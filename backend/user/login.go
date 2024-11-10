package user

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Credentials) LoginHandler() (bool, string) {
	log.Printf("Attempting login with username: %s and password: %s", c.Username, c.Password)
	token, success := SendRequest(*c, os.Getenv("SERVER_LOGIN_API_URL"))
	return success, token
}

func SendRequest(credentials Credentials, url string) (string, bool) {
	jsonData, err := json.Marshal(credentials)
	if err != nil {
		log.Fatalln("Error marshalling credentials:", err)
		return "", false
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalln("Error creating POST request:", err)
		return "", false
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error sending request:", err)
		return "", false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading response:", err)
		return "", false
	}

	log.Printf("Response from external API: %s", body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error parsing JSON response: %s", err)
		return "", false
	}

	// Извлечение токена
	token, tokenExists := result["token"].(string)

	// Проверка "bool" поля
	success := false
	if val, ok := result["bool"]; ok {
		if valStr, isString := val.(string); isString && valStr == "true" {
			success = true
		}
		if valBool, isBool := val.(bool); isBool && valBool {
			success = true
		}
	}

	if success && tokenExists {
		return token, true
	}

	return "", false
}
