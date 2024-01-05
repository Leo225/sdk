package storage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Leo225/sdk/mime"
)

type DownloadStorager interface {
	Downloader
	DispositionType() string
}

type DefaultDownloadStorage struct {
	*DefaultStorage
	dispositionType string
}

func (ds *DefaultDownloadStorage) DispositionType() string {
	return ds.dispositionType
}

func NewDefaultDownloadStorage(dispositionType string) *DefaultDownloadStorage {
	return &DefaultDownloadStorage{
		DefaultStorage:  NewDefaultStorage(),
		dispositionType: dispositionType,
	}
}

func DownloadHandler(ctx context.Context,
	w http.ResponseWriter, storage DownloadStorager, filename string) (err error) {
	if storage == nil {
		storage = NewDefaultDownloadStorage("attachment")
	}

	if _, ok := FromDownloadBeforeContext(ctx); !ok {
		ctx = NewDownloadBeforeContext(ctx, func(info DownloadFileInfoer) {
			md := info.Metadata()
			wr := false
			if md != nil {
				contentType := md.Get("Content-Type")
				if contentType != "" {
					w.Header().Set("Content-Type", contentType)
					wr = true
				}
			}
			if !wr {
				var contentType string
				contentType, err = mime.DeleteContentType(info.Filename())
				if err != nil {
					return
				}
				w.Header().Set("Content-Type", contentType)
			}

			dispositionType := storage.DispositionType()
			if dispositionType == "" || (dispositionType != "inline" && dispositionType != "attachment") {
				dispositionType = "attachment"
			}
			w.Header().Set("Content-Disposition",
				fmt.Sprintf("%s; filename=\"%s\"", dispositionType, info.Filename()))
		})
	}
	_, err = storage.Download(ctx, w, filename)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
