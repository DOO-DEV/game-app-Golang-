package config

import (
	"fmt"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
	"log"
	"strings"
)

const (
	prefix    = "GAMEAPP_"
	delimiter = "."
	separator = "__"
)

func callbackEnv(source string) string {
	base := strings.ToLower(strings.TrimPrefix(source, prefix))

	return strings.ReplaceAll(base, separator, delimiter)
}

func New() Config {
	k := koanf.New(".")

	// load default configuration from default function
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider("config.yml"), yaml.Parser()); err != nil {
		log.Printf("error loading config.yml: %s", err)
	}

	// load environment variables
	if err := k.Load(env.Provider(prefix, delimiter, callbackEnv), nil); err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	var instance Config
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshaling config: %s", err)
	}

	if instance.Debug {
		fmt.Printf("%+v\n", instance)
	}
	return instance
}
