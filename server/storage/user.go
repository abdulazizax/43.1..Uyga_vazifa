package storage

import (
	"database/sql"
	"server/models"
)

type User struct {
	DB *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{DB: db}
}

func (u *User) CreateUser(req models.UserRequest) (*models.UserResponse, error) {
	query := `
		INSERT INTO users (name, email, age) 
		VALUES ($1, $2, $3) 
		RETURNING id, name, email, age
	`
	row := u.DB.QueryRow(query, req.Name, req.Email, req.Age)
	var user models.UserResponse
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) GetAllUsers() ([]*models.UserResponse, error) {
	query := `
		SELECT id, name, email, age 
		FROM users
	`

	rows, err := u.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.UserResponse

	for rows.Next() {
		var user models.UserResponse
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *User) GetUserByID(id int) (*models.UserResponse, error) {
	query := `
		SELECT id, name, email, age
		FROM users
		WHERE id = $1
	`

	row := u.DB.QueryRow(query, id)

	var user models.UserResponse

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) UpdateUserByID(id int, req models.UserRequest) (*models.UserResponse, error) {
	query := `
		UPDATE users
		SET name = $1, email = $2, age = $3
		WHERE id = $4
		RETURNING id, name, email, age
	`

	row:= u.DB.QueryRow(query, req.Name, req.Email, req.Age, id)

	var user models.UserResponse

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) DeleteUserByID(id int) (*models.UserResponse, error) {
	query := `
		DELETE FROM users
		WHERE id = $1
		RETURNING id, name, email, age
	`

	row := u.DB.QueryRow(query, id)

	var user models.UserResponse

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
