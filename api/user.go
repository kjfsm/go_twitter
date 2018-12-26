package api

import (
	"encoding/json"

	"github.com/go-pg/pg"
)

// User ユーザー情報
type User struct {
	Name  string `json;"name"`
	Email string `json:"email`
}

// GetUsers ユーザー情報を取得
func GetUsers() []byte {
	tDatabase := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "password",
		Database: "go_twitter",
	})
	defer tDatabase.Close()

	var tUsers []User
	err := tDatabase.Select(&tUsers)
	if err != nil {
		panic(err)
	}

	tJSONBytes, err := json.Marshal(tUsers)
	if err != nil {
		panic(err)
	}
	return tJSONBytes
}
