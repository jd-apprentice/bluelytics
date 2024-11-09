package tests

import (
	"dolar/src/functions"
	"dolar/src/types"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetData(t *testing.T) {
	var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(types.ApiResponse{
			Oficial:     types.DolarData{ValueAvg: 100.0},
			Blue:        types.DolarData{ValueAvg: 120.0},
			OficialEuro: types.DolarData{ValueAvg: 90.0},
			BlueEuro:    types.DolarData{ValueAvg: 110.0},
		})
	})

	srv := httptest.NewServer(handler)
	defer srv.Close()

	data, err := functions.GetData(srv.URL)
	if err != nil {
		t.Fatal(err)
	}

	expected := types.ApiResponse{
		Oficial:     types.DolarData{ValueAvg: 100.0},
		Blue:        types.DolarData{ValueAvg: 120.0},
		OficialEuro: types.DolarData{ValueAvg: 90.0},
		BlueEuro:    types.DolarData{ValueAvg: 110.0},
	}

	if !reflect.DeepEqual(data, expected) {
		t.Errorf("expected %+v, got %+v", expected, data)
	}
}
