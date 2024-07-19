package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/transform", func(c *gin.Context) {
		content, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read the body"})
			return
		}

		newNamespace := c.GetHeader("Namespace")
		newName := c.GetHeader("Name")
		if newNamespace == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Namespace header is required"})
			return
		}
		if newName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Name header is required"})
			return
		}

		transformedContent, err := parseAndTransformYAML(content, newNamespace, newName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to transform the YAML"})
			return
		}

		c.Data(http.StatusOK, "application/x-yaml", transformedContent)
	})

	router.Run(":8080")
}
