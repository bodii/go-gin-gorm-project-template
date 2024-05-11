package utils

import (
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

// Read the configuration content of a yaml file type
//
//	@this is a generic func [T]
//
// param:
//   - file [string]: a string file path
//
// return:
//   - [T]: type
func ReadYamlConfig[T any](filepath string) (t T) {
	if 1 > strings.LastIndex(filepath, ".yaml") {
		log.Fatalf("input file name not is %s yaml file\n", filepath)
	}

	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = yaml.Unmarshal(content, &t)
	if err != nil {
		log.Fatalf("%s -> %v\n", filepath, err)
		os.Exit(1)
	}

	log.Printf("read %s config succuee!", filepath)

	return
}

// Read the configuration content of a toml file type
//
//	@this is a generic func [T]
//
// param:
//   - file [string]: a string file path
//
// return:
//   - [T]: type
func ReadTomlConfig[T any](filepath string) (t T) {
	if 1 > strings.LastIndex(filepath, ".toml") {
		log.Fatalf("input file name not is %s toml file\n", filepath)
	}

	// var t T
	if _, err := toml.DecodeFile(filepath, &t); err != nil {
		log.Fatalf("%s -> %v\n", filepath, err)
	}

	log.Printf("read %s config succuee!", filepath)

	return
}
