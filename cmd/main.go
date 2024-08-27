package main

import (
	handler "golang-cassandra-crud/internal/hadler"
	"golang-cassandra-crud/internal/model"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// Creating a new cluster configuration
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "system"
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Unable to connect to Cassandra cluster to create keyspace:", err)
	}
	defer session.Close()

	// Create keyspace and table
	createKeyspaceAndTable(session)

	cluster.Keyspace = "example"
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("Unable to connect to Cassandra cluster with new keyspace:", err)
	}
	defer session.Close()

	log.Println("Connected to Cassandra with keyspace 'example'!")

	user := model.User{
		ID:        gocql.TimeUUID(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}

	err = handler.CreateUser(session, user)
	if err != nil {
		log.Fatal("Failed to create user:", err)
	}
	log.Println("User created successful")

	_, err = handler.GetUser(session, user.ID)
	if err != nil {
		log.Fatal("failed to getuser", err)
	}
	log.Println("get user by id is sucessfull")
}

func createKeyspaceAndTable(session *gocql.Session) {
	if err := session.Query(`CREATE KEYSPACE IF NOT EXISTS example WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 1}`).Exec(); err != nil {
		log.Fatal("Failed to create keyspace:", err)
	}

	if err := session.Query(`CREATE TABLE IF NOT EXISTS example.users (
        id UUID,
        first_name text,
        last_name text,
        email text,
        PRIMARY KEY ((first_name, last_name), id)
    )`).Exec(); err != nil {
		log.Fatal("Failed to create table:", err)
	}

	log.Println("Keyspace and table created successfully!")
}
