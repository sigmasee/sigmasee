package ses

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/sigmasee/sigmasee/shared/enterprise/os"
)

type AwsSesHelper interface {
	GetSesService() (*ses.SES, error)
}

type awsSesHelper struct {
	osHelper         os.OsHelper
	cachedSesService *ses.SES
}

func NewAwsSesHelper(osHelper os.OsHelper) (AwsSesHelper, error) {
	return &awsSesHelper{
		osHelper: osHelper,
	}, nil
}

func (s *awsSesHelper) GetSesService() (*ses.SES, error) {
	if s.cachedSesService != nil {
		return s.cachedSesService, nil
	}

	var awsSession *session.Session
	var err error

	region := s.osHelper.GetEnvironmentVariable("AWS_REGION")
	if len(region) == 0 {
		awsSession, err = session.NewSession()
	} else {
		awsSession, err = session.NewSession(
			&aws.Config{
				Region: aws.String(region),
			},
		)
	}

	if err != nil {
		return nil, err
	}

	s.cachedSesService = ses.New(awsSession)

	return s.cachedSesService, nil
}
