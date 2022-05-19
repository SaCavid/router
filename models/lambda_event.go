package models

type LambdaRequest struct {
	Headers         map[string]string `json:"headers"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
	RawPath         string            `json:"rawPath"`
	RawQueryString  string            `json:"rawQueryString"`
	RequestContext  RequestContext    `json:"requestContext"`
	RouteKey        string            `json:"routeKey"`
	Version         string            `json:"version"`
}

type RequestContext struct {
	AccountId    string `json:"accountId"`
	ApiId        string `json:"apiId"`
	DomainName   string `json:"domainName"`
	DomainPrefix string `json:"domainPrefix"`
	Http         HTTP   `json:"http"`
	RequestId    string `json:"requestId"`
	RouteKey     string `json:"routeKey"`
	Stage        string `json:"stage"`
	Time         string `json:"time"`
	TimeEpoch    int64  `json:"timeEpoch"`
}

type HTTP struct {
	Method    string `json:"method"`
	Path      string `json:"path"`
	Protocol  string `json:"protocol"`
	SourceIp  string `json:"sourceIp"`
	UserAgent string `json:"userAgent"`
}

type LambdaResponse struct {
	StatusCode      int               `json:"statusCode"`
	Headers         map[string]string `json:"headers"`
	Body            string            `json:"body"`
	Cookies         []string          `json:"cookies"`
	IsBase64Encoded bool              `json:"isBase64Encoded"`
}

func (w LambdaResponse) Set(header, value string) {
	w.Headers[header] = value
}

func (w LambdaResponse) Get(header, value string) string {
	if _, ok := w.Headers[header]; ok {
		return w.Headers[header]
	}

	return ""
}
