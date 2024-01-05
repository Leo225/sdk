package path

import (
	"os"
	"path/filepath"
)

func Dirs(dirname string, dirs *[]string) error {
	files, err := os.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, f := range files {
		if f.IsDir() {
			d := filepath.Join(dirname, f.Name())
			*dirs = append(*dirs, d)
			err = Dirs(d, dirs)
			if err != nil {
				return err
			}
		}
	}
	return err
}
