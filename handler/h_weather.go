package handler

import (
	"fmt"
	"text/template"

	"Assignment3/weather"

	"github.com/gin-gonic/gin"
)

func GetStatusHandler(c *gin.Context) {
	statusValue := weather.GetStatusValue()

	var t, err = template.ParseFiles("web/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Execute(c.Writer, statusValue)
}
