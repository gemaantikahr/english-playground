package web

import (
	"english_playground/app/base"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Web struct{}

func (Web) Do(router *base.Router) {
	router.GET("/", func(c *gin.Context) {
		_, err := fmt.Fprintf(c.Writer, "Hello world!")
		if err != nil {
			return
		}
	})
}

func Init() Web {
	return Web{}
}
