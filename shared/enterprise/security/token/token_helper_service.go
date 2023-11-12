package token

import (
	"context"
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/golang-jwt/jwt/v5"
	"github.com/life4/genesis/maps"
	"github.com/life4/genesis/slices"
	"github.com/sigmasee/sigmasee/shared/enterprise/errors"
	"github.com/sigmasee/sigmasee/shared/enterprise/tuples"
	"go.uber.org/zap"
)

type tokenHelperService struct {
	ctx                 context.Context
	logger              *zap.SugaredLogger
	identityProviderUrl string
	globalMutex         sync.Mutex
	rsaPublicKeys       []*rsa.PublicKey
}

type jwkKey struct {
	Alg string `json:"alg"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N   string `json:"n"`
}

type jwk struct {
	Keys []jwkKey `json:"keys"`
}

func NewTokenHelperService(
	ctx context.Context,
	logger *zap.SugaredLogger,
	jobScheduler *gocron.Scheduler,
	identityProviderUrl string,
	startupTimeJwksKeys string,
	refreshRsaKeys bool) (*tokenHelperService, error) {
	instance := &tokenHelperService{
		ctx:                 ctx,
		logger:              logger,
		identityProviderUrl: identityProviderUrl,
		globalMutex:         sync.Mutex{},
		rsaPublicKeys:       []*rsa.PublicKey{},
	}

	if len(startupTimeJwksKeys) > 0 {
		logger.Infof("Loading keys from provided initial key list...")

		jwk := jwk{}
		if err := json.Unmarshal([]byte(startupTimeJwksKeys), &jwk); err != nil {
			return nil, err
		}

		if err := instance.storeJwkKeysAsRsaPublicKeys(&jwk); err != nil {
			return nil, err
		}

		logger.Infof("Keys are successfully loaded")
	}

	if !refreshRsaKeys {
		return instance, nil
	}

	if _, err := jobScheduler.Every(2).Minutes().Do(func() {
		if err := instance.Run(); err != nil {
			logger.Errorf("Failed to convert JWK key to RSA public key. Error: %v", err)
		}
	}); err != nil {
		return nil, err
	}

	return instance, nil
}

func (s *tokenHelperService) RunAsync() {
	go func() {
		if err := s.Run(); err != nil {
			s.logger.Errorf("Failed to convert JWK key to RSA public key. Error: %v", err)
		}
	}()
}

func (s *tokenHelperService) Run() error {
	s.globalMutex.Lock()
	defer s.globalMutex.Unlock()

	jwk, err := s.downloadPublicKey()
	if err != nil {
		s.logger.Errorf("Failed to retrieve JWK keys. Error: %v", err)

		return err
	}

	return s.storeJwkKeysAsRsaPublicKeys(jwk)
}

func (s *tokenHelperService) VerifyAndGetAllClaims(token string) (map[string]string, error) {
	claims, err := s.verifyTokenAndGetClaims(token)
	if err != nil {
		return nil, err
	}

	result := maps.MapValues(claims, func(item interface{}) tuples.ValueErrorTuple[string] {
		value, ok := item.(string)
		if !ok {
			return tuples.ValueErrorTuple[string]{
				Value: "",
				Error: nil,
			}
		}

		return tuples.ValueErrorTuple[string]{
			Value: value,
			Error: nil,
		}
	})

	filteredStringClaims := make(map[string]string)
	for key, value := range tuples.GetMapValues(result) {
		if len(key) > 0 {
			filteredStringClaims[key] = value
		}
	}

	return filteredStringClaims, nil
}

func (s *tokenHelperService) downloadPublicKey() (*jwk, error) {
	client := http.Client{Timeout: 2 * time.Minute}
	resp, err := client.Get(s.identityProviderUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jwk := jwk{}
	if err = json.Unmarshal(body, &jwk); err != nil {
		return nil, err
	}

	return &jwk, nil
}

func (s *tokenHelperService) storeJwkKeysAsRsaPublicKeys(jwk *jwk) error {
	result := slices.Map(jwk.Keys, func(item jwkKey) tuples.ValueErrorTuple[*rsa.PublicKey] {
		rsaPublicKey, err := s.convertKey(item)
		if err != nil {
			return tuples.ValueErrorTuple[*rsa.PublicKey]{
				Value: nil,
				Error: err,
			}
		}

		return tuples.ValueErrorTuple[*rsa.PublicKey]{
			Value: rsaPublicKey,
			Error: nil,
		}
	})

	if err := tuples.ReduceErrors(result); err != nil {
		return err
	}

	s.rsaPublicKeys = tuples.GetValues(result)

	return nil
}

func (s *tokenHelperService) convertKey(jwkKey jwkKey) (*rsa.PublicKey, error) {
	decodedE, err := base64.RawURLEncoding.DecodeString(jwkKey.E)
	if err != nil {
		return nil, err
	}

	if len(decodedE) < 4 {
		ndata := make([]byte, 4)
		copy(ndata[4-len(decodedE):], decodedE)
		decodedE = ndata
	}

	pubKey := &rsa.PublicKey{
		N: &big.Int{},
		E: int(binary.BigEndian.Uint32(decodedE[:])),
	}

	decodedN, err := base64.RawURLEncoding.DecodeString(jwkKey.N)
	if err != nil {
		return nil, err
	}

	pubKey.N.SetBytes(decodedN)

	return pubKey, nil
}

func (s *tokenHelperService) verifyTokenAndGetClaims(token string) (jwt.MapClaims, error) {
	for _, rsaPublicKey := range s.rsaPublicKeys {
		claims := jwt.MapClaims{}

		jwtToken, err := jwt.ParseWithClaims(token, claims, func(_ *jwt.Token) (interface{}, error) {
			return rsaPublicKey, nil
		})

		if err == nil && jwtToken.Valid {
			return claims, nil
		}
	}

	return nil, errors.ErrFailedToVerifyToken
}
