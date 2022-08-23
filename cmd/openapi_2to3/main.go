package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/general252/openapi_2_to_3"
	"gopkg.in/yaml.v3"
)

func main() {
	var (
		swag2file string
		swag3file string
	)

	flag.StringVar(&swag2file, "swag2", "swagger.json", "swag2 input file (*.json, *.yaml)")
	flag.StringVar(&swag3file, "swag3", "swagger3.json", "swag3 output file (*.json, *.yaml)")
	flag.Parse()

	swag2fileData, err := os.ReadFile(swag2file)
	if err != nil {
		fmt.Printf("read swag2 file fail. %v, [%v]", err, swag2file)
		return
	}

	var ylj YamlJson

	if ylj.isYaml(swag2file) {
		if swag2fileData, err = ylj.yamlToJson(swag2fileData); err != nil {
			fmt.Printf("yaml convert yaml fail. %v, [%v]", err, swag2file)
			return
		}
	}

	swag3fileData, err := openapi_2_to_3.Convert(swag2fileData)
	if err != nil {
		fmt.Printf("convert fail. %v", err)
		return
	}

	if ylj.isYaml(swag3file) {
		if swag3fileData, err = ylj.jsonToYaml(swag3fileData); err != nil {
			fmt.Printf("json convert to yaml fail. %v", err)
			return
		}
	} else {
		if swag3fileData, err = ylj.jsonFormat(swag3fileData); err != nil {
			fmt.Printf("json format fail. %v", err)
			return
		}
	}

	err = os.WriteFile(swag3file, swag3fileData, os.ModePerm)
	if err != nil {
		fmt.Printf("write swag3 file fail. %v", err)
		return
	}

	fmt.Printf("success. %v", swag3file)
}

type YamlJson struct {
}

func (*YamlJson) isYaml(swag2file string) bool {
	if strings.HasSuffix(swag2file, ".yaml") || strings.HasSuffix(swag2file, ".yml") {
		return true
	}

	return false
}

func (*YamlJson) jsonFormat(data []byte) ([]byte, error) {
	var val interface{}
	if err := json.Unmarshal(data, &val); err != nil {
		return nil, err
	}

	out, err := json.MarshalIndent(val, "", "  ")
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (*YamlJson) yamlToJson(data []byte) ([]byte, error) {
	var val interface{}
	if err := yaml.Unmarshal(data, &val); err != nil {
		return nil, err
	}

	out, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (*YamlJson) jsonToYaml(data []byte) ([]byte, error) {
	var val interface{}
	if err := json.Unmarshal(data, &val); err != nil {
		return nil, err
	}

	out, err := yaml.Marshal(val)
	if err != nil {
		return nil, err
	}

	return out, nil
}
