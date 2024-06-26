package web

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RaghavSood/isbtc1m/prices"
	"github.com/RaghavSood/isbtc1m/static"
	"github.com/RaghavSood/isbtc1m/templates"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Serve() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", s.index)
	router.GET("/why", s.why)

	router.GET("/ogimage/:slug", s.ogimage)

	router.StaticFS("/static", http.FS(static.Static))
	// Serve /favicon.ico and /robots.txt from the root
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.FileFromFS("favicon.ico", http.FS(static.Static))
	})

	router.GET("/robots.txt", func(c *gin.Context) {
		c.FileFromFS("robots.txt", http.FS(static.Static))
	})

	port := os.Getenv("ISBTC_PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}

func (s *Server) index(c *gin.Context) {
	price, err := prices.GetBTCUSDPrice()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	millionPrice := price / 1e6
	desc := fmt.Sprintf("1 BTC = $%.2fM", millionPrice)

	decision := "no"
	if price >= 1e6 {
		decision = "yes"
	}

	s.renderTemplate(c, "index.tmpl", map[string]interface{}{
		"Title":        "Is BTC 1M?",
		"Desc":         desc,
		"Price":        price,
		"MillionPrice": millionPrice,
		"OGImage":      fmt.Sprintf("https://isbtc1m.com/ogimage/%s-%.2f.png", decision, millionPrice),
	})
}

func (s *Server) why(c *gin.Context) {
	s.renderTemplate(c, "why.tmpl", map[string]interface{}{
		"Title": "Why?",
		"Desc":  "Unit bias is coming for Bitcoin",
	})
}

func (s *Server) renderTemplate(c *gin.Context, template string, params map[string]interface{}) {
	tmpl := templates.New()
	err := tmpl.Render(c.Writer, template, params)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
