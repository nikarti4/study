package http_handlers

import (
	"net/http"

	"L0/cache"
	"github.com/gin-gonic/gin"
)

func GetServerOrderByID(cache *cache.My_cache) func(c *gin.Context) {

	f := func(c *gin.Context) {
		id := c.Param("id")
		orderFromCache, is_ok := cache.Read(id)

		if is_ok {
			c.IndentedJSON(http.StatusOK, orderFromCache)
			return
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "order not found"})
	}

	return gin.HandlerFunc(f)
}
