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

type ClientEventGenerateGeneratorService interface {
	Generate(
		packageName string,
		outputPath string) error
}

type clientEventGenerateGeneratorService struct {
	logger   *zap.SugaredLogger
	osHelper os.OsHelper
}

func NewClientEventGenerateGeneratorService(
	logger *zap.SugaredLogger,
	osHelper os.OsHelper) (ClientEventGenerateGeneratorService, error) {
	return &clientEventGenerateGeneratorService{
		logger:   logger,
		osHelper: osHelper,
	}, nil
}

func (cehgs *clientEventGenerateGeneratorService) Generate(
	packageName string,
	outputPath string) error {
	type data struct {
		PackageName string
	}

	tmpl, err := template.New("").Parse(clienteventtemplates.GetGenerate())
	if err != nil {
		return err
	}

	var processed bytes.Buffer
	if err = tmpl.ExecuteTemplate(&processed, "", &data{PackageName: packageName}); err != nil {
		return err
	}

	formatted, err := format.Source(processed.Bytes())
	if err != nil {
		return fmt.Errorf("could not format processed template: %v", err)
	}

	err = cehgs.osHelper.CreateDir(outputPath)
	if err != nil {
		return err
	}

	return cehgs.osHelper.CreateBinaryFile(
		path.Join(outputPath, "generate.go"),
		formatted)
}
