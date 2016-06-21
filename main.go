package main

import (
	"github.com/Gacnt/pugit-csgo/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetHTMLTemplate(hlp.ParseTemplates())
	r.Static("/static", "./static")

	r.Run(":3000")
}
