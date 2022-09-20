package conf

import "testing"

func TestImplJson_Load(t *testing.T) {
	target := confTest{}
	confJson := NewImplJson()

	_ = confJson.Load("implJsonTestConf.json", &target)
	if 3 != target.Pool.Size {
		panic("pool.size error")
	}
	if "open" != target.SwitchStatus {
		panic("switchStatus error")
	}
	if 4 != target.AllowUid[3] {
		panic("allowUid[3] error")
	}

	err := confJson.Load("implJsonTestNo.json", &target)
	if err == nil {
		panic("load NO error")
	}

	err = confJson.Load("implJsonTestConfErr.json", &target)
	if err == nil {
		panic("load Err error")
	}
}
