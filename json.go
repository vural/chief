package chief

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"reflect"
	"strings"
)

func (c *handlerJson) Import(raw string) error {
	c.raw = string(raw)
	return json.Unmarshal([]byte(c.raw), &c.m)
}

func (c *handlerJson) ImportFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	c.raw = string(content)
	return json.Unmarshal(content, &c.m)
}

func (c *handlerJson) Get(raw string) (interface{}, error) {
	output := strings.Split(raw, ".")

	var current map[string]interface{} = c.m

	for _, get := range output {
		if val, ok := current[get]; ok {

			if reflect.TypeOf(val).Kind() == reflect.Map {
				current = val.(map[string]interface{})
				continue
			}

			return val, nil
		}
	}

	return nil, errors.New("Not found ")
}

func (c *handlerJson) Export() string {
	return c.raw
}
