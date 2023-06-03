package bd

import (
	"fmt"

	"github.com/divan1319/gambitoUser/models"
	"github.com/divan1319/gambitoUser/tools"
)

func SignUp(sign models.SignUp) error {
	fmt.Println("comienza funcion")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	query := "INSERT INTO users (User_Email,User_UUID,User_DateAdd) VALUES ('" + sign.UserEmail + "','" + sign.UserUUID + "','" + tools.FechaMySQL() + "')"

	_, err = Db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
