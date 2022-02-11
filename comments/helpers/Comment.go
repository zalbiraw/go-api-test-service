package helpers

import (
	"encoding/json"
	"io/ioutil"
	
	"github.com/zalbiraw/go-api-test-service/comments/graph/model"
)

var comments []*model.Comment

func LoadComments() error {
	byteValue, err := ioutil.ReadFile("./comments/helpers/db.json")

	if nil != err {
		return err
	}

	json.Unmarshal(byteValue, &comments)

	return nil
}

func GetComments() []*model.Comment {
	return comments
}
