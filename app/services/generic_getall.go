package services

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/domgolonka/threatdefender/app"
)

func GenericGetAll(app *app.App) (*[]string, error) {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/app/data/raw/generic_list.txt")
	if err != nil {
		app.Logger.Error(err)
		return nil, err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		app.Logger.Error(err)
		return nil, err
	}
	generic := strings.Split(string(b), "\n")

	return &generic, nil
}

func GenericGetEmail(app *app.App, emailAddress string) (*bool, error) {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/app/data/raw/generic_list.txt")
	if err != nil {
		app.Logger.Error(err)
		return nil, err
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		app.Logger.Error(err)
		return nil, err
	}
	hasEmail := strings.Contains(string(b), emailAddress)
	return &hasEmail, nil
}
