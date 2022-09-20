package counter

import (
	assert2 "github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestImplLink(t *testing.T) {
	assert := assert2.New(t)
	//需求：实例化一个窗口时间长度为20ms的计数器
	counter := NewCounterImplLink(20)
	//每3ms压入一个计数
	ticker := time.NewTicker(time.Millisecond * 3)

	counter.Add(1)
	<-ticker.C
	countNow, _ := counter.Count()
	assert.Equal(1, countNow)

	counter.Add(1)
	<-ticker.C
	countNow, _ = counter.Count()
	assert.Equal(2, countNow)

	counter.Add(1)
	<-ticker.C
	countNow, _ = counter.Count()
	assert.Equal(3, countNow)

	counter.Add(1)
	<-ticker.C
	countNow, _ = counter.Count()
	assert.Equal(4, countNow)

	counter.Add(1)
	<-ticker.C
	countNow, _ = counter.Count()
	assert.Equal(5, countNow)

	counter.Add(1)
	<-ticker.C
	countNow, _ = counter.Count()
	assert.Equal(6, countNow)

	counter.Add(1)
	<-ticker.C
	countNow, _ = counter.Count()
	assert.Equal(6, countNow)
}

func BenchmarkImplLink_Add(b *testing.B) {
	counter := NewCounterImplLink(20)

	for i := 0; i < b.N; i++ {
		counter.Add(1)
	}
}

func BenchmarkImplLink_Count(b *testing.B) {
	//5ms的时间窗口期
	counter := NewCounterImplLink(5)
	//每1ms，写入1次
	go func() {
		ticker := time.NewTicker(time.Millisecond)
		for {
			<- ticker.C
			counter.Add(1)
		}
	}()
	//测试counter的性能
	for i := 0; i < b.N; i++ {
		 counter.Count()
	}
}
