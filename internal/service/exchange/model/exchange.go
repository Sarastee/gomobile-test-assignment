package model

import "encoding/xml"

// CurrencyAPICodes ...
type CurrencyAPICodes struct {
	XMLName  xml.Name               `xml:"Valuta"`
	Elements []CurrencyAPICodesItem `xml:"Item"`
}

// CurrencyAPICodesItem ...
type CurrencyAPICodesItem struct {
	XMLName     xml.Name `xml:"Item"`
	ID          string   `xml:"ID,attr"`
	Name        string   `xml:"Name"`
	EngName     string   `xml:"EngName"`
	Nominal     int      `xml:"Nominal"`
	ISONumCode  int      `xml:"ISO_Num_Code"`
	ISOCharCode string   `xml:"ISO_Char_Code"`
}

// ValCurs ...
type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []Valute `xml:"Valute"`
}

// Valute ...
type Valute struct {
	XMLName   xml.Name `xml:"Valute"`
	ID        string   `xml:"ID,attr"`
	NumCode   string   `xml:"NumCode"`
	CharCode  string   `xml:"CharCode"`
	Nominal   string   `xml:"Nominal"`
	Name      string   `xml:"Name"`
	Value     string   `xml:"Value"`
	VunitRate string   `xml:"VunitRate"`
}

// ValuteOutput struct
type ValuteOutput struct {
	CharCode string
	Date     string
	Rate     float64
}
