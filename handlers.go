package chief

import (
	"github.com/kylelemons/go-gypsy/yaml"
	"github.com/vaughan0/go-ini"
)

type handler interface {
	Import(raw string) error
	ImportFile(filename string) error
	Get(raw string) (interface{}, error)
	Export() string
}

type handlerJson struct {
	m   map[string]interface{}
	raw string
}

type handlerIni struct {
	m   *ini.File
	raw string
}

type handlerYaml struct {
	m   *yaml.File
	raw string
}
