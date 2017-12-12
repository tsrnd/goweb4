package models

import (
	"fmt"

	"github.com/goweb4/database"
)

type User struct {
	Model
	ProId    string  `schema:"pro_id"`
	UserName string  `schema:"username"`
	Password string  `schema:"password"`
	Email    string  `schema:"email"`
	Gender   string  `schema:"gender"`
	Role     string  `schema:"role"`
	Avatar   string  `schema:"avatar"`
	Phone    string  `schema:"phone"`
	Address  string  `schema:"address"`
	Provider string  `schema:"provider"`
	Orders   []Order //has many order
}

const ADMIN_ROLE = "admin"
const USER_ROLE = "normal_user"

// func GetUser(userInfo User) (user User) {
// 	database.DBCon.Where(&userInfo).First(&user)
// 	return user
// }

func GetUserById(id uint) (user User, err error) {
	err = database.DBCon.QueryRow("SELECT password, email, role FROM users where id = $1", id).Scan(&user.Password, &user.Email, &user.Role)
	if err != nil {
		fmt.Println("get user by id has an error: ", err)
	}
	return user, err
}

func GetUserByUserName(name string) (user User, err error) {
	err = database.DBCon.QueryRow("SELECT password, username, email, phone, address, role FROM users where username = $1", name).Scan(&user.Password, &user.UserName, &user.Email, &user.Phone, &user.Address, &user.Role)
	if err != nil {
		fmt.Println("get user by name has an error: ", err)
	}
	return user, err
}

func CreateUser(user User) (err error) {
	err = database.DBCon.
		QueryRow("INSERT INTO users (user_name, email, password, phone, address, created_at) VALUES($1,$2,$3,$4,$5,$6) returning id;",
			user.UserName, user.Email, user.Password, user.Phone, user.Address, user.CreatedAt).Scan(&user.ID)
	return err
}

func UpdateUser(user *User) (errUpdate error) {
	if errUpdate == nil {
		errUpdate = database.DBCon.QueryRow("UPDATE users SET user_name=$1, email=$2, password=$3, phone=$4, address=$5, created_at=$6 WHERE id = $1").Scan(&user)
	}
	return errUpdate
}

func DeleteUser(id uint) (err error) {
	user := User{}
	user, err = GetUserById(id)
	if err == nil {
		err = database.DBCon.QueryRow("DELETE FROM users WHERE id = $1", id).Scan(&user)
	}
	return err
}
