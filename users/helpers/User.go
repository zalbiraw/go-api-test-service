package helpers

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/zalbiraw/go-api-test-service/users/graph/model"
)

var users []model.User

func LoadUsers() error {
	jsonFile, err := os.Open("./users/helpers/db.json")

	if nil != err {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &users)

	return nil
}

func GetUsers() *[]model.User {
	return &users
}
