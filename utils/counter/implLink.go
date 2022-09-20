package counter

import "time"



type ImplLink struct {
	expireTime  int64 //过期时间，以毫秒为单位
	count       int64 //有效节点数
	linkManager *linkManager
}

//Add 添加节点
func (this *ImplLink) Add(num int) error {
	_ = this.linkManager.Add(int64(num))
	return nil
}

//Count 统计当前链表的有效节点数
func (this *ImplLink) Count() (int64, error) {
	//获取最小时间
	return this.linkManager.maintain(this.getMinTime())
}

//getMinTime 获取最小时间
func (this *ImplLink) getMinTime() int64 {
	return time.Now().UnixNano()/1e6 - this.expireTime
}

//linkNode 链表节点
type linkNode struct {
	num  int64
	time int64 //压入时间，单位为毫秒
	pre  *linkNode
	next *linkNode
}

//linkManager 链表管理
type linkManager struct {
	header *linkNode
	tail   *linkNode
}

//Add 加入节点
func (this *linkManager) Add(num int64) error {
	//获取当前时间毫秒数
	timeNow := time.Now().UnixNano() / 1e6
	//生成新的节点
	newNode := linkNode{
		num:  num,
		time: timeNow,
	}
	if nil == this.header {
		this.header = &newNode
	}
	newNode.pre = this.tail
	//写入链表末尾
	if nil != this.tail {
		this.tail.next = &newNode
	}
	this.tail = &newNode
	return nil
}

//maintain 根据过期时间，维护节点，并且把当前有效的节点数返回
func (this *linkManager) maintain(minTime int64) (int64, error) {
	//获取节点的尾部
	var count int64
	tailNode := this.tail
	for {
		if tailNode == nil {
			break
		}
		//从尾部逐个向前扫描，判断时间是否过期；同时统计有效节点数
		if tailNode.time <= minTime { //如果当前节点的时间小于最小时间，那么认为已过期
			//遇到第一个过期的节点
			//把header指向第一个不过期的节点
			this.header = tailNode.next
			//抛弃第一个节点前的过期节点
			if nil != this.header {
				this.header.pre = nil
			}
			break
		} else {
			//未过期
			count += tailNode.num
			tailNode = tailNode.pre
		}
	}
	return count, nil
}
