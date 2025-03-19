package main

import (
	"fmt"
	"reflect"
)

func main() {
	// チャネルを作成
	ch1, ch2 := make(chan int), make(chan int)

	// ゴルーチンを使って非同期で値を送信
	go func() {
		ch1 <- 100  // ch1に100を送信
	}()
	
	go func() {
		ch2 <- 200  // ch2に200を送信
	}()

	// reflect.Selectを使用して、複数のチャネルを操作する
	i, v, ok := reflect.Select([]reflect.SelectCase{
		// case: <-ch1
		{
			Dir: reflect.SelectRecv,
			Chan: reflect.ValueOf(ch1),
			Send: reflect.ValueOf(nil),
		},
		// case: ch2 <- 100
		{
			Dir: reflect.SelectSend,
			Chan: reflect.ValueOf(ch2),
			Send: reflect.ValueOf(100),
		},
		// default case
		{
			Dir: reflect.SelectDefault,
			Chan: reflect.ValueOf(nil),
			Send: reflect.ValueOf(nil),
		},
	})

	// 結果の出力
	if ok {
		fmt.Printf("選ばれたケースindex: %d\n", i)
		if i == 0 { // ch1から受信
			fmt.Printf("ch1から受信した値: %d\n", v.Int())
		} else if i == 1 { // ch2に送信
			fmt.Printf("ch2に送信した値: %d\n", v.Int())
		}
	} else {
		fmt.Println("デフォルトケースが選ばれました")
	}
}