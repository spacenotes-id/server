package dto

import (
	"time"

	"github.com/spacenotes-id/SpaceNotes/server/database/postgres/sqlc"
)

type (
	CreateNoteRequest struct {
		SpaceID int    `json:"space_id" valid:"required~Space id is required"`
		Title   string `json:"title" valid:"required~Title is required,maxstringlength(50)~Note title length should not more than 50 characters"`
		Body    string `json:"body"`
	}
	CreateNoteResponse struct {
		Message string             `json:"message"`
		Data    sqlc.CreateNoteRow `json:"data"`
	}

	FindAllNotesResponse struct {
		Data []*sqlc.FindAllNotesRow `json:"data"`
	}

	FindAllNotesBySpaceIDResponse struct {
		Data []*sqlc.FindAllNotesBySpaceIDRow `json:"data"`
	}

	FindAllTrashedNotesResponse struct {
		Data []*sqlc.FindAllTrashedNotesRow `json:"data"`
	}

	FindAllFavoriteNotesResponse struct {
		Data []*sqlc.FindAllFavoriteNotesRow `json:"data"`
	}

	FindAllArchivedNotesResponse struct {
		Data []*sqlc.FindAllArchivedNotesRow `json:"data"`
	}

	FindNoteByIDResponse struct {
		Data sqlc.Note `json:"data"`
	}

	UpdateNoteRequest struct {
		Title   string `json:"title" valid:"maxstringlength(50)~Note title length should not more than 50 characters"`
		Body    string `json:"body"`
		SpaceID int    `json:"space_id"`
		// Status    string    `json:"status" valid:"in(normal|favorite|archived|trashed)~Invalid status"`
		UpdatedAt time.Time `json:"updated_at"`
	}
	UpdateNoteResponse struct {
		Message string             `json:"message"`
		Data    sqlc.UpdateNoteRow `json:"data"`
	}

	DeleteNoteResponse struct {
		Message string `json:"message"`
	}
)
