package configuration

import "fmt"

type AppConfig struct {
	ListeningInterface      string `yaml:"listeningInterface" env:"sigmasee_APP_LISTENING_INTERFACE"`
	DomainSource            string `yaml:"domainSource" env:"sigmasee_APP_DOMAIN_SOURCE"`
	AppSource               string `yaml:"appSource" env:"sigmasee_APP_APP_SOURCE"`
	PublicWebSiteBaseDomain string `yaml:"publicWebsiteBaseDomain" env:"sigmasee_APP_PUBLIC_WEBSITE_BASE_DOMAIN"`
	WebAppBaseDomain        string `yaml:"webappBaseDomain" env:"sigmasee_APP_WEBAPP_BASE_DOMAIN"`
}

type IntercomConfig struct {
	Secret string `yaml:"secret" env:"sigmasee_INTERCOM_SECRET"`
}

type AwsLambdaConfig struct {
	IsRetryTopic    bool `yaml:"isRetryTopic" env:"sigmasee_AWS_LAMBDA_IS_RETRY_TOPIC"`
	RetryTopicIndex int  `yaml:"RetryTopicIndex" env:"sigmasee_AWS_LAMBDA_RETRY_TOPIC_INDEX"`
}

func (s AppConfig) GetSource() string {
	return fmt.Sprintf("%s::%s", s.DomainSource, s.AppSource)
}
