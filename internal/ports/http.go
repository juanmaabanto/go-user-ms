package ports

import "github.com/juanmaabanto/go-user-ms/internal/app"

type HttpServer struct {
	app app.Application
}

func NewHttpServer(application app.Application) HttpServer {
	return HttpServer{
		app: application,
	}
}
