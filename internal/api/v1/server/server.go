package server

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"xtuples/pkg/app"
)

type server struct {
	app *app.App
}

func NewServer(app *app.App) *server {
	return &server{app}
}

// Start the server on the specified host and port + with the defined routes
func (s *server) Start() {
	backend := gin.Default() // create the server
	s.router(backend)        // use endpoints

	conf := s.app.Conf()
	address := fmt.Sprintf("%v:%v", conf.Backend.Host, conf.Backend.Port)
	backend.Run(address)
}
