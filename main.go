package main

import (
	"com.example/app/channels"
	"com.example/app/select_ex"
)

func main() {
	channels.UnbufferedChannel()
	channels.RangeChannel()
	channels.BufferedChannel()
	channels.ChannelDirection()
	channels.ChannelOwnership()
	select_ex.FirstSelect()
	select_ex.SecondSelect()
	select_ex.ThirdSelect()
}