// Copyright 2015 Bowery, Inc.

package conf

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// YAML implements DB for a yaml-based local db.
type YAML struct {
	path string
}

// NewYAML creates a new YAML local database at the given location.
func NewYAML(path string) (*YAML, error) {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm|os.ModeDir)
	if err != nil {
		return nil, err
	}

	return &YAML{
		path: path,
	}, nil
}

// Load decodes a yaml file into the given data.
func (y *YAML) Load(data interface{}) error {
	file, err := ioutil.ReadFile(y.path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(file, &data)
}

// Save encodes the given data into a yaml file.
func (y *YAML) Save(data interface{}) error {
	file, err := os.OpenFile(y.path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	dat, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(dat)
	_, err = io.Copy(file, buf)
	return err
}
