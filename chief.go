package chief

const (
	_ = iota
	JSON
	YAML
	INI
)

type configType map[interface{}]handler

var (
	config configType
)

func New() *configType {
	config = make(configType)

	config[JSON] = new(handlerJson)
	config[YAML] = new(handlerYaml)
	config[INI] = new(handlerIni)

	return &config
}

func (c *configType) Import(t interface{}, raw string) error {
	return config[t].Import(raw)
}

func (c *configType) ImportFile(t interface{}, filename string) error {
	return config[t].ImportFile(filename)
}

func (c *configType) Get(t interface{}, raw string) (interface{}, error) {
	return config[t].Get(raw)
}

func (c *configType) Export(t interface{}) string {
	return config[t].Export()
}
