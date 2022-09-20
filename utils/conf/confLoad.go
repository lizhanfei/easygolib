package conf

type Conf interface {
	Load(filePath string, t interface{}) error
}

func NewImplYaml() Conf {
	return &ImplYaml{}
}

func NewImplJson() Conf {
	return &ImplJson{}
}
