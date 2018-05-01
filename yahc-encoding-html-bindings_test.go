package yahc_encoding_html_bindings

import (
	"github.com/stretchr/testify/assert"
	"github.com/zlepper/yahc"
	"testing"
)

func TestCanParseResponse(t *testing.T) {
	type ExampleDotCom struct {
		Title string `css:"h1"`
	}

	var example ExampleDotCom
	yahc.Get("http://example.com/", yahc.NoOptions, &example)

	assert.Equal(t, "Example Domain", example.Title)
}
