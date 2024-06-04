package exchange

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/sarastee/gomobile-test-assignment/internal/service/exchange/model"
	"github.com/sarastee/gomobile-test-assignment/internal/utils/validator"
	"golang.org/x/net/html/charset"
)

// InsertDailyData method which gets data from API and inserts it postgres database
func (r *Repo) InsertDailyData(ctx context.Context) error {
	url := fmt.Sprintf(FindByDateEndpoint, validator.ParseTimeNowToDate())

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	request.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",
	)

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			r.logger.Warn().Msg("Failure while closing body")
		}
	}()

	var ValCurs model.ValCurs

	decoder := xml.NewDecoder(response.Body)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&ValCurs); err != nil {
		return err
	}

	rows := make([][]interface{}, 0)
	sqlDate := time.Now().Format("2006-01-02")

	for _, item := range ValCurs.Valutes {
		value, err := strconv.ParseFloat(strings.ReplaceAll(item.Value, ",", "."), 64)
		if err != nil {
			return err
		}

		vunitValue, err := strconv.ParseFloat(strings.ReplaceAll(item.VunitRate, ",", "."), 64)
		if err != nil {
			return err
		}

		rows = append(rows, []interface{}{
			item.ID,
			sqlDate,
			item.NumCode,
			item.CharCode,
			item.Nominal,
			item.Name,
			value,
			vunitValue},
		)
	}

	_, err = r.db.DB().CopyFromContext(
		ctx,
		pgx.Identifier{currenciesTable},
		[]string{
			valuteIDColumn,
			dateColumn,
			numericCodeColumn,
			characterCodeColumn,
			nominalColumn,
			nameColumn,
			valueColumn,
			vunitRateColumn},
		pgx.CopyFromRows(rows))

	if err != nil {
		return err
	}

	return nil
}
