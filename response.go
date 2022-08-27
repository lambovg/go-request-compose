package main

type Response struct {
	Body string
	Err  error
}

func (r Response) Response(log *BuiltinLogger) {
	log.Printf(r.Body)

	if r.Err != nil {
		log.Fatalln(r.Err)
	}
}
