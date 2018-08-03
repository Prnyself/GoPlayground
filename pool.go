package main

const chanSize = 10

func main() {
	syncChan := make(chan int, chanSize)

	// 生产者
	// 异步做完任务后，释放资源
	for i := 0; i < chanSize; i++ {
		go func() {
			//dosomething
			syncChan <- 1
		}()
	}

	// 消费者，实现同步等待
	// 一开始取不到则挂起
	// 直到chanSize个资源全部获取完毕
	for i := 0; i < chanSize; i++ {
		<-syncChan
	}

	// after同步后的操作
}
