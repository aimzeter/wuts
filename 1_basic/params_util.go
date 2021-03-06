package params

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func GetStudentID(body io.Reader) (uint64, error) {
	return extractId("student_id", body)
}

func extractId(fieldName string, body io.Reader) (uint64, error) {
	var id uint64
	var bodyMap map[string]json.RawMessage
	var err error

	err = json.NewDecoder(body).Decode(&bodyMap)
	if err == nil {
		if bodyMap == nil {
			return id, errors.New("body cant be empty")
		}

		if bodyMap[fieldName] == nil {
			m := fmt.Sprint(fieldName, " cant be empty")
			return id, errors.New(m)
		}

		err = json.Unmarshal(bodyMap[fieldName], &id)
	}
	return id, err
}
