package functions

import (
	"dolar/src/types"
	"encoding/json"
	"net/http"
)

func GetData(url string) (types.ApiResponse, error) {

	resp, err := http.Get(url)
	if err != nil {
		return types.ApiResponse{}, err
	}
	defer resp.Body.Close()

	var data types.ApiResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return types.ApiResponse{}, err
	}

	return data, nil
}
