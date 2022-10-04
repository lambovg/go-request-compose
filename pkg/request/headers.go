package request

import "strings"

// Headers .
type Headers struct {
	Add Header
	Set Header
}

// Header .
type Header map[string][]string

// AttachHeaders .
func (rq Request) AttachHeaders(p *Params) *Request {
	// set / override existing
	for key, val := range p.Headers.Set {
		rq.Header.Set(key, strings.Join(val, ","))
	}

	// add / extend definition of existing
	for key, val := range p.Headers.Add {
		rq.Header.Add(key, strings.Join(val, ","))
	}

	return &rq
}
