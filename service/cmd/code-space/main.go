package main

import (
	"fmt"
	// "net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/airbornharsh/github-codespace/service/pkg/helpers"
	"github.com/airbornharsh/github-codespace/service/pkg/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	// "github.com/gorilla/websocket"
)

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

func main() {
	var c time.Timer
	c.C = make(chan time.Time)
	gin.SetMode(gin.ReleaseMode)

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.GET("/*any", func(c *gin.Context) {
		// if c.Request.URL.Path == "/ws" {
		// 	done := helpers.StartWebSocket(c, &upgrader)
		// 	if !done {
		// 		c.JSON(500, gin.H{
		// 			"message": "Error starting WebSocket",
		// 		})
		// 	}
		// }

		id := strings.Split(c.Request.Host, ".")[0]
		if id == "www" {
			id = strings.Split(c.Request.Host, ".")[1]
		}

		containerInfo, err := helpers.ReadContainersData(id)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error reading containers data",
				"error":   err.Error(),
			})
			return
		}

		internalServerURL := "http://localhost:" + strconv.Itoa(containerInfo.Port)
		internalServer, err := url.Parse(internalServerURL)
		if err != nil {
			fmt.Println("Error parsing internal server URL:", err)
			return
		}

		httpProxy := httputil.NewSingleHostReverseProxy(internalServer)

		httpProxy.ServeHTTP(c.Writer, c.Request)
	})

	routes.Init(r)

	fmt.Println("Server Started at http://localhost:5000")
	r.Run(":5000")
}
