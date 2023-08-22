package test

import (
	"encoding/json"
	"gochiapp/entities"
	"gochiapp/model"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// initialize
var s *Server = StarterServer()

type StructAirportResponse struct {
	Id     string                   `json:"id"`
	Result model.CreateAirportModel `json:"result"`
}

func TestGetAllAirport(t *testing.T) {
	ctx, cancel := createTestContext()
	defer cancel()

	s.Router.Get("/api/v1/airport/", airportController.List)
	//delete all data first
	airportRepo.DeleteAll(ctx)

	//save dump-data
	airportRepo.Store(ctx, dataAirport)
	//get request
	req := httptest.NewRequest("GET", "/api/v1/airport/", nil)

	var res = httptest.NewRecorder()

	//handlerfunc
	s.Router.ServeHTTP(res, req)
	var response *http.Response = res.Result()

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
	airportRepo.DeleteAll(ctx)
}

// TODO: Change to table test
func TestGetAirportById(t *testing.T) {

	//create context
	ctx, cancel := createTestContext()
	defer cancel()
	airportRepo.DeleteAll(ctx)
	s.Router.Post("/api/v1/airport/{id}", airportController.FindById)

	airportRepo.Store(ctx, dataAirport)

	var request *http.Request = httptest.NewRequest(http.MethodPost, "/api/v1/airport/ICIKIWIR", nil)
	var recorder *httptest.ResponseRecorder = httptest.NewRecorder()

	//passing the 'fake' request and response
	s.Router.ServeHTTP(recorder, request)

	//get the result
	var result *http.Response = recorder.Result()
	var body []byte
	body, _ = io.ReadAll(result.Body)

	var bodyResponse model.ResponseWeb

	//unmarshal the body
	_ = json.Unmarshal(body, &bodyResponse)

	//header
	assert.Equal(t, 200, result.StatusCode)

	//response
	assert.Equal(t, int16(200), bodyResponse.StatusCode, "Invalid status code expected")
	assert.NotNil(t, bodyResponse.Message, "Invalid message expected")
	assert.Equal(t, "Success", bodyResponse.Status, "Invalid status expected")

	var bodyByteData []byte
	bodyByteData, _ = json.Marshal(bodyResponse.Data)

	var bodyData StructAirportResponse

	_ = json.Unmarshal(bodyByteData, &bodyData)

	assert.Equal(t, dataAirport.Id, bodyData.Id, "Invalid Id Expected")
	assert.Equal(t, dataAirport.AirportName, bodyData.Result.AirportName, "Invalid airport name expected")
	assert.Equal(t, dataAirport.Location, bodyData.Result.Location, "Invalid location expected")
	assert.Equal(t, dataAirport.LocationAcronym, bodyData.Result.LocationAcronym, "Invalid location acronym expected")

}
