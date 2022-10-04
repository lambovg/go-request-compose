package request

import (
	"bytes"
	"fmt"
)

// BuildUrl .
func (p Params) BuildUrl() string {
	var b bytes.Buffer
	b.WriteString(p.Protocol)
	b.WriteString("://")
	b.WriteString(p.Hostname)

	if p.Port != 0 {
		b.WriteString(":")
		b.WriteString(fmt.Sprintf("%d", p.Port))
	}

	b.WriteString(p.Path)

	if p.QueryString != "" {
		b.WriteString("?")
		b.WriteString(p.QueryString)
	}

	return b.String()
}