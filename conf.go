// Copyright 2015 Bowery, Inc.

package conf

// DB represents a local database which can load and save data.
type DB interface {
	Load(data interface{}) error
	Save(data interface{}) error
}
