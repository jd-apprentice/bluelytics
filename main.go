package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
	"github.com/olekukonko/tablewriter"
)

type AcessToken struct {
	AccessToken string
	ExpiresIn   int
	TokenType   string
}

type DolarData struct {
	ValueAvg float64 `json:"value_avg"`
}

type ApiResponse struct {
	Oficial     DolarData `json:"oficial"`
	Blue        DolarData `json:"blue"`
	OficialEuro DolarData `json:"oficial_euro"`
	BlueEuro    DolarData `json:"blue_euro"`
}

type Secret struct {
	ID          string `json:"id"`
	SecretKey   string `json:"secretKey"`
	SecretValue string `json:"secretValue"`
}

type Response struct {
	Secrets []Secret `json:"secrets"`
}

const (
	sackEmoji = "💰"
	arsLabel  = "ARS"
)

func initDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func getAccessToken(clientId string, clientSecret string) string {

	resp, err := http.PostForm("https://app.infisical.com/api/v1/auth/universal-auth/login", url.Values{
		"clientId":     {clientId},
		"clientSecret": {clientSecret},
	})

	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	defer resp.Body.Close()

	var data AcessToken
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	return data.AccessToken
}

func getSecrets(accessToken string, workspaceId string, environment string) ([]Secret, error) {
	const URL = "https://us.infisical.com/api/v3/secrets/raw"
	req, err := http.NewRequest("GET", URL+"?workspaceId="+workspaceId+"&environment="+environment, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return data.Secrets, nil
}

func getData(url string) (ApiResponse, error) {

	resp, err := http.Get(url)
	if err != nil {
		return ApiResponse{}, err
	}
	defer resp.Body.Close()

	var data ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return ApiResponse{}, err
	}

	return data, nil
}

func createTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"💱 Currency", "💸 Value"})
	table.SetColWidth(50)
	return table
}

func main() {

	initDotEnv()

	var token = getAccessToken(os.Getenv("INFISICAL_CLIENT_ID"), os.Getenv("INFISICAL_CLIENT_SECRET"))
	var projectId = os.Getenv("INFISICAL_PROJECT_ID")
	var environment = os.Getenv("INFISICAL_ENVIRONMENT")

	secrets, err := getSecrets(token, projectId, environment)

	if err != nil {
		log.Fatalf("Failed to get secrets: %v", err)
	}

	var config map[string]string

	for _, secret := range secrets {
		config = map[string]string{
			secret.SecretKey: secret.SecretValue,
		}
	}

	data, err := getData(config["API_ENDPOINT"])

	if err != nil {
		log.Fatalf("Error fetching data: %v", err)
	}

	table := createTable()
	table.Append([]string{"💵 Oficial", fmt.Sprintf("%s %.2f %s", sackEmoji, data.Oficial.ValueAvg, arsLabel)})
	table.Append([]string{"💵 Blue", fmt.Sprintf("%s %.2f %s", sackEmoji, data.Blue.ValueAvg, arsLabel)})
	table.Append([]string{"💶 Oficial_Euro", fmt.Sprintf("%s %.2f %s", sackEmoji, data.OficialEuro.ValueAvg, arsLabel)})
	table.Append([]string{"💶 Blue_euro", fmt.Sprintf("%s %.2f %s", sackEmoji, data.BlueEuro.ValueAvg, arsLabel)})

	table.Render()
}
