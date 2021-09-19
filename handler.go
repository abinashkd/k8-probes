package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type probeResponse struct {
	Pod            string `json:"pod"`
	Conatiner      string `json:"container"`
	LivenessProbe  string `json:"livenessprobe"`
	ReadinessProbe string `json:"readinessprobe"`
	StartupProbe   string `json:"startupprobe"`
}

func getProbes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, probeResponse)
}

func main() {

	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")

}
