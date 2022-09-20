package rateLimit

import (
	"fmt"
	"testing"
	"time"
)

func TestImplTokenBucket(t *testing.T) {
	//assert := assert2.New(t)
	//1s 允许1000个请求，最大并发100
	stopChan := false
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		<-ticker.C
		stopChan = true
	}()
	ok := 0

	//需求：每秒只允许1000个请求，最大并发300
	create := 10              //每次新增令牌数
	createFrequency := 10     //新增间隔时间；单位ms
	capacity := 300           //令牌桶容量
	var maxWaitTime int64 = 5 //最大的等待时间；单位ms
	rateLimit := NewImplTokenBucket(create, createFrequency, capacity)
	for {
		if stopChan {
			break
		}
		if rateLimit.Check(maxWaitTime) {
			ok++
		}
	}
	fmt.Println(fmt.Sprintf("v1 res %d", ok/10))
}

func TestImplTokenBucketDefault(t *testing.T) {
	stopChan := false
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		<-ticker.C
		stopChan = true
	}()
	ok := 0

	//需求：使用默认值
	rateLimit := NewImplTokenBucket(0, 0, 0)
	for {
		if stopChan {
			break
		}
		if rateLimit.Check(0) {
			ok++
		}
	}
	fmt.Println(fmt.Sprintf("default res %d", ok/10))
}

func TestImplTokenBucketV2(t *testing.T) {
	stopChan := false
	ticker := time.NewTicker(time.Second * 10)
	go func() {
		<-ticker.C
		stopChan = true
	}()
	ok := 0

	//需求：每秒只允许300个请求,并发最大不超过100
	create := 3              //每次新增令牌数
	createFrequency := 10     //新增间隔时间；单位ms
	capacity := 100           //令牌桶容量
	var maxWaitTime int64 = 5 //最大的等待时间；单位ms
	rateLimit := NewImplTokenBucket(create, createFrequency, capacity)
	for {
		if stopChan {
			break
		}
		if rateLimit.Check(maxWaitTime) {
			ok++
		}
	}
	fmt.Println(fmt.Sprintf("v2 res: %d", ok/10))
}

func BenchmarkImplTokenBucket(b *testing.B) {
	ok := 0
	rateLimit := NewImplTokenBucket(100, 10, 300)
	for i := 0; i < b.N; i++ {
		if rateLimit.Check(1) {
			ok++
		}
	}
}
