package responses

type ToText struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			Language                   string `json:"Language"`
			Industry                   string `json:"Industry"`
			IndustryKeyText            string `json:"IndustryKeyText"`
			ToCustomerSupplierIndustry struct {
				Deferred struct {
					URI string `json:"uri"`
				} `json:"__deferred"`
			} `json:"to_CustomerSupplierIndustry"`
		} `json:"results"`
	} `json:"d"`
}
