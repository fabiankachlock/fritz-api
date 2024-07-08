package transformer

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/go-shiori/dom"
	"golang.org/x/net/html"
)

var (
	NetCntDataRegex = regexp.MustCompile(`const data\s*=\s*(\{.*\});`)
)

func NetCntToJson(rawContent []byte) ([]byte, error) {
	doc, err := html.Parse(bytes.NewReader(rawContent))
	if err != nil {
		return nil, err
	}

	script := dom.QuerySelector(doc, "script")
	data := NetCntDataRegex.FindStringSubmatch(dom.InnerText(script))
	if len(data) < 2 {
		return []byte{}, fmt.Errorf(("no data found in script"))
	}
	return []byte(data[1]), nil
}
