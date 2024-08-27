package handler

import ( 
	"golang-cassandra-crud/internal/model"
	"github.com/gocql/gocql"
)

func CreateUser(session *gocql.Session, user model.User) error {
	if err := session.Query(`INSERT INTO example.users (id, first_name, last_name, email) VALUES (?, ?, ?, ?)`,
		user.ID, user.FirstName, user.LastName, user.Email).Exec(); err != nil {
		return err
	}
	return nil 
}

func GetUser(session *gocql.Session, id gocql.UUID) (model.User, error) {
	var user model.User
	var firstName, lastName, email string

	if err := session.Query(`SELECT id, first_name, last_name, email FROM example.users WHERE id = ? LIMIT 1`, id).
		Scan(&user.ID, &firstName, &lastName, &email); err != nil {
		return user, err
	}

	user.FirstName = firstName
	user.LastName = lastName
	user.Email = email
	return user, nil
}

func UpdateUser(session *gocql.Session, id gocql.UUID, email string) error {
	if err := session.Query(`UPDATE example.users SET email = ? WHERE id = ?`, email, id).Exec(); err != nil {
		return err
	}
	return nil
}

func DeleteUser(session *gocql.Session, id gocql.UUID) error {
	if err := session.Query(`DELETE FROM example.users WHERE id = ?`, id).Exec(); err != nil {
		return err
	}
	return nil
} 
