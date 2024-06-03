package exchange

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/sarastee/gomobile-test-assignment/internal/service"
	"github.com/sarastee/gomobile-test-assignment/internal/service/exchange/model"
	"golang.org/x/net/html/charset"
)

// GetExchangeRateFromAPI method which gets exchange rates from cbr.ru API
func (s *Service) GetExchangeRateFromAPI(_ context.Context, val string, date string) (json.RawMessage, error) {
	url := fmt.Sprintf(FindByDateEndpoint, date)

	request, err := http.NewRequest("GET", url, nil)
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

	var ValCurs model.ValCurs

	decoder := xml.NewDecoder(response.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&ValCurs); err != nil {
		return nil, err
	}

	var ValOutput model.ValuteOutput

	for _, item := range ValCurs.Valutes {
		if item.CharCode == val {
			ValOutput.Rate, _ = strconv.ParseFloat(strings.ReplaceAll(item.VunitRate, ",", "."), 64)
			ValOutput.Date = date
			ValOutput.CharCode = item.CharCode

			break
		}
	}

	if ValOutput.CharCode == "" {
		return nil, service.ErrNoDataFound
	}

	ValOutputDecoded, err := json.Marshal(ValOutput)
	if err != nil {
		return nil, err
	}

	return ValOutputDecoded, nil
}
