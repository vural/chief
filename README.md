## chief ##

chief is configuration manager for golang.

### handlers ###
    - json
    - yaml
    - ini

### installation ###

```
go get github.com/vural/chief
```

### usage ###

**example for json, yaml and ini handlers**
```go
package main

import (
    "fmt"
    "github.com/vural/chief"
)

var yamlD = `
map:
  foo: bar
  key: true
list:
  - item1
  - item2
config:
  server:
    - github.com
    - example.com
`

var jsonD = `{
        "foo": "bar",
        "key": true,
        "api": {
            "port": 3000
        }
}`

var iniD = `
# I am a comment
; So am I!

[api]
foo = bar
`

func main() {
    i := chief.New()

    // json handler
    json := i.Import(chief.JSON, jsonD)

    if json != nil {
        panic(json)
    }

    // yaml handler
    yaml := i.Import(chief.YAML, yamlD)

    if yaml != nil {
        panic(yaml)
    }

    // ini handler
    ini := i.Import(chief.INI, iniD)

    if ini != nil {
        panic(ini)
    }

    // import file
    // chief.JSON or chief.YAMl or chief.INI
    // jsonFile := i.ImportFile(chief.JSON, "example.file")

    // get value for json
    valJson, err := i.Get(chief.JSON, "api.port")

    if err != nil {
        panic(err)
    }

    // get value for yaml
    valYaml, err := i.Get(chief.YAML, "config.server[0]")

    if err != nil {
        panic(err)
    }

    // get value for ini
    valIni, err := i.Get(chief.INI, "api.foo") // output => bar

    if err != nil {
        panic(err)
    }
    fmt.Println(valJson) // output => 3000
    fmt.Println(valYaml) // output => github.com
    fmt.Println(valIni)  // output => bar

    // exporting
    fmt.Println(i.Export(chief.JSON))
}
```

**not to like**
```go
 // valJson, err := i.GetBool(foo)
```

**like better**
```go
    valJson, err := i.Get(chief.JSON, "key") // output => true

    if err != nil {
        panic(err)
    }

    if valJson.(bool) {
        fmt.Println("true !")
    }
```

![](http://www.topito.com/wp-content/uploads/2013/01/code-34.gif)