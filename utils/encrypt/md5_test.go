package encrypt

import (
	"testing"
)

func TestMd5(t *testing.T) {
	originStr := "234drfgtsdedf"
	resStr := Md5(originStr)

	if "f6a7214d5e1b7141f1f34571bdff99c8" != resStr {
		panic("md5 error")
	}

	originStr = "*#9中文daklasAISJDF"
	resStr = Md5(originStr)

	if "bec6452ea7cf29a358845344f3d07e97" != resStr {
		panic("md5 error")
	}
}

func BenchmarkMd5(b *testing.B) {
	originStr := "123456"
	for i := 0; i < b.N; i++ {
		_ = Md5(originStr)
	}
}
