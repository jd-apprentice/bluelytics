package functions

import (
	consts "dolar/src/constants"
	"dolar/src/types"
	handlers "dolar/src/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func getAccessToken(clientId string, clientSecret string) string {

	resp, err := http.PostForm(consts.InfisicalAccessToken, url.Values{
		"clientId":     {clientId},
		"clientSecret": {clientSecret},
	})

	if err != nil {
		handlers.LogFatalFmt(consts.FailedFetch, err)
	}

	defer resp.Body.Close()

	var data types.AcessToken
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		handlers.LogFatalFmt(consts.FailedFetch, err)
	}

	return data.AccessToken
}

func GetSecrets(workspaceId string, environment string) ([]types.Secret, error) {

	var accessToken = getAccessToken(os.Getenv("INFISICAL_CLIENT_ID"), os.Getenv("INFISICAL_CLIENT_SECRET"))

	var fullUrl = fmt.Sprintf("%s?workspaceId=%s&environment=%s", consts.InfisicalApiUrl, workspaceId, environment)

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return nil, handlers.ErrorMessageFmt(consts.CreateError, err)
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, handlers.ErrorMessageFmt(consts.SendError, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, handlers.ErrorMessageFmt(consts.ReadError, err)
	}

	var data types.Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, handlers.ErrorMessageFmt(consts.UnmarshalError, err)
	}

	return data.Secrets, nil
}
