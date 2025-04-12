package rest

import (
	"github.com/gofiber/fiber/v2"
	"golinks/internal/db"
	"golinks/internal/db/sqlc"
)

func RegisterLinksHandlers(app *fiber.App) {
	app.Get("/links", getLinks)
	app.Get("/links/:id", getLinkByID)
	app.Put("/links", putLink)
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
		return errorResponse(c, 500, err)
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
		return errorResponse(c, 400, err)
	}

	link_, err := db.Q.LinkGetByID(c.Context(), int32(id))
	if err != nil {
		return errorResponse(c, 404, err)
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

	err := c.BodyParser(&newLink)
	if err != nil {
		return errorResponse(c, 500, err)
	}

	link_, err := db.Q.LinkAdd(c.Context(), newLink.Url)
	if err != nil {
		return errorResponse(c, 500, err)
	}

	return c.JSON(IDRs{ID: int(link_.ID)})
}
