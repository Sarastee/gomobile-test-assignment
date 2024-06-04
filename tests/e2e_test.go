package tests

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type testTable struct {
	Name           string
	Method         string
	Route          string
	Body           string
	Headers        [][2]string
	ExpectedStatus int
	ParseResp      bool
	RespBody       interface{}
}

func createRequest(method, route, body string, headers [][2]string) (*http.Request, error) {
	host := os.Getenv("HTTP_HOST")
	if len(host) == 0 {
		return nil, errors.New("http host not found")
	}

	port := os.Getenv("HTTP_PORT")
	if len(port) == 0 {
		return nil, errors.New("http port not found")
	}

	url := fmt.Sprintf("http://%s:%s%s", host, port, route)

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	for _, header := range headers {
		req.Header.Add(header[0], header[1])
	}

	return req, nil
}

func sendRequest(t *testing.T, client *http.Client, req *http.Request, expectedStatus int, parseResp bool, respBody interface{}) {
	t.Helper()
	resp, err := client.Do(req)
	require.NoError(t, err)

	require.Equal(t, expectedStatus, resp.StatusCode)
	if parseResp {
		err = json.NewDecoder(resp.Body).Decode(respBody)
		require.NoError(t, err)
	}
	defer resp.Body.Close() //nolint
}

func TestGetExchangeRate(t *testing.T) {
	client := http.Client{}
	var exchangeRate json.RawMessage

	tests := []testTable{
		{
			Name:           "Get Exchange Rate no Content-type header",
			Method:         http.MethodGet,
			Route:          "/currency?date=05/10/2015&val=EUR",
			Body:           "",
			Headers:        [][2]string{},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get Exchange Rate incorrect Content-type header",
			Method:         http.MethodGet,
			Route:          "/currency?date=05/10/2015&val=EUR",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "invalidType"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get Exchange Rate incorrect data format",
			Method:         http.MethodGet,
			Route:          "/currency?date=05.10.2015&val=EUR",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get Exchange Rate incorrect data format (incorrect month number)",
			Method:         http.MethodGet,
			Route:          "/currency?date=31/31/2015&val=EUR",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get Exchange Rate incorrect val format",
			Method:         http.MethodGet,
			Route:          "/currency?date=15/05/2015&val=EURO",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get Exchange Rate incorrect val format",
			Method:         http.MethodGet,
			Route:          "/currency?date=15/05/2015&val=евро",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusBadRequest,
			ParseResp:      false,
			RespBody:       nil,
		},
		{
			Name:           "Get Exchange Rate Success (with data and val)",
			Method:         http.MethodGet,
			Route:          "/currency?date=15/05/2015&val=EUR",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusOK,
			ParseResp:      true,
			RespBody:       &exchangeRate,
		},
		{
			Name:           "Get Exchange Rate Success (with data and val lowercase format)",
			Method:         http.MethodGet,
			Route:          "/currency?date=15/05/2015&val=eur",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusOK,
			ParseResp:      true,
			RespBody:       &exchangeRate,
		},
		{
			Name:           "Get Exchange Rate Success (without date)",
			Method:         http.MethodGet,
			Route:          "/currency?val=EUR",
			Body:           "",
			Headers:        [][2]string{{"Content-Type", "application/json"}},
			ExpectedStatus: http.StatusOK,
			ParseResp:      true,
			RespBody:       &exchangeRate,
		},
	}

	for _, tc := range tests {
		t.Log(tc.Name)

		req, err := createRequest(tc.Method, tc.Route, tc.Body, tc.Headers)
		if err != nil {
			t.Log(err.Error())
		}
		require.NoError(t, err)

		sendRequest(t, &client, req, tc.ExpectedStatus, tc.ParseResp, tc.RespBody)
	}
}
