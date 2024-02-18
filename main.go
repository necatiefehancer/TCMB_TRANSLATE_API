package main

import (
	"log"
	structs "modules/Structs"
)

func main() {
	xmlTarih_Date, err := structs.GetXmlCurrencyies()
	if err != nil {
		log.Fatal(err.Error())
	} else {
		goTarihDate := structs.TranslateXMLtoGoStruct(xmlTarih_Date)
		structs.GetTranslateJSON("tcmbApi", &goTarihDate)
	}
}
