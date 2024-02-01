package use_cases

import (
	"net/http"
)

const url = "https://testgab.free.beeceptor.com/"

func GetOneSynced(tool, id, token string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url+tool+"/"+id, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("token", token)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetAllSynced(token, tool string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url+tool, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("token", token)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
