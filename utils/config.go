package utils

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Data struct {
	Publisher string
	Topics    []string
}
type conf map[string]interface{}

var configFile string

func SetConfigFile(file string) {
	configFile = file
}

func GetConfig() string {
	return configFile
}

func GetStruct() Data {
	structData := Data{}
	return (structData)
}

func NewConfigFromFile(fname string) (conf, error) {
	var publish conf
	if _, err := toml.DecodeFile(fname, &publish); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return publish, nil
}

func GetData() conf {
	configFile := GetConfig()
	conf, err := NewConfigFromFile(configFile)
	if err != nil {
		fmt.Println("ERROR:", err)
		return nil
	}
	return conf
}

func (node *Data) Flush() {
	var config = map[string]interface{}{
		node.Publisher: node,
	}
	f, err := os.OpenFile(configFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var fileBuffer bytes.Buffer
	e := toml.NewEncoder(&fileBuffer)
	err1 := e.Encode(config)
	if err1 != nil {
		log.Fatal(err)
	}
	log.Println("FILE: ", fileBuffer.String())
	f.WriteString(fileBuffer.String())
}
