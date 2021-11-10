package utils

import (
	"strings"

	"cloud.google.com/go/datastore"
)

func CreateKey(kind string, keys ...string) *datastore.Key {
	if len(keys) == 0 {
		return datastore.IncompleteKey(kind, nil)
	}

	return datastore.NameKey(kind, keys[0], nil)
}

func ConvertUserIdKey(id string) string {
	return strings.Join([]string{"user", id}, ":")
}
