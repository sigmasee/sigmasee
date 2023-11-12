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

type ClientEventHandlerGeneratorService interface {
	Generate(
		packageName string,
		eventType string,
		outputPath string) error
}

type clientEventHandlerGeneratorService struct {
	logger   *zap.SugaredLogger
	osHelper os.OsHelper
}

func NewClientEventHandlerGeneratorService(
	logger *zap.SugaredLogger,
	osHelper os.OsHelper) (ClientEventHandlerGeneratorService, error) {
	return &clientEventHandlerGeneratorService{
		logger:   logger,
		osHelper: osHelper,
	}, nil
}

func (cehgs *clientEventHandlerGeneratorService) Generate(
	packageName string,
	eventType string,
	outputPath string) error {
	type data struct {
		PackageName string
		EventType   string
	}

	tmpl, err := template.New("").Parse(clienteventtemplates.GetHandler())
	if err != nil {
		return err
	}

	var processed bytes.Buffer
	if err = tmpl.ExecuteTemplate(&processed, "", &data{PackageName: packageName, EventType: eventType}); err != nil {
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
		path.Join(outputPath, "handler_eventgen.go"),
		formatted)
}
