package dto

import (
	"time"

	"github.com/spacenotes-id/SpaceNotes/server/database/postgres/sqlc"
)

type (
	CreateSpaceRequest struct {
		Name     string `json:"name" valid:"required~Name is required,stringlength(3|50)~Name length should be at least between 3 - 50 characters"`
		Emoji    string `json:"emoji" valid:"stringlength(5|14)~Emoji unified code should be at least between 5 - 14 characters"`
		IsLocked bool   `json:"is_locked"`
	}
	CreateSpaceResponse struct {
		Message string              `json:"message"`
		Data    sqlc.CreateSpaceRow `json:"data"`
	}

	FindAllSpacesByUserIDResponse struct {
		Data []*sqlc.FindAllSpacesByUserIDRow `json:"data"`
	}

	FindSpaceByIDResponse struct {
		Data sqlc.Space `json:"data"`
	}

	UpdateSpaceRequest struct {
		Name      string    `json:"name" valid:"stringlength(3|50)~Name length should be at least between 3 - 50 characters"`
		Emoji     string    `json:"emoji" valid:"stringlength(5|14)~Emoji unified code should be at least between 5 - 14 characters"`
		IsLocked  bool      `json:"is_locked"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	UpdateSpaceResponse struct {
		Message string              `json:"message"`
		Data    sqlc.UpdateSpaceRow `json:"data"`
	}

	DeleteSpaceResponse struct {
		Message string `json:"message"`
	}
)
