package rateLimit

type RateLimit interface {
	//CheckWithTimeOut 指定超时时间限流检查
	Check(waitTimeOut int64) bool
}

//NewImplTokenBucket 实例化基于令牌桶的限流器；
//create	每次新增令牌数；默认 10
//createFrequency  新增间隔时间；单位ms；默认10ms
//capacity   最大容量；默认 100
func NewImplTokenBucket(create, createFrequency, capacity int) RateLimit {
	if create <= 0 {
		create = 10
	}
	if createFrequency <= 0 {
		createFrequency = 10
	}
	if capacity <= 0 {
		capacity = 100
	}
	bucket := make(chan bool, capacity)
	rateLimit := &ImplTokenBucket{
		Create:          create,
		CreateFrequency: createFrequency,
		Capacity:        capacity,
		tokenBucket:     bucket,
	}
	rateLimit.worker()
	return rateLimit
}