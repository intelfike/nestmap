package nestmap

// if has child then
import (
	"errors"
	"strconv"
)

// update child.
//
// else
// make child.
//
// Can't make array.
func (f *Nestmap) Set(i interface{}) error {
	if len(f.path) == 0 {
		*f.master = i
		return nil
	} else {
		// topがマップや配列じゃなかったら作る
		switch (*f.master).(type) {
		case []interface{}, map[string]interface{}:
		default:
			*f.master = map[string]interface{}{}
		}
	}
	cur := new(interface{})
	*cur = *f.master
	for n, pathv := range f.path {
		switch mas := (*cur).(type) {
		case map[string]interface{}:
			pt, ok := pathv.(string)
			if !ok {
				pt = strconv.Itoa(pathv.(int))
			}
			*cur, ok = mas[pt]
			if !ok {
				paths := []string{}
				for _, v := range f.path[n:] {
					s, ok := v.(string)
					if !ok {
						s = strconv.Itoa(v.(int))
					}
					paths = append(paths, s)
				}
				mapNest(mas, i, 0, paths...)
				return nil
			}
			if n == len(f.path)-1 {
				mas[pt] = i
			}
		case []interface{}:
			switch pt := pathv.(type) {
			case string:
				f.Parent().Set(map[string]interface{}{})
				f.Set(i)
			case int:
				if 0 > pt || pt >= len(mas) {
					return errors.New("Array index out of range.")
				}
				if n == len(f.path)-1 { // 最後の要素なら
					mas[pt] = i
					return nil
				}
				*cur = mas[pt]
			}
		default:
		}
	}
	return nil
}
func mapNest(m map[string]interface{}, val interface{}, depth int, s ...string) {
	if depth == len(s)-1 {
		m[s[depth]] = val
		return
	}
	mm := map[string]interface{}{s[depth+1]: nil}
	m[s[depth]] = mm
	mapNest(mm, val, depth+1, s...)
}

// It like append().
//
// If json node is array then add i.
//
// else then set []interface{i} (initialize for array).
func (f *Nestmap) Push(a interface{}) error {
	if !f.IsArray() {
		f.Set([]interface{}{a})
		return nil
	}
	pt, err := f.GetInterfacePt()
	if err != nil {
		return err
	}
	ar := (*pt).([]interface{})
	return f.Set(append(ar, a))
}

// Remove() remove map or array element
func (f *Nestmap) Remove() error {
	if len(f.path) == 0 {
		return errors.New("Root can't remove. Use Empty().")
	}
	path := f.path[len(f.path)-1]
	pt, err := f.Parent().GetInterfacePt()
	if err != nil {
		return err
	}
	switch t := (*pt).(type) {
	case map[string]interface{}:
		st, ok := path.(string)
		if !ok {
			st = strconv.Itoa(path.(int))
		}
		_, ok = t[st]
		if !ok {
			return errors.New("JSON node not exists.")
		}
		delete(t, st)
	case []interface{}:
		it, ok := path.(int)
		if !ok {
			return errors.New("JSON node is not Map().")
		}
		if 0 > it || it >= len(t) {
			return errors.New("Array index out of range.")
		}
		return f.Parent().Set(append(t[:it], t[it+1:]...))
	default:
		return errors.New("JSON node not exists. ")
	}
	return nil
}
func (f *Nestmap) Empty() error {
	if f.IsArray() {
		f.Set([]interface{}{})
	} else if f.IsMap() {
		f.Set(map[string]interface{}{})
	} else {
		return errors.New("JSON node is not Array or Map. or node not exists.")
	}
	return nil
}
