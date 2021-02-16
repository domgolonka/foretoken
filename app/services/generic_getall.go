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

	app.Logger.Error(generic)
	return &generic, nil
}
