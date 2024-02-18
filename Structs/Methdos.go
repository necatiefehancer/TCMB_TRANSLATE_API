package structs

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetXmlCurrencyies() (XMLTarih_Date, error) {
	const requestUrl string = "http://www.tcmb.gov.tr/kurlar/today.xml"
	var xmlTarih_Date = new(XMLTarih_Date)
	response, err := http.Get(requestUrl)
	if err != nil {
		log.Fatal("Fatal Error ", err.Error())
	} else {
		if response.StatusCode != http.StatusNotFound {
			defer response.Body.Close()
			xmlDecoder := xml.NewDecoder(response.Body)
			err := xmlDecoder.Decode(&xmlTarih_Date)
			if err != nil {
				log.Fatal("Decode Error ", err.Error())
			} else {
				return *xmlTarih_Date, nil
			}
		}
	}
	return XMLTarih_Date{}, errors.New("RequsetFatal")
}

func GetTranslateJSON(filename string, value interface{}) {
	jsonFile, err := os.Create(filename + ".json")
	defer jsonFile.Close()
	if err != nil {
		log.Fatal("Create Json File error ! ", err.Error())
	} else {
		jsonEncoder := json.NewEncoder(jsonFile)
		err := jsonEncoder.Encode(value)
		if err != nil {
			log.Fatal("Translate Json Encode Error ! ", err.Error())
		} else {
			log.Println("Created File Success !")
		}
	}
}

func TranslateXMLtoGoStruct(xmldata XMLTarih_Date) GOTarih_Date {
	goCurrency := GoCurrency{}
	goTarihDate := GOTarih_Date{}

	goTarihDate.Date = xmldata.Date
	goTarihDate.Tarih = xmldata.Tarih
	goTarihDate.Bulten_No = xmldata.Bulten_No

	for _, v := range xmldata.XmlCurrencyies {
		goCurrency.Unit, _ = strconv.Atoi(v.Unit)
		goCurrency.Kod = v.Unit
		goCurrency.BanknoteBuying, _ = strconv.ParseFloat(v.BanknoteBuying, 64)
		goCurrency.BanknoteSelling, _ = strconv.ParseFloat(v.BanknoteSelling, 64)
		goCurrency.CrossOrder, _ = strconv.Atoi(v.CrossOrder)
		goCurrency.CrossRateOther, _ = strconv.ParseFloat(v.CrossRateOther, 64)
		goCurrency.CrossRateUSD, _ = strconv.ParseFloat(v.CrossRateUSD, 64)
		goCurrency.CurrencyCode = v.CurrencyCode
		goCurrency.CurrencyName = v.CurrencyName
		goCurrency.ForexBuying, _ = strconv.ParseFloat(v.ForexBuying, 64)
		goCurrency.ForexSelling, _ = strconv.ParseFloat(v.ForexSelling, 64)
		goTarihDate.GoCurrencyies = append(goTarihDate.GoCurrencyies, goCurrency)
	}

	return goTarihDate
}
