package yahc_encoding_html_bindings

import (
	"github.com/zlepper/encoding-html"
	"github.com/zlepper/yahc"
	"io"
)

type ResponseParser struct {
}

func (ResponseParser) Convert(r io.Reader, v interface{}) error {
	return html.NewDecoder(r).Decode(v)
}

func init() {
	parser := ResponseParser{}
	yahc.RegisterResponseParser("text/html; charset=UTF-8", parser)
	yahc.RegisterResponseParser("text/html", parser)
	yahc.RegisterResponseParser("text/html;", parser)
}
