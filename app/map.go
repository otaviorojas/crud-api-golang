package app

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Start() {
	port := os.Getenv("PORT")

	if port == "" {
		//log.Fatal("$PORT must be set")
		port = "8080"
	}

	ConfigureRouter()
	router.Run(":" + port)
}

func ConfigureRouter() {
	router = gin.Default()
	mapUrlsToControllers()
}

//This function is just to avoid making router variable globally visible
func ServeHTTP(w http.ResponseWriter, req *http.Request) {
	router.ServeHTTP(w, req)
}
