package request

type Post struct {
	Params Request
	Url    string
	Body   string
}

func (r Post) Request() {
	//TODO: implementation
}

func (r Params) Post() {
	//TODO: implementation
}
