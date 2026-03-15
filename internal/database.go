// server layer for the GraphQL API
// internal/database.go
package internal

import "github.com/keodevspace/go-graphql-api/graph/model"

type DB struct {
	Tasks []*model.Task
}

func NewDB() *DB {
	return &DB{
		Tasks: []*model.Task{}, // Initializing an empty slice (List)
	}
}
