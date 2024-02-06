package main

import (
	"github.com/gin-gonic/gin"
)

func testPage(c *gin.Context) {
	c.String(200, "Hell the World!!")
}
