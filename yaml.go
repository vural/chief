package chief

import (
	"io/ioutil"

	"github.com/kylelemons/go-gypsy/yaml"
)

func (c *handlerYaml) Import(raw string) error {
	c.m = yaml.Config(raw)
	c.raw = raw
	return nil
}

func (c *handlerYaml) ImportFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	c.raw = string(content)
	c.m = yaml.Config(c.raw)
	return nil
}

func (c *handlerYaml) Get(raw string) (interface{}, error) {
	return c.m.Get(raw)
}

func (c *handlerYaml) Export() string {
	return c.raw
}
