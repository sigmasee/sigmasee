package configuration

import "fmt"

type AppConfig struct {
	ListeningInterface      string `yaml:"listeningInterface" env:"SIGMASEE_APP_LISTENING_INTERFACE"`
	DomainSource            string `yaml:"domainSource" env:"SIGMASEE_APP_DOMAIN_SOURCE"`
	AppSource               string `yaml:"appSource" env:"SIGMASEE_APP_APP_SOURCE"`
	PublicWebSiteBaseDomain string `yaml:"publicWebsiteBaseDomain" env:"SIGMASEE_APP_PUBLIC_WEBSITE_BASE_DOMAIN"`
	WebAppBaseDomain        string `yaml:"webappBaseDomain" env:"SIGMASEE_APP_WEBAPP_BASE_DOMAIN"`
}

type IntercomConfig struct {
	Secret string `yaml:"secret" env:"SIGMASEE_INTERCOM_SECRET"`
}

type AwsLambdaConfig struct {
	IsRetryTopic    bool `yaml:"isRetryTopic" env:"SIGMASEE_AWS_LAMBDA_IS_RETRY_TOPIC"`
	RetryTopicIndex int  `yaml:"RetryTopicIndex" env:"SIGMASEE_AWS_LAMBDA_RETRY_TOPIC_INDEX"`
}

func (s AppConfig) GetSource() string {
	return fmt.Sprintf("%s::%s", s.DomainSource, s.AppSource)
}
