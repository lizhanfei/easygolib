package conf

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestImplJson_Load(t *testing.T) {
	assert := assert2.New(t)
	target := confTest{}
	confJson := NewImplJson()

	_ = confJson.Load("implJsonTestConf.json", &target)
	assert.Equal(target.Pool.Size, 3)
	assert.Equal(target.SwitchStatus, "open")
	assert.Equal(target.AllowUid[3], 4)

	err := confJson.Load("implJsonTestNo.json", &target)
	assert.NotNil(err)

	err = confJson.Load("implJsonTestConfErr.json", &target)
	assert.NotNil(err)
}
