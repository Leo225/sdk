package storage

import (
	"context"
	"io"
	"os"
	"path/filepath"

	"github.com/Leo225/sdk/mime"
)

type Uploader interface {
	Upload(ctx context.Context, reader io.Reader, filename string) (fullName string, err error)
}

type Downloader interface {
	Download(ctx context.Context, writer io.Writer, filename string) (info DownloadFileInfoer, err error)
}

type Remover interface {
	Remove(ctx context.Context, filename string) (err error)
}

type Exister interface {
	Exist(ctx context.Context, filename string) (exist bool, err error)
}

type Storager interface {
	Uploader
	Downloader
	Remover
	Exister
}

type DownloadFileInfoer interface {
	Size() int64
	Filename() string
	Metadata() Metadata
}

type downloadFileInfo struct {
	size     int64
	filename string
	metadata Metadata
}

type DefaultStorage struct {
	BasePath string
}

func NewDefaultStorage() *DefaultStorage {
	return &DefaultStorage{
		BasePath: "./",
	}
}

func (ds *DefaultStorage) Upload(ctx context.Context,
	reader io.Reader, filename string) (fullName string, err error) {
	if renameFunc, ok := FromRenameContext(ctx); ok {
		filename = renameFunc(filename)
	}

	saveFilename := filepath.Join(ds.BasePath, filename)
	dir := filepath.Dir(saveFilename)
	_, e := os.Stat(dir)
	if e != nil {
		if os.IsNotExist(e) {
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				return
			}
		} else {
			err = e
			return
		}
	}

	var dist *os.File
	dist, err = os.Create(saveFilename)
	if err != nil {
		return
	}
	defer dist.Close()

	_, err = io.Copy(dist, reader)
	if err != nil {
		return
	}
	fullName = filename
	return
}

func (ds *DefaultStorage) Download(ctx context.Context,
	writer io.Writer, filename string) (info DownloadFileInfoer, err error) {
	fullName := filepath.Join(ds.BasePath, filename)
	file, e := os.Open(fullName)
	if e != nil {
		err = e
		return
	}
	defer file.Close()

	md := Metadata{}
	if mimeType, exist := mime.Lookup(filepath.Ext(filename)); exist {
		md.Set("Content-Type", mimeType)
	}

	downloadFilename, downloadFilenameExist := FromDownloadFilenameContext(ctx)
	if !downloadFilenameExist {
		downloadFilename = filepath.Base(filename)
	}

	var fileInfo os.FileInfo
	fileInfo, err = file.Stat()
	if err != nil {
		return
	}

	info = &downloadFileInfo{
		size:     fileInfo.Size(),
		metadata: md,
		filename: downloadFilename,
	}
	if downloadBeforeFunc, downloadBeforeExist := FromDownloadBeforeContext(ctx); downloadBeforeExist {
		downloadBeforeFunc(info)
	}
	_, err = io.Copy(writer, file)
	if err != nil {
		info = nil
	}
	return
}

func (ds *DefaultStorage) Remove(ctx context.Context, filename string) (err error) {
	fullName := filepath.Join(ds.BasePath, filename)
	err = os.Remove(fullName)
	return
}

func (ds *DefaultStorage) Exist(ctx context.Context, filename string) (exist bool, err error) {
	fullName := filepath.Join(ds.BasePath, filename)
	_, err = os.Stat(fullName)
	if err != nil {
		if os.IsNotExist(err) {
			err = nil
		}
		return
	}
	exist = true
	return
}

func (dfi *downloadFileInfo) Size() int64 {
	return dfi.size
}

func (dfi *downloadFileInfo) Filename() string {
	return dfi.filename
}

func (dfi *downloadFileInfo) Metadata() Metadata {
	return dfi.metadata
}
