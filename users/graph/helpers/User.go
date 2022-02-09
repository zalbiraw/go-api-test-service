package helpers

import (
	"encoding/json"
	"github.com/zalbiraw/go-api-test-service/users/graph/model"
	"io/ioutil"
	"os"
)

func GetUsers() (*[]model.User, error) {
	jsonFile, err := os.Open("./users/graph/db.json")

	if nil != err {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []model.User
	json.Unmarshal(byteValue, &users)

	return &users, nil
}
