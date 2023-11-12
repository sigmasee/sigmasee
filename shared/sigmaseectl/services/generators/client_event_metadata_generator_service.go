package generators

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"path"

	"github.com/sigmasee/sigmasee/shared/enterprise/os"
	clienteventtemplates "github.com/sigmasee/sigmasee/shared/sigmaseectl/templates/client/event"
	"go.uber.org/zap"
)

type ClientEventMetadataGeneratorService interface {
	Generate(
		packageName string,
		topicName string,
		retryTopicNamePrefix string,
		retryTopicNameCount int,
		deadLetterTopicName string,
		outputPath string) error
}

type clientEventMetadataGeneratorService struct {
	logger   *zap.SugaredLogger
	osHelper os.OsHelper
}

func NewClientEventMetadataGeneratorService(
	logger *zap.SugaredLogger,
	osHelper os.OsHelper) (ClientEventMetadataGeneratorService, error) {
	return &clientEventMetadataGeneratorService{
		logger:   logger,
		osHelper: osHelper,
	}, nil
}

func (cemgs *clientEventMetadataGeneratorService) Generate(
	packageName string,
	topicName string,
	retryTopicNamePrefix string,
	retryTopicNameCount int,
	deadLetterTopicName string,
	outputPath string) error {
	type data struct {
		PackageName          string
		TopicName            string
		RetryTopicNamePrefix string
		RetryTopicNameCount  int
		DeadLetterTopicName  string
	}

	tmpl, err := template.New("").Parse(clienteventtemplates.GetMetadata())
	if err != nil {
		return err
	}

	var processed bytes.Buffer
	if err = tmpl.ExecuteTemplate(
		&processed,
		"",
		&data{
			PackageName:          packageName,
			TopicName:            topicName,
			RetryTopicNamePrefix: retryTopicNamePrefix,
			RetryTopicNameCount:  retryTopicNameCount,
			DeadLetterTopicName:  deadLetterTopicName,
		}); err != nil {
		return err
	}

	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		return fmt.Errorf("could not format processed template: %v", err)
	}

	err = cemgs.osHelper.CreateDir(outputPath)
	if err != nil {
		return err
	}

	return cemgs.osHelper.CreateBinaryFile(
		path.Join(outputPath, "metadata_eventgen.go"),
		formatted)
}
