package internal

import (
	"awesomeProject19/logger"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

func addAge(name string) (int, error) {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	respAge, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		if closeErr := Body.Close(); closeErr != nil {
			logger.Log.Errorf("Error closing response: %v", closeErr)
		}
	}(respAge.Body)

	bodyAge, err := io.ReadAll(respAge.Body)
	if err != nil {
		return 0, err
	}
	age := int(gjson.Get(string(bodyAge), "age").Int())
	logger.Log.Debugf("Retrieved age %d for name %s", age, name)
	return age, nil
}

func addGender(name string) (string, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	respGender, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		if closeErr := Body.Close(); closeErr != nil {
			logger.Log.Errorf("Error closing response: %v", closeErr)
		}
	}(respGender.Body)

	bodyGender, err := io.ReadAll(respGender.Body)
	if err != nil {
		return "", err
	}
	gender := gjson.Get(string(bodyGender), "gender").String()
	logger.Log.Debugf("Retrieved gender %s for name %s", gender, name)
	return gender, nil
}

func addNationality(name string) (string, error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	respNation, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		if closeErr := Body.Close(); closeErr != nil {
			logger.Log.Errorf("Error closing response: %v", closeErr)
		}
	}(respNation.Body)

	bodyNation, err := io.ReadAll(respNation.Body)
	if err != nil {
		return "", err
	}
	nationality := gjson.Get(string(bodyNation), "country.0.country_id").String()
	logger.Log.Debugf("Retrieved nationality %s for name %s", nationality, name)
	return nationality, nil
}
