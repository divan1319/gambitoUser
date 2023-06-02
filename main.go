package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/divan1319/gambitoUser/awsgo"
	"github.com/divan1319/gambitoUser/bd"
	"github.com/divan1319/gambitoUser/models"
)

func main() {
	lambda.Start(LamdaExec)
}

func LamdaExec(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitAWS()

	if !ValidarParametros() {
		fmt.Println("Error en los parametros, no hay SecretName")
		err := errors.New("Error en los parametros, no hay secretName")

		return event, err
	}

	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
		case "sub":
			datos.UserUUID = att
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		return event, err
	}

}

func ValidarParametros() bool {
	var getParam bool

	_, getParam = os.LookupEnv("SecretName")

	return getParam
}
