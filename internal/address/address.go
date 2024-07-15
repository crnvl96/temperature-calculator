package address

import (
	"encoding/json"
	"net/http"
)

type Address struct {
	City string `json:"localidade"`
}

func GetAddressCity(zipcode string) (string, error) {
	req, err := http.NewRequest("GET", "https://viacep.com.br/ws/"+zipcode+"/json/", nil)
	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var address Address
	err = json.NewDecoder(resp.Body).Decode(&address)
	if err != nil {
		return "", err
	}

	return address.City, nil
}
