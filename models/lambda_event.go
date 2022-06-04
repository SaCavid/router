package models

type LambdaRequest struct {
	Version               string            `json:"version"`
	RouteKey              string            `json:"routeKey"`
	RawPath               string            `json:"rawPath"`
	RawQueryString        string            `json:"rawQueryString"`
	Cookies               []string          `json:"cookies"`
	Headers               map[string]string `json:"headers"`
	QueryStringParameters map[string]string `json:"queryStringParameters"`
	RequestContext        `json:"requestContext"`
	Body                  string      `json:"body"`
	PathParameters        interface{} `json:"pathParameters"`
	IsBase64Encoded       bool        `json:"isBase64Encoded"`
	StageVariables        interface{} `json:"stageVariables"`

	// Vars internal usage - for variables in url example: /path/{variable}
	Vars map[string]string
}

type RequestContext struct {
	AccountID      string      `json:"accountId"`
	APIID          string      `json:"apiId"`
	Authentication interface{} `json:"authentication"`
	Authorizer     struct {
		Iam struct {
			AccessKey       string      `json:"accessKey"`
			AccountID       string      `json:"accountId"`
			CallerID        string      `json:"callerId"`
			CognitoIdentity interface{} `json:"cognitoIdentity"`
			PrincipalOrgID  interface{} `json:"principalOrgId"`
			UserArn         string      `json:"userArn"`
			UserID          string      `json:"userId"`
		} `json:"iam"`
	} `json:"authorizer"`
	DomainName   string `json:"domainName"`
	DomainPrefix string `json:"domainPrefix"`
	HTTP         struct {
		Method    string `json:"method"`
		Path      string `json:"path"`
		Protocol  string `json:"protocol"`
		SourceIP  string `json:"sourceIp"`
		UserAgent string `json:"userAgent"`
	} `json:"http"`
	RequestID string `json:"requestId"`
	RouteKey  string `json:"routeKey"`
	Stage     string `json:"stage"`
	Time      string `json:"time"`
	TimeEpoch int64  `json:"timeEpoch"`
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
