package server

import (
	res "xtuples/internal/api/response"

	"github.com/gin-gonic/gin"

	"xtuples/internal/api/v1/modules/cron_module"
	"xtuples/internal/api/v1/modules/exec_module"
)

func (s *server) router(app *gin.Engine) {

	app.GET("/", func(c *gin.Context) {
		res.Response(c, 200, nil, "Hello There!")
	})

	apiPath := s.app.Conf().Backend.APIPath
	api := app.Group(apiPath) // prefix for routes
	_ = api
	{

		cron_module.Router__(api)
		exec_module.Router__(api)
	}

}
