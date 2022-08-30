package request

type Post struct {
	Params Request
	Url    string
	Body   string
}

func (r Post) Request() {
	//TODO: implementation
}

func (p Params) Post() {
	//TODO: implementation
}
