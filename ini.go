package chief

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/vaughan0/go-ini"
)

func (c *handlerIni) Import(raw string) error {
	r := bytes.NewReader([]byte(raw))
	file, err := ini.Load(r)
	c.m = &file
	c.raw = raw
	return err
}

func (c *handlerIni) ImportFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	r := bytes.NewReader([]byte(content))
	file, err := ini.Load(r)
	c.m = &file
	return err
}

func (c *handlerIni) Get(raw string) (interface{}, error) {

	output := strings.Split(raw, ".")

	if len(output) != 2 {
		return nil, errors.New("Should to be two args")
	}

	name, err := c.m.Get(output[0], output[1])
	if !err {
		return nil, errors.New("Not found")
	}
	return name, nil
}

func (c *handlerIni) Export() string {
	return c.raw
}
