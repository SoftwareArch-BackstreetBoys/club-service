// Package api_gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api_gen

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/oapi-codegen/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create a new club
	// (POST /club)
	CreateClub(c *fiber.Ctx) error
	// Get club info
	// (GET /club/{clubId})
	GetClubInfo(c *fiber.Ctx, clubId string) error
	// Get all clubs
	// (GET /clubs)
	GetAllClubs(c *fiber.Ctx) error
	// Search clubs by keyword
	// (GET /clubs/search)
	SearchClubs(c *fiber.Ctx, params SearchClubsParams) error
	// Get clubs joined by the user
	// (GET /clubs/user/{userId})
	GetJoinedClub(c *fiber.Ctx, userId string) error
	// Check if a user belongs to the club
	// (GET /clubs/{clubId}/is-belong)
	IsBelongToClub(c *fiber.Ctx, clubId string, params IsBelongToClubParams) error
	// Join a club
	// (POST /clubs/{clubId}/join)
	JoinClub(c *fiber.Ctx, clubId string) error
	// Leave a club
	// (POST /clubs/{clubId}/leave)
	LeaveClub(c *fiber.Ctx, clubId string) error
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

type MiddlewareFunc fiber.Handler

// CreateClub operation middleware
func (siw *ServerInterfaceWrapper) CreateClub(c *fiber.Ctx) error {

	return siw.Handler.CreateClub(c)
}

// GetClubInfo operation middleware
func (siw *ServerInterfaceWrapper) GetClubInfo(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "clubId" -------------
	var clubId string

	err = runtime.BindStyledParameterWithOptions("simple", "clubId", c.Params("clubId"), &clubId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter clubId: %w", err).Error())
	}

	return siw.Handler.GetClubInfo(c, clubId)
}

// GetAllClubs operation middleware
func (siw *ServerInterfaceWrapper) GetAllClubs(c *fiber.Ctx) error {

	return siw.Handler.GetAllClubs(c)
}

// SearchClubs operation middleware
func (siw *ServerInterfaceWrapper) SearchClubs(c *fiber.Ctx) error {

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params SearchClubsParams

	var query url.Values
	query, err = url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
	}

	// ------------- Required query parameter "keyword" -------------

	if paramValue := c.Query("keyword"); paramValue != "" {

	} else {
		err = fmt.Errorf("Query argument keyword is required, but not found")
		c.Status(fiber.StatusBadRequest).JSON(err)
		return err
	}

	err = runtime.BindQueryParameter("form", true, true, "keyword", query, &params.Keyword)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter keyword: %w", err).Error())
	}

	return siw.Handler.SearchClubs(c, params)
}

// GetJoinedClub operation middleware
func (siw *ServerInterfaceWrapper) GetJoinedClub(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "userId" -------------
	var userId string

	err = runtime.BindStyledParameterWithOptions("simple", "userId", c.Params("userId"), &userId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter userId: %w", err).Error())
	}

	return siw.Handler.GetJoinedClub(c, userId)
}

// IsBelongToClub operation middleware
func (siw *ServerInterfaceWrapper) IsBelongToClub(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "clubId" -------------
	var clubId string

	err = runtime.BindStyledParameterWithOptions("simple", "clubId", c.Params("clubId"), &clubId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter clubId: %w", err).Error())
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params IsBelongToClubParams

	var query url.Values
	query, err = url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for query string: %w", err).Error())
	}

	// ------------- Required query parameter "userId" -------------

	if paramValue := c.Query("userId"); paramValue != "" {

	} else {
		err = fmt.Errorf("Query argument userId is required, but not found")
		c.Status(fiber.StatusBadRequest).JSON(err)
		return err
	}

	err = runtime.BindQueryParameter("form", true, true, "userId", query, &params.UserId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter userId: %w", err).Error())
	}

	return siw.Handler.IsBelongToClub(c, clubId, params)
}

// JoinClub operation middleware
func (siw *ServerInterfaceWrapper) JoinClub(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "clubId" -------------
	var clubId string

	err = runtime.BindStyledParameterWithOptions("simple", "clubId", c.Params("clubId"), &clubId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter clubId: %w", err).Error())
	}

	return siw.Handler.JoinClub(c, clubId)
}

// LeaveClub operation middleware
func (siw *ServerInterfaceWrapper) LeaveClub(c *fiber.Ctx) error {

	var err error

	// ------------- Path parameter "clubId" -------------
	var clubId string

	err = runtime.BindStyledParameterWithOptions("simple", "clubId", c.Params("clubId"), &clubId, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Errorf("Invalid format for parameter clubId: %w", err).Error())
	}

	return siw.Handler.LeaveClub(c, clubId)
}

// FiberServerOptions provides options for the Fiber server.
type FiberServerOptions struct {
	BaseURL     string
	Middlewares []MiddlewareFunc
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router fiber.Router, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, FiberServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router fiber.Router, si ServerInterface, options FiberServerOptions) {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	for _, m := range options.Middlewares {
		router.Use(fiber.Handler(m))
	}

	router.Post(options.BaseURL+"/club", wrapper.CreateClub)

	router.Get(options.BaseURL+"/club/:clubId", wrapper.GetClubInfo)

	router.Get(options.BaseURL+"/clubs", wrapper.GetAllClubs)

	router.Get(options.BaseURL+"/clubs/search", wrapper.SearchClubs)

	router.Get(options.BaseURL+"/clubs/user/:userId", wrapper.GetJoinedClub)

	router.Get(options.BaseURL+"/clubs/:clubId/is-belong", wrapper.IsBelongToClub)

	router.Post(options.BaseURL+"/clubs/:clubId/join", wrapper.JoinClub)

	router.Post(options.BaseURL+"/clubs/:clubId/leave", wrapper.LeaveClub)

}
