package handler

import "encoding/json"

func errorJSON(msg string) []byte {
	parsedError := struct {
		Message string `json:"message"`
	}{
		Message: msg,
	}

	r, err := json.Marshal(parsedError)

	if err != nil {
		return []byte(err.Error())
	}

	return []byte(r)
}
