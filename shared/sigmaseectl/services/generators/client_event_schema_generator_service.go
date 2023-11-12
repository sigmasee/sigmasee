package generators

import (
	"path"

	"github.com/sigmasee/sigmasee/shared/enterprise/os"
	"go.uber.org/zap"
)

type ClientEventSchemaGeneratorService interface {
	Generate(
		protobufFilePath string,
		outputPath string) error
}

type clientEventSchemaGeneratorService struct {
	logger   *zap.SugaredLogger
	osHelper os.OsHelper
}

func NewClientEventSchemaGeneratorService(
	logger *zap.SugaredLogger,
	osHelper os.OsHelper) (ClientEventSchemaGeneratorService, error) {
	return &clientEventSchemaGeneratorService{
		logger:   logger,
		osHelper: osHelper,
	}, nil
}

func (cesgs *clientEventSchemaGeneratorService) Generate(
	protobufFilePath string,
	outputPath string) error {
	return cesgs.osHelper.CopyFile(
		protobufFilePath,
		path.Join(outputPath, "schema", "schema.proto"))
}
