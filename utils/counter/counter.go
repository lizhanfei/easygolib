package counter

type WindowCounter interface {
	//Add 新增
	Add(num int) error
	//Count 获取当前总数
	Count() (int64, error)
}

//NewCounterImplLink 基于链表的滑动窗口计数器
//expireTime 窗口时间长度；单位ms
func NewCounterImplLink(expireTime int64) WindowCounter {
	return &ImplLink{
		expireTime: expireTime,
		count:      0,
		linkManager: &linkManager{
			header: nil,
			tail:   nil,
		},
	}
}
