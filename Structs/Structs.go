package structs

import (
	"encoding/xml"
)

type GoCurrency struct {
	CrossOrder      int
	Kod             string
	CurrencyCode    string
	Unit            int
	CurrencyName    string
	ForexBuying     float64
	ForexSelling    float64
	BanknoteBuying  float64
	BanknoteSelling float64
	CrossRateUSD    float64
	CrossRateOther  float64
}

type GOTarih_Date struct {
	Tarih         string
	Date          string
	Bulten_No     string
	GoCurrencyies []GoCurrency
}

type XMLTarih_Date struct {
	Tarih          string        `xml:"Tarih,attr"`
	Date           string        `xml:"Date,attr"`
	Bulten_No      string        `xml:"Bulten_No,attr"`
	XmlCurrencyies []XMLCurrency `xml:"Currency"`
}

type XMLCurrency struct {
	XMLName         xml.Name `xml:"Currency"`
	CrossOrder      string   `xml:"CrossOrder,attr"`
	Kod             string   `xml:"Kod,attr`
	CurrencyCode    string   `xml:"CurrencyCode,attr`
	Unit            string   `xml:"Unit"`
	CurrencyName    string   `xml:"CurrencyName"`
	ForexBuying     string   `xml:"ForexBuying"`
	ForexSelling    string   `xml:"ForexSelling"`
	BanknoteBuying  string   `xml:"BanknoteBuying"`
	BanknoteSelling string   `xml:"BanknoteSelling"`
	CrossRateUSD    string   `xml:"CrossRateUSD"`
	CrossRateOther  string   `xml:"CrossRateOther"`
}
