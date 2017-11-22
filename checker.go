package nestmap

func (f *Nestmap) Exists() bool {
	return f.ReferError() == nil
}
func (f *Nestmap) ReferError() error {
	_, err := f.GetInterfacePt()
	return err
}

func (f *Nestmap) IsString() bool {
	_, ok := (f.Interface()).(string)
	return ok
}
func (f *Nestmap) IsBool() bool {
	_, ok := (f.Interface()).(bool)
	return ok
}
func (f *Nestmap) IsInt() bool {
	_, ok := (f.Interface()).(int)
	return ok
}
func (f *Nestmap) IsUint() bool {
	_, ok := (f.Interface()).(uint)
	return ok
}
func (f *Nestmap) IsFloat() bool {
	_, ok := (f.Interface()).(float64)
	return ok
}
func (f *Nestmap) IsNull() bool {
	if !f.Exists() {
		return false
	}
	return f.Interface() == nil
}
func (f *Nestmap) IsArray() bool {
	_, ok := (f.Interface()).([]interface{})
	return ok
}
func (f *Nestmap) IsMap() bool {
	_, ok := (f.Interface()).(map[string]interface{})
	return ok
}

func (f *Nestmap) HasChild(a interface{}) bool {
	return f.Child(a).Exists()
}
