package model

import (
	"errors"

	"github.com/bits-and-atoms/Go_REST_API/db"
	"github.com/bits-and-atoms/Go_REST_API/utils"
)

type User struct{
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error{
	query := "insert into users (email,password) values (?,?)"
	st,err := db.DB.Prepare(query)
	if err != nil{
		return err
	}
	defer st.Close()
	hashpass , err := utils.HashPassword(u.Password)
	if err != nil{
		return err
	}
	result,err := st.Exec(u.Email,hashpass)
	if err != nil{
		return err
	}
	id,err := result.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCreds() error{
	query := "select id,password from users where email = ?"
	row := db.DB.QueryRow(query,u.Email)
	var hashpass string
	var id int64
	err := row.Scan(&id,&hashpass)
	if err != nil{
		return err
	}
	isok := utils.CheckPassHash(u.Password,hashpass)
	if isok == false{
		return errors.New("Credentials invalid")
	}
	u.ID = id
	return nil
}