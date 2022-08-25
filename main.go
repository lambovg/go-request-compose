package main

func main() {

	var client = new(Request)
	client.hostname = "d2kgi8nio2h9bn.cloudfront.net"
	client.protocol = "https"
	client.path = "hello-world.json"
	client.url = "https://d2kgi8nio2h9bn.cloudfront.net/hello-world.json"

	Get{*client}.Request()

	client.url = "https://d2kgi8nio2h9bn.cloudfront.net/ping.json"
	Get{*client}.Request()

	/*
		extend := Get{url: "https://d2kgi8nio2h9bn.cloudfront.net"}
		Get{url: extend.url + "/ping.json"}.Request()
		Get{url: extend.url + "/hello-world.json"}.Request()
	*/
}
