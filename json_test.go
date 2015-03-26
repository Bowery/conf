// Copyright 2014 Bowery, Inc.

package conf

import (
	"os"
	"testing"
)

var (
	testJSONDBPath = "test_json_db"
	testJSONDB     *JSON
	testJSONDBData *testconfig
)

func TestNewJSON(t *testing.T) {
	var err error
	testJSONDB, err = NewJSON(testJSONDBPath)
	if err != nil {
		t.Error(err)
	}
}

func TestSaveJSON(t *testing.T) {
	testJSONDBData = new(testconfig)
	testJSONDBData.Field = "data"

	err := testJSONDB.Save(testJSONDBData)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadJSON(t *testing.T) {
	err := testJSONDB.Load(testJSONDBData)
	if err != nil {
		t.Error(err)
	}

	os.RemoveAll(testJSONDBPath)
}
