package models

func (w LambdaResponse) Set(header, value string) {
	w.Headers[header] = value
}

func (w LambdaResponse) Get(header, value string) string {
	if _, ok := w.Headers[header]; ok {
		return w.Headers[header]
	}

	return ""
}
