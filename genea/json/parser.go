package json

import (
	"strings"
	"time"

	"github.com/json-iterator/go"
	"gopkg.in/go-playground/validator.v9"
)

var val *validator.Validate

const DateFormat = "2006-01-02"

func init() {
	val = validator.New()
}

func Parse(content []byte, data interface{}) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(content, data)
	if err != nil {
		return err
	}
	return val.Struct(data)
}

func (dt *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	parsed, err := time.Parse(DateFormat, s)
	if err != nil {
		return err
	}
	*dt = Date(parsed)
	return nil
}

func (dt *Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(*dt).Format("\"" + DateFormat + "\"")), nil
}
