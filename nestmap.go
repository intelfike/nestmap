// json管理用のgithub.com/intelfike/jsonbaseをシンプルにしたもの。
// 一部関数を破棄して、また目的をjsonの管理からmap/arrayの入れ子状態の管理に変更した。
// golang標準のjsonパッケージと組み合わせて利用してもらうことになる。
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
