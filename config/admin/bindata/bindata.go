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
			if filepath.Match(pattern, key) {
				matches = append(matches, key)
			}
		}
		return matches
	}

	return assetFS.Bindata.AssetFileSystem.Glob(pattern)
}

func (assetFS *Bindata) Compile() error {
	assetFS.Bindata.CopyFiles(assetFS.Path)
	return assetFS.Bindata.AssetFileSystem.Compile()
}
