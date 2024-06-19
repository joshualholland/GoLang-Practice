package controllers

import (
	"encoding/json"
	"example/hello/models"
	"io"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func GetDriver(c *gin.Context) {
	id := c.Param("id")
	u, _ := url.Parse("https://api.openf1.org/v1/drivers")
	q := u.Query()
	q.Add("driver_number", id)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		c.AbortWithStatus(400)
	}

	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	var driver []models.Driver
	err = json.Unmarshal([]byte(responseBody), &driver)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	if len(driver) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Driver number",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Driver": driver})
}
