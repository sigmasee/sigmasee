package configuration

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/sigmasee/sigmasee/shared/enterprise/os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

type unmarshaller func([]byte, interface{}) error

type ConfigurationHelper interface {
	LoadYaml(configFilePath string, config interface{}) error
}

type configurationHelper struct {
	logger   *zap.SugaredLogger
	osHelper os.OsHelper
}

func NewConfigurationHelper(
	logger *zap.SugaredLogger,
	osHelper os.OsHelper,
) (ConfigurationHelper, error) {
	return &configurationHelper{
		logger:   logger,
		osHelper: osHelper,
	}, nil
}

func (ch *configurationHelper) LoadYaml(configFilePath string, config interface{}) error {
	if !ch.osHelper.FileExist(configFilePath) {
		ch.logger.Warnf("Configurtion file '%s' does not exist, returning provided configuration untouched!!!", configFilePath)
		return nil
	}

	fileContent, err := ch.osHelper.GetFileAsByteArray(configFilePath)
	if err != nil {
		return err
	}

	return ch.unmarshal(fileContent, config, yaml.Unmarshal)
}

func (ch *configurationHelper) unmarshal(fileContent []byte, config interface{}, um unmarshaller) error {
	if err := um(fileContent, config); err != nil {
		return err
	}

	return ch.dive(reflect.ValueOf(config), "")
}

func (ch *configurationHelper) dive(v reflect.Value, env string) error {
	switch v.Kind() {
	case reflect.Ptr:
		originalValue := v.Elem()
		return ch.dive(originalValue, env)

	case reflect.Interface:
		originalValue := v.Elem()
		return ch.dive(originalValue, env)

	case reflect.Struct:
		for i := 0; i < v.NumField(); i += 1 {
			tag := reflect.Indirect(v).Type().Field(i).Tag
			env := tag.Get("env")
			err := ch.dive(v.Field(i), env)
			if err != nil {
				return err
			}
		}
		return nil

	case reflect.Slice:
		for i := 0; i < v.Len(); i += 1 {
			err := ch.dive(v.Index(i), env)
			if err != nil {
				return err
			}
		}
		return nil

	case reflect.Map:
		for _, key := range v.MapKeys() {
			originalValue := v.MapIndex(key)
			err := ch.dive(originalValue, env)
			if err != nil {
				return err
			}
		}
		return nil

	default:
		return ch.envUnmarshaller(v, env)
	}
}

func (ch *configurationHelper) envUnmarshaller(elem reflect.Value, env string) error {
	if val := ch.osHelper.GetEnvironmentVariable(env); val != "" {
		tipe := elem.Type()

		switch tipe.Name() {
		case "string":
			elem.SetString(val)
		case "int", "int8", "int16", "int32", "int64":
			elem.SetInt(toInt(val))
		case "uint", "uint8", "uint16", "uint32", "uint64":
			elem.SetUint(toUint(val))
		case "float32", "float64":
			elem.SetFloat(toFloat(val))
		case "bool":
			elem.SetBool(toBool(val))
		}
	}

	return nil
}

func toBool(sbool string) bool {
	return strings.ToLower(sbool) == "true"
}

func toInt(sint string) int64 {
	num, _ := strconv.ParseInt(sint, 10, 64)
	return num
}

func toUint(sint string) uint64 {
	num, _ := strconv.ParseUint(sint, 10, 64)
	return num
}

func toFloat(sfloat string) float64 {
	num, _ := strconv.ParseFloat(sfloat, 64)
	return num
}
