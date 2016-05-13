// -build bindata
package bindata

import "fmt"

var _bindata []string

func Asset(name string) ([]byte, error) {
	return nil, fmt.Errorf("Asset %s not found", name)
}
