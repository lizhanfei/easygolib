package conf

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

type confTest struct {
	Pool         pool   `yaml:"pool"`
	SwitchStatus string `yaml:"switchStatus"`
	AllowUid     []int  `yaml:"allowUid"`
}

type pool struct {
	Size int `yaml:"size"`
}

func TestImplYaml_Load(t *testing.T) {
	assert := assert2.New(t)
	target := confTest{}
	confYaml := NewImplYaml()

	_ = confYaml.Load("implYamlTestConf.yaml", &target)
	assert.Equal(3, target.Pool.Size)
	assert.Equal("open", target.SwitchStatus)
	assert.Equal(4, target.AllowUid[3])

	err := confYaml.Load("imlYamlTestNo.yaml", &target)
	assert.NotNil(err)

	err = confYaml.Load("implYamlTestConfErr.yaml", &target)
	assert.NotNil(err)
}
