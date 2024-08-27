package model

import "github.com/gocql/gocql"

type User struct {
	ID        gocql.UUID
	FirstName string
	LastName  string
	Email     string
}