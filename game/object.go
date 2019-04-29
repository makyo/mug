package game

type Object struct {
	Type     Flag
	Flags    []Flag
	Name     string
	Owner    DBRef
	Settings map[string]string
}

func (o *Object) GetSetting(key string) (string, error) {
}

func (o *Object) SetSetting(key string) (string, error) {
}
