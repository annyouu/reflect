package main

import (
	"fmt"
	"reflect"
)

func main() {
	// int型のチャネルを作成
	ch1 := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, reflect.TypeOf(0)), 1)
	ch2 := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, reflect.TypeOf(0)), 1)

	// 送信
	ch1.Send(reflect.ValueOf(10))
	ch2.Send(reflect.ValueOf(20))

	// Selectcaseの定義(受信を行う)
	cases := []reflect.SelectCase{
		{Dir: reflect.SelectRecv, Chan: ch1},
		{Dir: reflect.SelectRecv, Chan: ch2}, // case val := <-ch2:
		{Dir: reflect.SelectDefault},         // default:
	}

	// Selectを実行
	chosen, recv, recvOK := reflect.Select(cases)

	if recvOK {
		fmt.Printf("チャネル%dから受信: %d\n", chosen, recv.Int())
	} else {
		fmt.Println("デフォルト case が選ばれました")
	}
}