package rest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golinks/internal/db/sqlc"
	"golinks/internal/link"

	"golinks/internal/db"
)

func RegisterLinksHandlers(app *fiber.App) {
	app.Get("/links", getLinks)
	app.Get("/links/:id", getLinkByID)
	app.Put("/links", putLink)
	app.Post("/test", getLinkImages)
}

// @Summary Get all links
// @Description Long desc here
// @Tags link
// @Accept json
// @Produce json
// @Success 200 {object} sqlc.Link
// @Router /links [get]
func getLinks(c *fiber.Ctx) error {
	params := sqlc.LinkGetListParams{
		Limit:  10,
		Offset: 0,
	}
	links, err := db.Q.LinkGetList(c.Context(), params)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Internal server error")
	}

	if links == nil {
		links = []sqlc.Link{} // todo wtf?
	}

	return c.JSON(links)
}

// @Summary Get a link by its ID
// @Description
// @Tags link
// @Accept json
// @Produce json
// @Success 200 {object} ScmLink
// @Param id path int true "Link ID"
// @Router /links/{id} [get]
func getLinkByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).SendString("Bad request")
	}

	link_, err := db.Q.LinkGetByID(c.Context(), int32(id))
	if err != nil {
		fmt.Println(err)
		return c.Status(404).SendString("Link not found") // another err
	}

	return c.JSON(link_)
}

// @Summary Add new link
// @Description
// @Tags link
// @Accept json
// @Produce json
// @Param request body ScmLinkAdd true "Тело запроса"
// @Success 200 {object} ScmLink
// @Router /links [put]
func putLink(c *fiber.Ctx) error {
	var newLink LinkAddRq
	if err := c.BodyParser(&newLink); err != nil {
		return c.Status(500).SendString("Internal server error")
	}

	link_, err := db.Q.LinkAdd(c.Context(), newLink.Url)
	if err != nil {
		fmt.Println(err)
		return c.Status(500).SendString("Internal server error")
	}

	return c.JSON(IDRs{ID: int(link_.ID)})
}

// @Summary Get link images
// @Description
// @Tags test
// @Accept json
// @Produce json
// @Param request body ScmLinkAdd true "Тело запроса"
// @Router /test [post]
func getLinkImages(c *fiber.Ctx) error {
	var newLink LinkAddRq
	if err := c.BodyParser(&newLink); err != nil {
		return c.Status(500).SendString("Internal server error")
	}

	page, _ := link.LoadHtml(newLink.Url)
	images, _ := link.GetImagesFromHtml(page)

	return c.JSON(images)
}
