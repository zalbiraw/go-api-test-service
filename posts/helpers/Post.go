package helpers

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/zalbiraw/go-api-test-service/posts/graph/model"
)

var posts []model.Post

func LoadPosts() error {
	jsonFile, err := os.Open("./posts/helpers/db.json")

	if nil != err {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &posts)

	return nil
}

func GetPosts() *[]model.Post {
	return &posts
}
