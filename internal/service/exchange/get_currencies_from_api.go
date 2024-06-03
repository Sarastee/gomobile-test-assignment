package exchange

import (
	"encoding/xml"
	"net/http"

	"github.com/sarastee/gomobile-test-assignment/internal/service/exchange/model"
	"golang.org/x/net/html/charset"
)

// GetCurrenciesFromAPI method which gets currency codes from cbr.ru API
func (s *Service) GetCurrenciesFromAPI() (map[string]bool, error) {
	currencyCodesMap := make(map[string]bool)

	request, err := http.NewRequest("GET", CurrencyCodesEndpoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",
	)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			s.logger.Warn().Msg("Failure while closing body")
		}
	}()

	var codes model.CurrencyAPICodes
	decoder := xml.NewDecoder(response.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&codes); err != nil {
		return nil, err
	}

	for _, item := range codes.Elements {
		currencyCodesMap[item.ISOCharCode] = true
	}

	return currencyCodesMap, nil
}
