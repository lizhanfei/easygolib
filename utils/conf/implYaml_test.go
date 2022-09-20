package conf

import "testing"

type confTest struct {
	Pool         pool   `yaml:"pool"`
	SwitchStatus string `yaml:"switchStatus"`
	AllowUid     []int  `yaml:"allowUid"`
}

type pool struct {
	Size int `yaml:"size"`
}

func TestImplYaml_Load(t *testing.T) {
	target := confTest{}
	confYaml := NewImplYaml()

	_ = confYaml.Load("implYamlTestConf.yaml", &target)
	if 3 != target.Pool.Size {
		panic("pool.size error")
	}
	if "open" != target.SwitchStatus {
		panic("switchStatus error")
	}
	if 4 != target.AllowUid[3] {
		panic("allowUid[3] error")
	}

	err := confYaml.Load("imlYamlTestNo.yaml", &target)
	if err == nil {
		panic("load NO error")
	}

	err = confYaml.Load("implYamlTestConfErr.yaml", &target)
	if err == nil {
		panic("load Err error")
	}
}
