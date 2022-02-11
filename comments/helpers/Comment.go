package helpers

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/zalbiraw/go-api-test-service/comments/graph/model"
)

var comments []model.Comment

func LoadComments() error {
	jsonFile, err := os.Open("./comments/helpers/db.json")

	if nil != err {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &comments)
	return nil
}

func GetComments() []model.Comment {
	return comments
}
