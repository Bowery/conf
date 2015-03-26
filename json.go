// Copyright 2015 Bowery, Inc.

package conf

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

// JSON implements DB for a json-based local db.
type JSON struct {
	path string
}

// NewJSON creates a new JSON local database at the given location.
func NewJSON(path string) (*JSON, error) {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm|os.ModeDir)
	if err != nil {
		return nil, err
	}

	return &JSON{
		path: path,
	}, nil
}

// Load decodes a json file into the given data.
func (j *JSON) Load(data interface{}) error {
	file, err := os.OpenFile(j.path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}

	return nil
}

// Save encodes the given data to a json file.
func (j *JSON) Save(data interface{}) error {
	file, err := os.OpenFile(j.path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	dat, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(dat)
	_, err = io.Copy(file, buf)
	return err
}
