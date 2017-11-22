package nestmap

// Get json root.
func (f Nestmap) Root() *Nestmap {
	f.path = []interface{}{}
	return &f
}

// Child(...interface{}.(type) == string or int)
func (f Nestmap) Child(path ...interface{}) *Nestmap {
	for _, v := range path {
		switch v.(type) {
		case string, int:
			f.path = append(f.path, v)
		default:
			panic("Child(...interface{}.(type) == string or int)")
		}
	}
	return &f
}

// Get json parent.
func (f Nestmap) Parent() *Nestmap {
	return f.Ancestor(1)
}

// Ancestor(1) == Parent()
func (f Nestmap) Ancestor(i int) *Nestmap {
	if i < 0 {
		panic("Ancestor() argument can't set under 0")
	}
	anc := len(f.path) - i
	if 0 > anc {
		panic("JSON root has not parent.")
	}
	f.path = f.path[:anc]
	return &f
}

func (f *Nestmap) Path() []interface{} {
	return f.path
}
func (f *Nestmap) BottomPath() interface{} {
	if len(f.path) == 0 {
		return nil
	}
	return f.path[len(f.path)-1]
}

func (f *Nestmap) Depth() int {
	return len(f.path)
}
