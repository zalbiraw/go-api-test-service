package helpers

import (
	"encoding/json"
	"github.com/zalbiraw/go-api-test-service/posts/graph/model"
	"io/ioutil"
	"os"
)

func GetPosts() (*[]model.Post, error) {
	jsonFile, err := os.Open("./graph/db.json")

	if nil != err {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var posts []model.Post
	json.Unmarshal(byteValue, &posts)

	return &posts, nil
}
