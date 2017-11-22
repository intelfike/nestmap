// Gopherのためのjson操作パッケージ var1.1
package nestmap

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Nestmap struct {
	master *interface{}
	path   []interface{}
	Indent string
}

var _ fmt.Stringer = Nestmap{}

func New() *Nestmap {
	f := new(Nestmap)
	f.master = new(interface{})
	return f
}

// f location become to new json root
func (f *Nestmap) Clone() (*Nestmap, error) {
	if !f.Exists() {
		return nil, errors.New("JSON node not exists.")
	}
	newfb := new(Nestmap)
	newfb.Indent = f.Indent
	newfb.master = new(interface{})
	b, err := f.Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &newfb.master)
	return newfb, err
}
