package bindata

import "github.com/qor/admin/bindata"

type Bindata struct {
	Path string
	*bindata.Bindata
}

var AssetFS *Bindata

func init() {
	AssetFS = &Bindata{Bindata: bindata.New(), Path: "config/admin/bindata"}
}

func (bindata *Bindata) Asset(name string) ([]byte, error) {
	return bindata.Bindata.AssetFileSystem.Asset(name)
}

func (bindata *Bindata) Glob(pattern string) (matches []string, err error) {
	return bindata.Bindata.AssetFileSystem.Glob(pattern)
}

func (bindata *Bindata) Compile() error {
	bindata.Bindata.CopyFiles("")
	return bindata.Bindata.AssetFileSystem.Compile()
}
