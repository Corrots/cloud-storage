package model

type (
	RequestLogin struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
)

type File struct {
	Size       int
	Hash       string
	Name       string
	Location   string
	UploadedAt string
}
