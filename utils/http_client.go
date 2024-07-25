package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PeopleInfo struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

func GetPeopleInfo(passportSerie, passportNumber string) (*PeopleInfo, error) {
	url := fmt.Sprintf("http://external-api/info?passportSerie=%s&passportNumber=%s", passportSerie, passportNumber)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	var info PeopleInfo
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}

	return &info, nil
}
