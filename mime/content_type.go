package mime

import (
	"io"
	"net/http"
)

func GetFileContentType(r io.Reader) (contentType string, err error) {
	buffer := make([]byte, 512)
	_, err = r.Read(buffer)
	if err != nil {
		return
	}
	contentType = http.DetectContentType(buffer)
	return
}
