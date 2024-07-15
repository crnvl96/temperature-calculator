package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnTemperatureInCelsiusFahrenheitAndKelvin(t *testing.T) {
	defer gock.Off()

	gock.New("https://viacep.com.br").
		Get("/ws/01001000/json/").
		Reply(200).
		JSON(map[string]string{
			"cep":         "01001-000",
			"logradouro":  "Praça da Sé",
			"complemento": "lado ímpar",
			"bairro":      "Sé",
			"localidade":  "São Paulo",
			"uf":          "SP",
			"ibge":        "3550308",
			"gia":         "1004",
			"ddd":         "11",
			"siafi":       "7107",
		})

	gock.New("https://api.weatherapi.com").
		Get("/v1/current.json").
		Reply(200).
		JSON(map[string]interface{}{
			"location": map[string]interface{}{
				"name": "São Paulo",
			},
			"current": map[string]interface{}{
				"temp_c": 20.0,
				"temp_f": 68.0,
			},
		})

	r := http.NewServeMux()

	NewWeatherHandler(r)

	req := httptest.NewRequest("GET", "/calculate?zipcode=01001000", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	expected := `{"temp_C":20,"temp_F":68,"temp_K":293}`

	assert.Equal(t, expected, strings.TrimRight(w.Body.String(), "\n"))
}

func TestShouldReturnErrorWhenZipcodeIsInvalid(t *testing.T) {
	r := http.NewServeMux()

	NewWeatherHandler(r)

	req := httptest.NewRequest("GET", "/calculate?zipcode=123", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 422 {
		t.Errorf("Expected status code 422, got %d", w.Code)
	}

	expected := "invalid zipcode\n"

	assert.Equal(t, expected, w.Body.String())
}

func TestShouldReturnErrorWhenZipcodeIsNotFound(t *testing.T) {
	defer gock.Off()

	gock.New("https://viacep.com.br").
		Get("/ws/12345678/json/").
		Reply(404).
		JSON(map[string]bool{
			"erro": true,
		})

	r := http.NewServeMux()

	NewWeatherHandler(r)

	req := httptest.NewRequest("GET", "/calculate?zipcode=12345678", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("Expected status code 404, got %d", w.Code)
	}

	expected := "can not find zipcode\n"

	assert.Equal(t, expected, w.Body.String())
}

func TestShouldReturnErrorWhenWeatherAPIIsUnavailable(t *testing.T) {
	defer gock.Off()

	gock.New("https://viacep.com.br").
		Get("/ws/01001000/json/").
		Reply(200).
		JSON(map[string]string{
			"cep":         "01001-000",
			"logradouro":  "Praça da Sé",
			"complemento": "lado ímpar",
			"bairro":      "Sé",
			"localidade":  "São Paulo",
			"uf":          "SP",
			"ibge":        "3550308",
			"gia":         "1004",
			"ddd":         "11",
			"siafi":       "7107",
		})

	gock.New("https://api.weatherapi.com").
		Get("/v1/current.json").
		Reply(500)

	r := http.NewServeMux()

	NewWeatherHandler(r)

	req := httptest.NewRequest("GET", "/calculate?zipcode=01001000", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 500 {
		t.Errorf("Expected status code 500, got %d", w.Code)
	}
}
