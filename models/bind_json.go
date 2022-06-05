package models

import "encoding/json"

func (r LambdaRequest) BindJson(d any) error {

	err := json.Unmarshal([]byte(r.Body), &d)
	if err != nil {
		return err
	}

	return nil
}
