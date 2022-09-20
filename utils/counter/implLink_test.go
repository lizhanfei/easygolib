package counter

import (
	"testing"
	"time"
)

func TestImplLink(t *testing.T) {
	//需求：实例化一个窗口时间长度为20ms的计数器
	counter := NewCounterImplLink(20)
	//每3ms压入一个计数
	ticker := time.NewTicker(time.Millisecond * 3)

	counter.Add(1)
	<-ticker.C
	countNow, _ := counter.Count()
	if countNow != 1 {
		panic("error")
	}

	counter.Add(1)
	<-ticker.C
	countNow, _ = counter.Count()
	if countNow != 2 {
		panic("error")
	}

	counter.Add(1)
	<-ticker.C
	countNow, _ = counter.Count()
	if countNow != 3 {
		panic("error")
	}

	counter.Add(1)
	<-ticker.C
	countNow, _ = counter.Count()
	if countNow != 4 {
		panic("error")
	}

	counter.Add(1)
	<-ticker.C
	countNow, _ = counter.Count()
	if countNow != 5 {
		panic("error")
	}

	counter.Add(1)
	<-ticker.C
	countNow, _ = counter.Count()
	if countNow != 6 {
		panic("error")
	}

	counter.Add(1)
	<-ticker.C
	countNow, _ = counter.Count()
	if countNow != 6 {
		panic("error")
	}

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
