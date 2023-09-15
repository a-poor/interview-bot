package main

import (
	bolt "go.etcd.io/bbolt"
)

func newStorageClient(path string) (*bolt.DB, error) {
	return bolt.Open(path, 0600, nil)
}
