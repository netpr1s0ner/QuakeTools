package model

// httpInfo represents the details of an HTTP service.
type httpInfo struct {
	HTMLHash        string      `json:"html_hash"`
	Title           string      `json:"title"`
	Host            string      `json:"host"`
	Path            string      `json:"path"`
	Body            string      `json:"body"`
	Robots          string      `json:"robots"`
	RobotsHash      string      `json:"robots_hash"`
	Sitemap         string      `json:"sitemap"`
	SitemapHash     string      `json:"sitemap_hash"`
	Favicon         faviconInfo `json:"favicon"`
	Server          string      `json:"server"`
	XPoweredBy      string      `json:"x_powered_by"`
	MetaKeywords    string      `json:"meta_keywords"`
	SecurityText    string      `json:"security_text"`
	StatusCode      int         `json:"status_code"`
	ResponseHeaders string      `json:"response_headers"`
}

// faviconInfo holds information about the favicon.
type faviconInfo struct {
	Hash     string `json:"hash"`
	Location string `json:"location"`
	Data     string `json:"data"`
}

// serviceData represents each individual service data within ServiceInfo.
type serviceData struct {
	Time      string        `json:"time"` // Time, as string, will be parsed manually
	Transport string        `json:"transport"`
	Service   serviceDetail `json:"service"`
	//Components []componentDetail `json:"components"`
	//Location   locationInfo      `json:"location"`
	Hostname string `json:"hostname"`
	//Org        string            `json:"org"`
	IsIpv6 bool   `json:"is_ipv6"`
	IP     string `json:"ip"`
	Port   int    `json:"port"`
}

// serviceDetail holds the HTTP service details.
type serviceDetail struct {
	HTTP     httpInfo `json:"http"`
	Version  string   `json:"version"`
	Name     string   `json:"name"`
	Product  *string  `json:"product"`
	Banner   string   `json:"banner"`
	Response string   `json:"response"`
}

// componentDetail holds the information about individual components.
type componentDetail struct {
	ProductLevel   string   `json:"product_level"`
	ProductType    []string `json:"product_type"`
	ProductVendor  string   `json:"product_vendor"`
	ProductNameCN  string   `json:"product_name_cn"`
	ProductNameEN  string   `json:"product_name_en"`
	ID             string   `json:"id"`
	ProductCatalog []string `json:"product_catalog"`
	Version        string   `json:"version"`
}

// locationInfo contains geographical and location details.
type locationInfo struct {
	//ProvinceCn  string    `json:"province_cn"`
	//CountryCode string    `json:"country_code"`
	//Gps         []float64 `json:"gps"`
}

// paginationMeta holds metadata for pagination.
type paginationMeta struct {
	PaginationID string `json:"pagination_id"`
}

// ServiceInfo represents the service information response.
type ServiceInfo struct {
	Code int `json:"code"`
	//Message string         `json:"message"`
	Data []serviceData  `json:"data"`
	Meta paginationMeta `json:"meta"`
}
