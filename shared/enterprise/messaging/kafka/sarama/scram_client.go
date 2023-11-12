package kafka

import (
	"crypto/sha256"
	"crypto/sha512"

	"github.com/xdg-go/scram"
)

type xdgScramClient struct {
	*scram.Client
	*scram.ClientConversation
	scram.HashGeneratorFcn
}

var (
	SHA256 scram.HashGeneratorFcn = sha256.New
	SHA512 scram.HashGeneratorFcn = sha512.New
)

func (xsc *xdgScramClient) Begin(userName, password, authzID string) (err error) {
	xsc.Client, err = xsc.HashGeneratorFcn.NewClient(userName, password, authzID)

	if err != nil {
		return err
	}

	xsc.ClientConversation = xsc.Client.NewConversation()

	return nil
}

func (xsc *xdgScramClient) Step(challenge string) (string, error) {
	return xsc.ClientConversation.Step(challenge)
}

func (xsc *xdgScramClient) Done() bool {
	return xsc.ClientConversation.Done()
}
