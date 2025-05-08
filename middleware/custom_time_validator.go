package middleware

import (
	"time"
	"encoding/json"
	"fmt"
	"strings"
)


var isoLayouts = []string{
	"2006-01-02T15:04:05.000Z07:00",
	"2006-01-02T15:04:05Z07:00",    
}


type CustomTime struct {
	time.Time
	ParseError error `json:"-"`
}


func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	for _, layout := range isoLayouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			ct.Time = t
			ct.ParseError = nil
			return nil
		}
		ct.ParseError = fmt.Errorf("formato inválido")
	}
	return nil
}


func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(ct.Format(isoLayouts[0]))
}
