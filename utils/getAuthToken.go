package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"

	"topinambur02.com/m/v2/configs"
	"topinambur02.com/m/v2/model"
)

func GetAuthToken() (*model.AuthResponse, int) {
	formData := url.Values{}
	config := configs.NewConfig()
	ulogin := config.ULOGIN
	upassword := config.UPASSWORD

	formData.Set("ulogin", ulogin)
	formData.Set("upassword", upassword)

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://e.mospolytech.ru/old/lk_api.php", bytes.NewBufferString(formData.Encode()))

	if err != nil {
		log.Println("Error creating request: ", err)
		return nil, 500
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Origin", "https://e.mospolytech.ru")
	req.Header.Set("Referer", "https://e.mospolytech.ru/")

	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error making POST request: ", err)
		return nil, 500
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body: ", err)
		return nil, 500
	}

	var authResp model.AuthResponse
	
	if err := json.Unmarshal(body, &authResp); err != nil {
		log.Println("Error parsing JSON: ", err)
		return nil, 500
	}

	return &authResp, resp.StatusCode
}
