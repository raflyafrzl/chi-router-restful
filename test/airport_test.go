package test

import (
	"encoding/json"
	"gochiapp/entities"
	"gochiapp/model"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllAirport(t *testing.T) {
	ctx, cancel := createTestContext()
	defer cancel()
	//delete all data first
	repository.DeleteAll(ctx)

	s := StarterServer()
	s.AllHandlers()

	//save dump-data
	repository.Store(ctx, dataAirport)
	//get request
	req := httptest.NewRequest("GET", "/api/v1/airport/", nil)

	var res = httptest.NewRecorder()

	//handlerfunc
	s.Router.ServeHTTP(res, req)
	response := res.Result()

	//read the result from response
	body, _ := io.ReadAll(response.Body)

	var result model.ResponseWeb

	err := json.Unmarshal(body, &result)

	if err != nil {
		t.Fatal("Error occured")
	}

	assert.Equal(t, int16(200), result.StatusCode, "Invalid status code expected")
	assert.Equal(t, "Success", result.Status, "Invalid status expected")

	//get response data
	rawdata, _ := json.Marshal(result.Data)
	//unmarshal data
	var dataResult []entities.Airport

	_ = json.Unmarshal(rawdata, &dataResult)

	assert.Equal(t, dataAirport.AirportCode, dataResult[0].AirportCode)
	assert.Equal(t, dataAirport.AirportName, dataResult[0].AirportName)
	assert.Equal(t, dataAirport.Id, dataResult[0].Id)
	assert.Equal(t, dataAirport.Location, dataResult[0].Location)
	assert.Equal(t, dataAirport.LocationAcronym, dataResult[0].LocationAcronym)
	repository.DeleteAll(ctx)
}

func TestGetAirportById(t *testing.T) {

}
