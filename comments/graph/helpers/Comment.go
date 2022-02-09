package helpers

import (
	"encoding/json"
	"github.com/zalbiraw/go-api-test-service/comments/graph/model"
	"io/ioutil"
	"os"
)

func GetComments() (*[]model.Comment, error) {
	jsonFile, err := os.Open("./comments/graph/db.json")

	if nil != err {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var comments []model.Comment
	json.Unmarshal(byteValue, &comments)

	return &comments, nil
}
