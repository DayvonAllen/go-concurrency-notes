package main

import (
	"com.example/app/cancel"
	"fmt"
	"runtime"
	"time"
)

func main() {
	//channels.UnbufferedChannel()
	//channels.RangeChannel()
	//channels.BufferedChannel()
	//channels.ChannelDirection()
	//channels.ChannelOwnership()
	//select_ex.FirstSelect()
	//select_ex.SecondSelect()
	//select_ex.ThirdSelect()
	//sync_ex.MutexEx()
	//sync_ex.AtomicEx()
	//sync_ex.CondEx()
	//sync_ex.PoolEx()
	//race.RaceEx()
	//pipeline.FirstPipeline()
	//pipeline.Fan()
	cancel.CancelGoRou()
	time.Sleep(10 * time.Millisecond)

	g := runtime.NumGoroutine()
	fmt.Println(g)
}