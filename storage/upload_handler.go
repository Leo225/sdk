package storage

import (
	"context"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

const (
	defaultMaxMemory = 32 << 20 // 32MB
)

type UploadStorager interface {
	Uploader
	MaxMemory() int64
}

type DefaultUploadStorage struct {
	*DefaultStorage
}

func (ds *DefaultUploadStorage) MaxMemory() int64 {
	return defaultMaxMemory
}

func NewDefaultUploadStorage() *DefaultUploadStorage {
	return &DefaultUploadStorage{
		DefaultStorage: NewDefaultStorage(),
	}
}

type UploadFileInfoer interface {
	FullName() string
	Filename() string
	Size() int64
	Header() textproto.MIMEHeader
}

type uploadFileInfo struct {
	fullName string
	filename string
	size     int64
	header   textproto.MIMEHeader
}

func (uf *uploadFileInfo) FullName() string {
	return uf.fullName
}

func (uf *uploadFileInfo) Filename() string {
	return uf.filename
}

func (uf *uploadFileInfo) Size() int64 {
	return uf.size
}

func (uf *uploadFileInfo) Header() textproto.MIMEHeader {
	return uf.header
}

func UploadHandler(ctx context.Context,
	r *http.Request, storage UploadStorager, name string) (infos []UploadFileInfoer, err error) {
	if storage == nil {
		storage = NewDefaultUploadStorage()
	}

	if r.MultipartForm == nil {
		err = r.ParseMultipartForm(storage.MaxMemory())
		if err != nil {
			return
		}
	}

	for _, f := range r.MultipartForm.File[name] {
		var file multipart.File
		file, err = f.Open()
		if err != nil {
			return
		}

		var fullName string
		fullName, err = storage.Upload(ctx, file, f.Filename)
		file.Close()
		if err != nil {
			return
		}

		infos = append(infos, &uploadFileInfo{
			fullName: fullName,
			filename: f.Filename,
			size:     f.Size,
			header:   f.Header,
		})
	}
	return
}

type sizer interface {
	Size() int64
}

type stater interface {
	Stat() (os.FileInfo, error)
}

// GetMultipartFileSize 获取上传文件大小
// 相关问题：https://github.com/golang/go/issues/19501
func GetMultipartFileSize(file multipart.File) (size int64, err error) {
	if s, ok := file.(sizer); ok {
		size = s.Size()
	} else if st, ok := file.(stater); ok {
		var fileInfo os.FileInfo
		fileInfo, err = st.Stat()
		if err != nil {
			return
		}
		size = fileInfo.Size()
	}
	return
}
