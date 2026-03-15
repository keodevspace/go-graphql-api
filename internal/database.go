// Package internal provides a mock data store.
package internal

import "github.com/keodevspace/go-graphql-api/graph/model"

type DB struct {
	Tasks []*model.Task // Maiúsculo para ser visível fora da pasta
}

func NewDB() *DB {
	return &DB{
		Tasks: make([]*model.Task, 0),
	}
}
