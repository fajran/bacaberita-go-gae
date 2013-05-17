package feed

import (
	"encoding/xml"
)

func Parse(content []byte) (*Feed, error) {
	var data struct {
		Feed Feed `xml:"channel"`
	}
	err := xml.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}

	return &data.Feed, nil
}
