package helpers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/zalbiraw/go-api-test-service/services/graphql-subgraphs/users/graph/model"
)

var users []*model.User

func LoadUsers() error {
	byteValue, err := ioutil.ReadFile("../../../../helpers/users-db.json")

	if nil != err {
		return err
	}

	json.Unmarshal(byteValue, &users)

	return nil
}

func GetUsers() []*model.User {
	return users
}
