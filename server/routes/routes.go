package routes

import (
	"github.com/camcelve/frigate-app/server/api/todo/route"
	
	"github.com/camcelve/frigate-app/server/common/static"
	
	"gopkg.in/gin-gonic/gin.v1"
)

func Init(r *gin.Engine) {
	
	static.Init(r)
	
	todoroute.Init(r)
}
