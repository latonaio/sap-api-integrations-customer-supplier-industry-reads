package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-customer-supplier-industry-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToCustomerSupplierIndustry(raw []byte, l *logger.Logger) ([]CustomerSupplierIndustry, error) {
	pm := &responses.CustomerSupplierIndustry{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CustomerSupplierIndustry. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	customerSupplierIndustry := make([]CustomerSupplierIndustry, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		customerSupplierIndustry = append(customerSupplierIndustry, CustomerSupplierIndustry{
			Industry: data.Industry,
			ToText:   data.ToText.Deferred.URI,
		})
	}

	return customerSupplierIndustry, nil
}

func ConvertToText(raw []byte, l *logger.Logger) ([]Text, error) {
	pm := &responses.Text{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Text. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	text := make([]Text, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		text = append(text, Text{
			Language:        data.Language,
			Industry:        data.Industry,
			IndustryKeyText: data.IndustryKeyText,
		})
	}

	return text, nil
}

func ConvertToToText(raw []byte, l *logger.Logger) ([]ToText, error) {
	pm := &responses.ToText{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToText. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	toText := make([]ToText, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		toText = append(toText, ToText{
			Language:        data.Language,
			Industry:        data.Industry,
			IndustryKeyText: data.IndustryKeyText,
		})
	}

	return toText, nil
}
