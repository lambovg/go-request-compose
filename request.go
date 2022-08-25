package main

type Request struct {
	url      string
	hostname string
	protocol string
	path     string
}

type Get struct {
	params Request
}

type Post struct {
	params Request
}
