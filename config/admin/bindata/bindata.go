package bindata

import (
	"path/filepath"

	"github.com/qor/admin/bindata"
)

type Bindata struct {
	Path string
	*bindata.Bindata
}

var AssetFS *Bindata

func init() {
	AssetFS = &Bindata{Bindata: bindata.New(), Path: "config/admin/bindata"}
}

func (assetFS *Bindata) Asset(name string) ([]byte, error) {
	if len(_bindata) > 0 {
		return Asset(name)
	}
	return assetFS.Bindata.AssetFileSystem.Asset(name)
}

func (assetFS *Bindata) Glob(pattern string) (matches []string, err error) {
	if len(_bindata) > 0 {
		for key, _ := range _bindata {
			if ok, err := filepath.Match(pattern, key); ok && err == nil {
				matches = append(matches, key)
			}
		}
		return matches, nil
	}

	return assetFS.Bindata.AssetFileSystem.Glob(pattern)
}

func (assetFS *Bindata) Compile() error {
	assetFS.Bindata.CopyFiles(filepath.Join(assetFS.Path, "templates"))
	return nil
}
