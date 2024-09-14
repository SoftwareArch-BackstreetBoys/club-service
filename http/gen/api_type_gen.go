// Package api_gen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package api_gen

// Club defines model for Club.
type Club struct {
	Description string  `json:"description"`
	Id          *string `json:"id,omitempty"`
	Name        string  `json:"name"`
}

// SearchClubsParams defines parameters for SearchClubs.
type SearchClubsParams struct {
	Keyword string `form:"keyword" json:"keyword"`
}

// IsBelongToClubParams defines parameters for IsBelongToClub.
type IsBelongToClubParams struct {
	UserId string `form:"userId" json:"userId"`
}

// JoinClubJSONBody defines parameters for JoinClub.
type JoinClubJSONBody struct {
	UserId string `json:"userId"`
}

// LeaveClubJSONBody defines parameters for LeaveClub.
type LeaveClubJSONBody struct {
	UserId string `json:"userId"`
}

// CreateClubJSONRequestBody defines body for CreateClub for application/json ContentType.
type CreateClubJSONRequestBody = Club

// JoinClubJSONRequestBody defines body for JoinClub for application/json ContentType.
type JoinClubJSONRequestBody JoinClubJSONBody

// LeaveClubJSONRequestBody defines body for LeaveClub for application/json ContentType.
type LeaveClubJSONRequestBody LeaveClubJSONBody
