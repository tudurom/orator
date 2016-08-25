package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Site configuration from config.yaml
type SiteConfig struct {
	Title       string
	Subtitle    string
	Description string
	Keywords    string
	Author      string
	BasePath    string
	Extra       map[string]interface{}
}

func (sc *SiteConfig) ReadConfig(fpath string) {
	contents, err := ioutil.ReadFile(fpath)

	if err != nil {
		log.Fatal(err)
	}

	yaml.Unmarshal(contents, sc)
}
