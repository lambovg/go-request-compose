package response

import(
	"github.com/lambovg/go-request-compose/pkg/logger"
)

type Response struct {
	Body string
	Err  error
}

func (r Response) Response(log *logger.BuiltinLogger) {
	log.Printf(r.Body)
	
	if r.Err != nil {
		log.Println(r.Err)
	}
}
