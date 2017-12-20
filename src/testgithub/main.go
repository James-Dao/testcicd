package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "servicebroker"
	app.Usage = "Start components hi10"
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) {
		router := gin.New()
		router.Use(gin.Recovery())

		events := router.Group("/events")
		{
			events.POST("/github", Handle)
		}
		router.GET("/healthz", healthz)
		router.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong")
		})
		server := &http.Server{
			Addr:           ":8080",
			Handler:        router,
			ReadTimeout:    300 * time.Second,
			WriteTimeout:   300 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		server.ListenAndServe()
	}
	app.Run(os.Args)
}

func healthz(c *gin.Context) {
	c.String(http.StatusOK, http.StatusText(http.StatusOK))
}

func Handle(c *gin.Context) {
	event := c.Request.Header.Get("X-GitHub-Event")
	switch event {
	case "ping":
		log.Print("Received ping from GitHub")
		c.JSON(200, gin.H{"message": "OK"})
		return
	case "push", "pull_request", "create", "release", "status", "commit_comment", "pull_request_review":
		handleEvent(c, event)
		return
	default:
		// Issue #127: Don't return an error for unimplemented events.
		log.Printf("Unsupported event %q", event)
		c.JSON(200, gin.H{"message": "Ignored"})
		return
	}
}

func handleEvent(c *gin.Context, eventType string) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Failed to read body: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Malformed body"})
		return
	}
	defer c.Request.Body.Close()

	e, err := github.ParseWebHook(eventType, body)
	if err != nil {
		log.Printf("Failed to parse body: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Malformed body"})
		return
	}
	log.Panicf("github body : %+v", body)
	c.JSON(http.StatusOK, gin.H{"status": "Complete"})
}
