package helpers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/zalbiraw/go-api-test-service/services/rest/comments/model"
)

var comments []*model.Comment

func LoadComments() error {
	byteValue, err := ioutil.ReadFile("/go/src/github.com/zalbiraw/go-api-test-service/helpers/comments-db.json")

	if nil != err {
		return err
	}

	json.Unmarshal(byteValue, &comments)

	return nil
}

func GetComments() []*model.Comment {
	return comments
}
