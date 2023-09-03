package dto

import "github.com/spacenotes-id/SpaceNotes/server/database/postgres/sqlc"

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
)
