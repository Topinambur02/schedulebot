package utils

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"topinambur02.com/m/v2/configs"
	"topinambur02.com/m/v2/model"
)

func GetSchedule() (*model.ScheduleResponse, int) {
	authResp, statusCode := GetAuthToken()
	if authResp == nil || statusCode != 200 {
		return nil, statusCode
	}

	config := configs.NewConfig()
	url := config.URL_GET_SCHEDULE_API + authResp.Token
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln("Error creating request:", err)
		return nil, 500
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("Error making request:", err)
		return nil, 500
	}
	
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("HTTP Error: %d\n", resp.StatusCode)
		return nil, resp.StatusCode
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error reading response body:", err)
		return nil, 500
	}

	if len(body) == 0 || body[0] != '{' {
		log.Fatalln("Invalid JSON response")
		return nil, 500
	}

	var scheduleResp model.ScheduleResponse
	if err := json.Unmarshal(body, &scheduleResp); err != nil {
		log.Fatalln("Error parsing JSON:", err)
		log.Fatalf("Response body: %s\n", string(body))
		return nil, 500
	}

	return &scheduleResp, resp.StatusCode
}
