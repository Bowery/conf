// Copyright 2015 Bowery, Inc.

package conf

import (
	"os"
	"testing"
)

var (
	testYAMLDBPath = "test_yaml_db"
	testYAMLDB     *YAML
	testYAMLDBData *testconfig
)

func TestNewYAML(t *testing.T) {
	var err error
	testYAMLDB, err = NewYAML(testYAMLDBPath)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveYAML(t *testing.T) {
	testYAMLDBData = new(testconfig)
	testYAMLDBData.Field = "data"

	err := testYAMLDB.Save(testYAMLDBData)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadYAML(t *testing.T) {
	err := testYAMLDB.Load(testYAMLDBData)
	if err != nil {
		t.Error(err)
	}

	os.RemoveAll(testYAMLDBPath)
}
