package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/airbornharsh/github-codespace/service/pkg/helpers"
	"github.com/airbornharsh/github-codespace/service/pkg/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	var c time.Timer
	c.C = make(chan time.Time)
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// r.GET("/ws", func(c *gin.Context) {
	// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	// 	if err != nil {
	// 		return
	// 	}
	// 	defer conn.Close()
	// 	i := 0
	// 	for {
	// 		i++
	// 		conn.WriteMessage(websocket.TextMessage, []byte("Hello, WebSocket!"))
	// 		time.Sleep(time.Second)
	// 		fmt.Println("Sent message", i)
	// 	}
	// })

	r.GET("/*any", func(c *gin.Context) {
		if c.Request.URL.Path == "/ws" {
			done := helpers.StartWebSocket(c, &upgrader)
			if !done {
				c.JSON(500, gin.H{
					"message": "Error starting WebSocket",
				})
			}
		}

		imageId := strings.Split(c.Request.Host, ".")[0]

		containersData, err := helpers.ReadContainersData()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Error reading containers data",
				"error":   err.Error(),
			})
			return
		}

		containerInfo, ok := containersData[imageId]
		if !ok {
			c.JSON(404, gin.H{
				"message": "Container not found",
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
