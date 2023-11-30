// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package note

import (
	"context"

	"NTNU_APP_2023_Final_Backend/api/note/v1"
)

type INoteV1 interface {
	CreateNote(ctx context.Context, req *v1.CreateNoteReq) (res *v1.CreateNoteRes, err error)
	DeleteNote(ctx context.Context, req *v1.DeleteNoteReq) (res *v1.DeleteNoteRes, err error)
	GetNoteContent(ctx context.Context, req *v1.GetNoteContentReq) (res *v1.GetNoteContentRes, err error)
	GetNotes(ctx context.Context, req *v1.GetNotesReq) (res *v1.GetNotesRes, err error)
	UpdateNoteContent(ctx context.Context, req *v1.UpdateNoteContentReq) (res *v1.UpdateNoteContentRes, err error)
	UpdateNoteTitle(ctx context.Context, req *v1.UpdateNoteTitleReq) (res *v1.UpdateNoteTitleRes, err error)
}
