package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/app/api"
	"github.com/zngue/go_helper/pkg/app/service"
)

func Router(apiRouter *gin.RouterGroup) {
	userService := service.NewService()
	var userApi = []api.FnApi{
		api.DataApiAdd(),
		api.DataApiDelete(),
		api.DataApiList(),
		api.DataApiUpdate(),
		api.DataApiUpdateField(),
	}
	api.Router[service.ItemService]("user", userService, apiRouter, userApi...)

}
