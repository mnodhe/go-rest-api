package models

import (
	"go-rest-api/db"
	"go-rest-api/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email,password) VALUES (?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.Id = userId
	return nil
}
func (u User) GetUserByEmail(id string) (*User, error) {
	query := "SELECT * FROM users WHERE email=?"
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&u.Id, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func (u User) IsEmailExist(email string) (bool, error) {
	query := "SELECT COUNT(*) FROM users WHERE email=?"
	var count int
	err := db.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
