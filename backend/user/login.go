package user

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"gorm.io/gorm"
)

type Credentials struct {
	gorm.Model
	Username string `gorm:"username" json:"username"`
	Password string `gorm:"password" json:"password"`
}

func (c *Credentials) LoginHandler() error {
	SendLoginRequest(*c, os.Getenv("SERVER_LOGIN_API_URL"))
	return nil
}

func SendLoginRequest(credentials Credentials, url string) {
	jsonData, err := json.Marshal(credentials)
	if err != nil {
		log.Fatalf("Error of turning credentials in json: %s", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating POST request: %s", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %s", err)
		return
	}
	defer resp.Body.Close()

}
