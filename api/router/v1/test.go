package v1

import (
	"github.com/gin-gonic/gin"
	v1 "my-framework/api/controller/v1"
)

func ApiTestRouter(r *gin.RouterGroup) {
	r.GET("/ping", v1.Test)
}
