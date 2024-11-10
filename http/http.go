package http_server

import (
	"context"
	"fmt"

	"github.com/SoftwareArch-BackstreetBoys/club-service/application"
	"github.com/SoftwareArch-BackstreetBoys/club-service/http/auth_util"
	api_gen "github.com/SoftwareArch-BackstreetBoys/club-service/http/gen"
	"github.com/SoftwareArch-BackstreetBoys/club-service/model"
	"github.com/gofiber/fiber/v2"
)

type Http struct {
	app application.Application
}

func NewHttp(app application.Application) api_gen.ServerInterface {
	return &Http{app: app}
}

func (h *Http) DeleteClub(c *fiber.Ctx, clubId string) error {
	user, err := auth_util.GetUserFromFiberContext(c)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "invalid authentication"})
	}

	club, err := h.app.GetClubInfo(context.Background(), clubId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "club not found"})
	}

	if club.CreatedByID != user.Id {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "you are not the club owner"})
	}

	deletedClub, err := h.app.DeleteClub(context.Background(), clubId)

	return c.Status(fiber.StatusOK).JSON(deletedClub)
}

func (h *Http) PatchClubInfo(c *fiber.Ctx, clubId string) error {
	user, err := auth_util.GetUserFromFiberContext(c)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "invalid authentication"})
	}

	club, err := h.app.GetClubInfo(context.Background(), clubId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "club not found"})
	}

	if club.CreatedByID != user.Id {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "you are not the club owner"})
	}

	var clubPatchInfo api_gen.PatchClubInfoJSONRequestBody
	if err := c.BodyParser(&clubPatchInfo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	updatedClub, err := h.app.UpdateClub(context.Background(), clubId, model.UpdateClubInfo{
		Name:        clubPatchInfo.Name,
		Description: clubPatchInfo.Description,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update club"})
	}

	return c.Status(fiber.StatusOK).JSON(updatedClub)
}

func (h *Http) GetAllClubs(c *fiber.Ctx) error {
	clubs, err := h.app.GetAllClubs(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(clubs)
}

func (h *Http) CreateClub(c *fiber.Ctx) error {
	user, err := auth_util.GetUserFromFiberContext(c)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "invalid authentication"})
	}

	var club api_gen.CreateClubJSONRequestBody
	if err := c.BodyParser(&club); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
	}

	newClub, err := h.app.CreateClub(context.Background(), model.Club{
		Name:              club.Name,
		Description:       club.Description,
		CreatedByID:       user.Id,
		CreatedByFullName: user.FullName,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(newClub)
}

func (h *Http) SearchClubs(c *fiber.Ctx, params api_gen.SearchClubsParams) error {
	clubs, err := h.app.SearchClubs(context.Background(), params.Keyword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(clubs)
}

func (h *Http) GetJoinedClub(c *fiber.Ctx) error {
	user, err := auth_util.GetUserFromFiberContext(c)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "invalid authentication"})
	}

	userId := user.Id

	clubs, err := h.app.GetJoinedClub(context.Background(), userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(clubs)
}

func (h *Http) GetClubInfo(c *fiber.Ctx, clubId string) error {
	club, err := h.app.GetClubInfo(context.Background(), clubId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(club)
}

func (h *Http) IsBelongToClub(c *fiber.Ctx, clubId string, params api_gen.IsBelongToClubParams) error {
	isBelong, err := h.app.IsBelongToClub(context.Background(), clubId, params.UserId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"isBelong": isBelong})
}

func (h *Http) JoinClub(c *fiber.Ctx, clubId string) error {
	user, err := auth_util.GetUserFromFiberContext(c)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "invalid authentication"})
	}

	userId := user.Id

	err = h.app.JoinClub(context.Background(), clubId, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Successfully joined"})
}

func (h *Http) LeaveClub(c *fiber.Ctx, clubId string) error {
	user, err := auth_util.GetUserFromFiberContext(c)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "invalid authentication"})
	}

	userId := user.Id

	err = h.app.LeaveClub(context.Background(), clubId, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Successfully left"})
}

func (h *Http) GetHealthService(c *fiber.Ctx) error {
	serviceStatus := "Service is running"

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"service": serviceStatus,
	})
}

func (h *Http) GetHealthDatabase(c *fiber.Ctx) error {
	dbErr := h.app.CheckDatabaseConnection(context.Background())
	var dbStatus string
	if dbErr != nil {
		dbStatus = "Database connection failed: " + dbErr.Error()
	} else {
		dbStatus = "Database connection is healthy"
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"database": dbStatus,
	})
}
