package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-customer-supplier-industry-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetCustomerSupplierIndustry(industry, language, industryKeyText string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "CustomerSupplierIndustry":
			func() {
				c.CustomerSupplierIndustry(industry)
				wg.Done()
			}()
		case "Text":
			func() {
				c.Text(industry, language, industryKeyText)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) CustomerSupplierIndustry(industry string) {
	customerSupplierIndustryData, err := c.callCustomerSupplierIndustrySrvAPIRequirementCustomerSupplierIndustry("A_CustomerSupplierIndustry", industry)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(customerSupplierIndustryData)
	}

	textData, err := c.callToText(customerSupplierIndustryData[0].ToText)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(textData)
	}
	return
}

func (c *SAPAPICaller) callCustomerSupplierIndustrySrvAPIRequirementCustomerSupplierIndustry(api, industry string) ([]sap_api_output_formatter.CustomerSupplierIndustry, error) {
	url := strings.Join([]string{c.baseURL, "API_CUSTOMERSUPPLIERINDUSTRY_SRV", api}, "/")
	param := c.getQueryWithCustomerSupplierIndustry(map[string]string{}, industry)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCustomerSupplierIndustry(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToText(url string) ([]sap_api_output_formatter.ToText, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Text(industry, language, industryKeyText string) {
	data, err := c.callCustomerSupplierIndustrySrvAPIRequirementText("A_CustomerSupplierIndustryText", industry, language, industryKeyText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callCustomerSupplierIndustrySrvAPIRequirementText(api, industry, language, industryKeyText string) ([]sap_api_output_formatter.Text, error) {
	url := strings.Join([]string{c.baseURL, "API_CUSTOMERSUPPLIERINDUSTRY_SRV", api}, "/")

	param := c.getQueryWithText(map[string]string{}, industry, language, industryKeyText)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToText(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithCustomerSupplierIndustry(params map[string]string, industry string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Industry eq '%s'", industry)
	return params
}

func (c *SAPAPICaller) getQueryWithText(params map[string]string, industry, language, industryKeyText string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("Industry eq '%s' and Language eq '%s' and IndustryKeyText eq '%s'", industry, language, industryKeyText)
	return params
}
