package exchange

import (
	"errors"
	"net/http"
	"strings"

	"github.com/sarastee/gomobile-test-assignment/internal/api"
	"github.com/sarastee/gomobile-test-assignment/internal/service"
	"github.com/sarastee/gomobile-test-assignment/internal/utils/response"
	"github.com/sarastee/gomobile-test-assignment/internal/utils/validator"
)

var globalAPICodes map[string]bool

// GetExchangeRate API method which handles GET /currency endpoint
func (i *Implementation) GetExchangeRate(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := r.Body.Close()
		if err != nil {
			i.logger.Warn().Msg(err.Error())
		}
	}()

	err := validator.CheckContentType(r)
	if err != nil {
		i.logger.Info().Msg(err.Error())
		response.SendError(w, http.StatusBadRequest, err, i.logger)
		return
	}

	date := r.URL.Query().Get(DateParam)

	switch {
	case date == "":
		date = validator.ParseTimeNowToDate()
	case !validator.ValidateDateQueryParam(date):
		i.logger.Info().Msg(api.ErrInvalidDateFormat.Error())
		response.SendError(w, http.StatusBadRequest, api.ErrInvalidDateFormat, i.logger)
		return
	}

	val := r.URL.Query().Get(ValParam)
	val = strings.ToUpper(val)

	if globalAPICodes == nil {
		globalAPICodes, err = i.exchangeService.GetCurrenciesFromAPI()
		if err != nil {
			i.logger.Info().Msg(err.Error())
			response.SendError(w, http.StatusBadRequest, err, i.logger)
			return
		}
	}

	switch {
	case val == "":
		i.logger.Info().Msg(api.ErrInvalidValFormat.Error())
		response.SendError(w, http.StatusBadRequest, api.ErrInvalidValFormat, i.logger)
		return
	case val != "":
		if !globalAPICodes[val] {
			i.logger.Info().Msg(api.ErrInvalidValFormat.Error())
			response.SendError(w, http.StatusBadRequest, api.ErrInvalidValFormat, i.logger)
			return
		}
	}

	exchangeRate, err := i.exchangeService.GetExchangeRateFromAPI(r.Context(), val, date)
	if err != nil {
		if errors.Is(err, service.ErrNoDataFound) {
			i.logger.Info().Msg(err.Error())
			response.SendError(w, http.StatusNotFound, err, i.logger)
			return
		}
		i.logger.Info().Msg(err.Error())
		response.SendError(w, http.StatusBadRequest, err, i.logger)
		return
	}

	response.SendStatus(w, http.StatusOK, exchangeRate, i.logger)
}