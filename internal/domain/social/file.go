package social

import "time"

type FileType string

const (
	FileTypeAvatar FileType = "avatar"
)

type FileMeta struct {
	FileName  string    `json:"file_name"`
	FileType  FileType  `json:"file_type"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type FileMetaRepository interface {
	Create(fileMeta *FileMeta) error
	Delete(fileName string) error
	GetByName(fileName string) (*FileMeta, error)
	GetByUserId(userId string) ([]*FileMeta, error)
	Exists(fileName string) (bool, error)
}

type File interface {
	Read(p []byte) (n int, err error)
}

type FileStorage interface {
	SaveAs(fileName string, file File) error
	Delete(fileName string) error
}
