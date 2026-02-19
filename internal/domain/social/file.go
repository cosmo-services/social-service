package social

import (
	"errors"
	"slices"
	"time"
)

var (
	ErrFileTypeNotAllowed = errors.New("FILE_TYPE_NOT_ALLOWED")
)

type FileType string

const (
	FileTypeAvatar FileType = "avatar"
)

type FileMeta struct {
	FileName  string    `json:"file_name"`
	FileType  FileType  `json:"file_type"`
	MimeType  string    `json:"mime_type"`
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
	MimeType() string
}

type FileStorage interface {
	Save(file File) (fileName string, err error)
	Delete(fileName string) (err error)
}

type FileService struct {
	metaRepo FileMetaRepository
	storage  FileStorage
}

func NewFileService(
	metaRepo FileMetaRepository,
	storage FileStorage,
) *FileService {
	return &FileService{
		metaRepo: metaRepo,
		storage:  storage,
	}
}

func (s *FileService) CreateFile(file File, userId string, fileType FileType) (*FileMeta, error) {
	mimeType := file.MimeType()
	if err := s.ValidateAllowedFileType(mimeType, fileType); err != nil {
		return nil, err
	}

	fileName, err := s.storage.Save(file)
	if err != nil {
		return nil, err
	}

	meta := &FileMeta{
		FileName: fileName,
		FileType: fileType,
		MimeType: mimeType,
		UserId:   userId,
	}

	if err := s.metaRepo.Create(meta); err != nil {
		s.storage.Delete(fileName)
		return nil, err
	}

	return meta, nil
}

func (s *FileService) ValidateAllowedFileType(mimeType string, fileType FileType) error {
	allowed := s.GetAllowedMimeType(fileType)
	if slices.Contains(allowed, mimeType) {
		return nil
	}
	return ErrFileTypeNotAllowed
}

func (s *FileService) GetAllowedMimeType(fileType FileType) []string {
	switch fileType {
	case FileTypeAvatar:
		return []string{"jpg", "jpeg", "png", "gif", "bmp"}
	}
	return nil
}
