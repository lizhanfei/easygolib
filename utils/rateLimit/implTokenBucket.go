package rateLimit

import "time"

//基于令牌桶实现的限流器
type ImplTokenBucket struct {
	Create          int       //每次新增令牌数；默认 10
	CreateFrequency int       //新增间隔时间；单位ms；默认10ms
	Capacity        int       //最大容量；默认 100
	tokenBucket     chan bool //令牌桶
}

//Check 检查是否允许请求；waitTimeOut 指定超时时间 ms 默认使用 this.CreateFrequency
func (this *ImplTokenBucket) Check(waitTimeOut int64) bool {
	if waitTimeOut <= 0 {
		waitTimeOut = int64(this.CreateFrequency) + 1
	}
	ticker := time.NewTicker(time.Millisecond * time.Duration(waitTimeOut))
	defer ticker.Stop()
	select {
	case <-ticker.C:
		return false
	case <-this.tokenBucket:
		return true
	}
}

//addToken 添加令牌到桶
func (this *ImplTokenBucket) addToken() {
	create := this.Create
	for {
		if create <= 0 || len(this.tokenBucket) >= this.Capacity {
			break
		}
		this.tokenBucket <- true
		create--
	}
}

func (this *ImplTokenBucket) worker() {
	go func() {
		ticker := time.NewTicker(time.Millisecond * time.Duration(this.CreateFrequency))
		for {
			<-ticker.C
			this.addToken()
		}
	}()
}
