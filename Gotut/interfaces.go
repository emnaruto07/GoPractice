package main

import (
	"fmt"
)

// type NumberRemote int

// func (NumberRemote) VolumeUp() {}
// func (NumberRemote) VolumeDown() {}

// func (NumberRemote) SwitchChannel(int) {}
// func (NumberRemote) CurrentChannel() int {}

type Remote interface {
	VolumeUp()
	VolumeDown()
	SwitchChannel(int) // 200, 300, 12
	CurrentChannel() int
}

type TataPlayRemote struct {
	currentChannel int
	volume         int
}

func (t *TataPlayRemote) VolumeUp() {
	t.volume += 1
}

func (t *TataPlayRemote) VolumeDown() {
	t.volume -= 1
}

func (t *TataPlayRemote) SwitchChannel(number int) {
	t.currentChannel = number
}

func (t *TataPlayRemote) CurrentChannel() int {
	return t.currentChannel
}

func main() {
	var remote Remote = NumberRemote(200)
	remote.CurrentChannel()
	remote.VolumeUp()
	fmt.Println("Test")
}
