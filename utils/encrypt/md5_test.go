package encrypt

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	assert := assert2.New(t)
	originStr := "234drfgtsdedf"
	resStr := Md5(originStr)
	assert.Equal("f6a7214d5e1b7141f1f34571bdff99c8", resStr)

	originStr = "*#9中文daklasAISJDF"
	resStr = Md5(originStr)
	assert.Equal("bec6452ea7cf29a358845344f3d07e97", resStr)
}

func BenchmarkMd5(b *testing.B) {
	originStr := "123456"
	for i := 0; i < b.N; i++ {
		_ = Md5(originStr)
	}
}
