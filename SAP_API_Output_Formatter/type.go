package sap_api_output_formatter

type CustomerSupplierIndustryReads struct {
	ConnectionKey                string `json:"connection_key"`
	Result                       bool   `json:"result"`
	RedisKey                     string `json:"redis_key"`
	Filepath                     string `json:"filepath"`
	Product                      string `json:"Product"`
	APISchema                    string `json:"api_schema"`
	CustomerSupplierIndustryCode string `json:"customer_supplier_industry_code"`
	Deleted                      string `json:"deleted"`
}

type CustomerSupplierIndustry struct {
	Industry string `json:"Industry"`
	ToText   string `json:"to_Text"`
}

type Text struct {
	Language        string `json:"Language"`
	Industry        string `json:"Industry"`
	IndustryKeyText string `json:"IndustryKeyText"`
}

type ToText struct {
	Language        string `json:"Language"`
	Industry        string `json:"Industry"`
	IndustryKeyText string `json:"IndustryKeyText"`
}
