package user

import (
	"bytes"
	"encoding/json"
	"io"
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

func (c *Credentials) LoginHandler() bool {
	return SendLoginRequest(*c, os.Getenv("SERVER_LOGIN_API_URL"))
}
func SendLoginRequest(credentials Credentials, url string) bool {
	jsonData, err := json.Marshal(credentials)
	if err != nil {
		log.Fatalln("error marshalling credentials")
		return false
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln("error creating POST request")
		return false
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("error sending request")
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("error reading response")
		return false
	}
	//responce:
	var result map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error parsing JSON response: %s", err)
		return false
	}

	if val, ok := result["bool"]; ok && val == "true" {
		return true
	}
	return false

}
