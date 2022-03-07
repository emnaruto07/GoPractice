package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

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
	cc     int
	volume int
}

func (t *TataPlayRemote) VolumeUp() {
	t.volume += 1
}

func (t *TataPlayRemote) VolumeDown() {
	t.volume -= 1
}

func (t *TataPlayRemote) SwitchChannel(number int) {
	t.cc = number
}

func (t *TataPlayRemote) CurrentChannel() int {
	return t.cc
}

type TV struct {
	r Remote // interface
}

// package io
type Reader interface {
	Read() ([]byte, int)
}

// os, net, net/http
// File, TCPConn, http.Request

type Writer interface {
	Write([]byte) int
}

func work(r io.Reader) {

}

type Adder interface {
	Add(int) int
}

func Add5(a Adder) {
	a.Add(5)
}

type Num struct {
	n int
}

func (num *Num) String() string {
	return fmt.Sprintf("Num=%d", num.n)
}

func (num *Num) Add(input int) int {
	num.n += input
	return num.n
}

func (i *int) Add(k int) int {}

// var remote Remote = TataPlayRemote{100, 0}
// remote.CurrentChannel()
// remote.VolumeUp()
// fmt.Println("Test")

// number
// var num Adder = &Num{0}
// Add5(num)
// k, ok := num.(*int) // type conversion
// fmt.Println(k.n)

// fmt.Println(k)

func typeof(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("fasdfas")
	case string:
		fmt.Println("string")
	}
}

func main() {
	typeof("sdafdsa") // string
	typeof(int(5))    // int
}
