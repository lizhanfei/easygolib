package encrypt

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestSha1(t *testing.T) {
	assert := assert2.New(t)
	originStr := "234drfgtsdedf"
	resStr := Sha1(originStr)
	assert.Equal("450e8520c10ae2d38f42eeaf0c83d2966e6f15a3", resStr)

	originStr = "*#9中文daklasAISJDF"
	resStr = Sha1(originStr)
	assert.Equal("ef22d61641285ee41163c57b66887394ef796fa8", resStr)
}

func BenchmarkSha1(b *testing.B) {
	originStr := "234drfgtsdedf"
	for i := 0; i < b.N; i++ {
		_ = Sha1(originStr)
	}
}



func TestSha256(t *testing.T) {
	assert := assert2.New(t)
	originStr := "234drfgtsdedf"
	resStr := Sha256(originStr)
	assert.Equal("9cd4b5b392d062e75bf7d6ad6c53dd16c8fff6e5ba73171fbfb01984a2770add", resStr)

	originStr = "*#9中文daklasAISJDF"
	resStr = Sha256(originStr)
	assert.Equal("1de19f1ec69cd94aeb98f5951598689da9f0ba8cdf76b48391f539ff5e1f00df", resStr)
}

func BenchmarkSha256(b *testing.B) {
	originStr := "234drfgtsdedf"
	for i := 0; i < b.N; i++ {
		_ = Sha256(originStr)
	}
}