package httpboomer

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func (tc *TestCase) toStruct() *TCase {
	tcStruct := TCase{
		Config: tc.Config,
	}
	for _, step := range tc.TestSteps {
		tcStruct.TestSteps = append(tcStruct.TestSteps, step.ToStruct())
	}
	return &tcStruct
}

func (tc *TestCase) dump2JSON(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		log.Printf("convert absolute path error: %v, path: %v", err, path)
		return err
	}
	log.Printf("dump testcase to json path: %s", path)
	tcStruct := tc.toStruct()
	file, _ := json.MarshalIndent(tcStruct, "", "    ")
	err = ioutil.WriteFile(path, file, 0644)
	if err != nil {
		log.Printf("dump json path error: %v", err)
		return err
	}
	return nil
}

func (tc *TestCase) dump2YAML(path string) error {
	path, err := filepath.Abs(path)
	if err != nil {
		log.Printf("convert absolute path error: %v, path: %v", err, path)
		return err
	}
	log.Printf("dump testcase to yaml path: %s", path)

	// init yaml encoder
	buffer := new(bytes.Buffer)
	encoder := yaml.NewEncoder(buffer)
	encoder.SetIndent(4)

	// encode
	tcStruct := tc.toStruct()
	err = encoder.Encode(tcStruct)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, buffer.Bytes(), 0644)
	if err != nil {
		log.Printf("dump yaml path error: %v", err)
		return err
	}
	return nil
}

func loadFromJSON(path string) (*TCase, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		log.Printf("convert absolute path error: %v, path: %v", err, path)
		return nil, err
	}
	log.Printf("load testcase from json path: %s", path)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("dump json path error: %v", err)
		return nil, err
	}

	tc := &TCase{}
	err = json.Unmarshal(file, tc)
	return tc, err
}

func loadFromYAML(path string) (*TCase, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		log.Printf("convert absolute path error: %v, path: %v", err, path)
		return nil, err
	}
	log.Printf("load testcase from yaml path: %s", path)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("dump yaml path error: %v", err)
		return nil, err
	}

	tc := &TCase{}
	err = yaml.Unmarshal(file, tc)
	return tc, err
}
