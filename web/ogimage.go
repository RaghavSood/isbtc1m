package web

import (
	"fmt"
	"image"
	"image/color"
	"net/http"
	"strings"

	"github.com/RaghavSood/isbtc1m/static"
	"github.com/RaghavSood/ogimage"
	"github.com/gin-gonic/gin"
)

func (s *Server) ogimage(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.String(http.StatusBadRequest, "missing slug")
		return
	}

	// Generate the image
	img, err := s.generateOGImage(slug)
	if err != nil {
		c.String(http.StatusInternalServerError, "failed to generate image")
		return
	}

	// Write the image to the response
	c.Header("Content-Type", "image/png")
	c.Header("Content-Length", fmt.Sprintf("%d", len(img)))
	c.Writer.Write(img)
}

func (s *Server) generateOGImage(slug string) ([]byte, error) {
	parts := strings.Split(slug, "-")

	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid slug: %s", slug)
	}

	if parts[0] != "no" && parts[0] != "yes" {
		return nil, fmt.Errorf("invalid slug: %s", slug)
	}

	slugParts := strings.Split(parts[1], ".")

	if len(slugParts) != 3 {
		return nil, fmt.Errorf("invalid slug: %s", slug)
	}

	if slugParts[2] != "png" {
		return nil, fmt.Errorf("invalid slug: %s", slug)
	}

	var title string
	var subtitle string
	var description string
	var err error

	title = "Is BTC 1M?"
	subtitle = strings.ToUpper(parts[0])
	description = fmt.Sprintf("1 BTC = $%s.%sM", slugParts[0], slugParts[1])

	templateBytes, err := static.Static.ReadFile(fmt.Sprintf("template-%s.png", parts[0]))
	if err != nil {
		return nil, err
	}

	logoBytes, err := static.Static.ReadFile("logo.png")
	if err != nil {
		return nil, err
	}

	fontBytes, err := static.Static.ReadFile("menlo.ttf")
	if err != nil {
		return nil, err
	}

	ogImage, err := ogimage.NewOgImage(templateBytes, logoBytes)

	titleText := ogimage.Text{
		Content:  title,
		FontData: fontBytes,
		FontSize: 32,
		Color:    color.White,
		Point:    image.Point{520, 200},
	}

	subStart := 485
	if parts[0] == "yes" {
		subStart = 380
	}
	subtitleText := ogimage.Text{
		Content:  subtitle,
		FontData: fontBytes,
		FontSize: 256,
		Color:    color.White,
		Point:    image.Point{subStart, 410},
	}

	descriptionText := ogimage.Text{
		Content:  description,
		FontData: fontBytes,
		FontSize: 48,
		Color:    color.White,
		Point:    image.Point{430, 470},
	}

	config := ogimage.Config{
		Position: ogimage.TopRight,
		Padding:  20,
		Texts:    []ogimage.Text{titleText, subtitleText, descriptionText},
	}

	imageBytes, err := ogImage.Generate(config)
	if err != nil {
		return nil, err
	}

	return imageBytes, nil
}
