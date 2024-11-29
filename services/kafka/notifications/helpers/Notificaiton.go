package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/zalbiraw/go-api-test-service/services/kafka/users/model"
)

var users []*model.User

// LoadUsers loads the list of users from the JSON file into the `users` slice.
func LoadUsers() error {
	byteValue, err := ioutil.ReadFile("/go/src/github.com/zalbiraw/go-api-test-service/helpers/users-db.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		return err
	}

	return nil
}

// GetRandomNotification generates a random notification and returns it in a map format that fits the Avro schema.
func GetRandomNotification() interface{} {
	// Ensure randomness
	rand.Seed(time.Now().UnixNano())

	// Randomly select `x` and `y` where x cannot be y
	x := rand.Intn(len(users))
	y := rand.Intn(len(users))

	// Ensure x and y are not the same
	for x == y {
		y = rand.Intn(len(users))
	}

	// Randomly choose a notification type and map it to the Avro enum values
	notificationTypes := []string{
		"liked",        // Map to "LIKE"
		"commented on", // Map to "COMMENT"
		"shared",       // Map to "SHARE"
	}

	// Map the string to enum-like value
	var notificationType string
	switch notificationTypes[rand.Intn(len(notificationTypes))] {
	case "liked":
		notificationType = "LIKE"
	case "commented on":
		notificationType = "COMMENT"
	case "shared":
		notificationType = "SHARE"
	}

	// Generate the notification message
	msg := fmt.Sprintf("%s %s %s's post", users[x].Name, notificationTypes[rand.Intn(len(notificationTypes))], users[y].Name)

	// Return the map that matches Avro schema
	return map[string]interface{}{
		"type": notificationType,
		"msg":  msg,
	}
}
