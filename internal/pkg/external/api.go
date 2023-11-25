package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	errors "go-api-echo/internal/pkg/helpers/errors"
	"io"
	"net/http"
)

func GetRequest(url string, query map[string]string, resStruct *interface{}) *errors.Error {
	req, err := http.NewRequest(`GET`, url, nil)
	if err != nil {
		jsonErr := *errors.InternalServerError
		jsonErr.AddError(`error creating external request`)
		return &jsonErr
	}

	if len(query) > 0 {
		queryString := req.URL.Query()

		for key, value := range query {
			queryString.Add(key, value)
		}

		req.URL.RawQuery = queryString.Encode()
	}

	fetch(req, resStruct)

	return nil
}

func PostRequest(url string, body interface{}, resStruct *interface{}) *errors.Error {
	var payloadBuf *bytes.Buffer
	if body != nil {
		payload, err := json.Marshal(body)
		if err != nil {
			jsonErr := *errors.InternalServerError
			jsonErr.AddError(`error serializing payload`)
			return &jsonErr
		}
		payloadBuf = bytes.NewBuffer(payload)
	}
	req, err := http.NewRequest(`POST`, url, payloadBuf)
	if err != nil {
		jsonErr := *errors.InternalServerError
		jsonErr.AddError(`error creating external request`)
		return &jsonErr
	}

	fetch(req, resStruct)

	return nil
}

func fetch(req *http.Request, res *interface{}) *errors.Error {
	req.Header.Set(`Content-Type`, `application/json`)
	client := &http.Client{}
	respApi, err := client.Do(req)
	if err != nil {
		fmt.Printf(`Error sending request: %s`, err)
		jsonErr := *errors.BadGatewayError
		return &jsonErr
	}

	if respApi.StatusCode == http.StatusOK || respApi.StatusCode == http.StatusCreated {
		data, err := io.ReadAll(respApi.Body)
		if err != nil {
			jsonErr := *errors.InternalServerError
			jsonErr.AddError(`error reading response body`)
			return &jsonErr
		}

		err = json.Unmarshal(data, res)
		if err != nil {
			jsonErr := *errors.InternalServerError
			jsonErr.AddError(`error parsing response`)
			return &jsonErr
		}
	} else {
		err := *errors.InternalServerError
		msg := fmt.Sprintf("error fetching request: %d", respApi.StatusCode)
		err.AddError(msg)
		return &err
	}

	return nil
}
