package main

import (
	"context"
	"fmt"
	"time"
)

/**
基于 Channel 编写一个简单的单线程生产者消费者模型：
1. 队列： 队列长度 10，队列元素类型为 int
2. 生产者： 每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
3. 消费者： 每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
*/
func main() {
	message := make(chan int, 10)
	//done := make(chan bool)
	defer close(message)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// producer
	go Producer(message, ctx)
	go Consumer(message, ctx)

	time.Sleep(12 * time.Second)
	fmt.Println("main process exit!")
}

func Producer(message chan int, ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	i := 0
	for _ = range ticker.C {
		select {
		case <-ctx.Done():
			fmt.Println("生产完毕！")
			return
		default:
			fmt.Printf("生产产品: %d\n", i)
			message <- i
			i = i + 1
		}

	}
}

func Consumer(message chan int, ctx context.Context) {
	ticker := time.NewTicker(time.Second)
	for _ = range ticker.C {
		select {
		case <-ctx.Done():
			fmt.Println("消费完毕！")
			return
		default:
			m := <-message
			fmt.Printf("消费产品：%d\n", m)
		}
	}
}
