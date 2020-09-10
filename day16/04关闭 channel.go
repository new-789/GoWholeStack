package main

import (
	"fmt"
	"time"
)

func main0401() {
	CacheCh := make(chan int, 3)
	// 有缓冲 channel 关闭后读数据测试
	go func() {
		for i := 0; i < 3; i++ {
			CacheCh <- i
		}
		close(CacheCh)
		fmt.Println("在有缓冲 channel 测试中已经关闭了 channel")
	}()

	time.Sleep(time.Second * 2)
	/*
	   for {
	   		// 从 channel 中获取数据并判断 channel 是否关闭
	   		if value, ok := <-CacheCh; ok {
	   			fmt.Println("主 go 程测试有缓冲 channel 读到数据:", value)
	   		} else {
	   			n := <-CacheCh
	   			fmt.Println("有缓冲 channel 关闭后读到数据:", n)
	   			break
	   		}
	   	}
	*/

	// 通过遍历获取 channel 中的数据防止阻塞
	for num := range CacheCh {
		fmt.Println("主 go 程测试有缓冲 channel 读到数据:", num)
	}

	/*
		// 无缓冲 channel 关闭后读数据测试
		NoCacheCh := make(chan int)
		go func() {
			for i := 0;i < 2 ;i++  {
				NoCacheCh <- i
			}
			close(NoCacheCh)
			fmt.Println("在无缓冲 channel 测试中已经关闭了 channel")
		}()

		time.Sleep(time.Second * 2)
		for {
			if value, ok := <- NoCacheCh; ok {
				fmt.Println("主 go 程测试无缓冲 channel 读到数据:", value)
			} else {
				n := <- NoCacheCh
				fmt.Println("无缓冲 channel 关闭后读到数据:", n)
				break
			}
		}
	*/
}
