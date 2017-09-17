package main

import "github.com/mrmiguu/rest"

func main() {
	rest.Connect("http://127.0.0.1/")
	h := rest.New("login")
	w, r := h.Int()
	w(69)
	println(r())
	select {}
}

// package main

// import (
// 	"bytes"
// 	"encoding/binary"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	w, r := String()
// 	w("lowercase")
// 	println(r())
// }

// const (
// 	tbytes byte = iota
// 	tstring
// 	tint
// )

// func write(t byte, b []byte) {
// 	b = append([]byte{t}, b...)
// 	http.Post("http://127.0.0.1/post", "text/plain", bytes.NewReader(b))
// }

// func read(t byte) []byte {
// 	b := []byte{t}
// 	resp, err := http.Post("http://127.0.0.1/get", "text/plain", bytes.NewReader(b))
// 	if err != nil {
// 		panic(err)
// 	}
// 	b, err = ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return b
// }

// func Bytes() (func([]byte), func() []byte) {
// 	w := func(b []byte) {
// 		write(tbytes, b)
// 	}
// 	r := func() []byte {
// 		return read(tbytes)
// 	}
// 	return w, r
// }

// func String() (func(string), func() string) {
// 	w := func(s string) {
// 		write(tstring, []byte(s))
// 	}
// 	r := func() string {
// 		return string(read(tstring))
// 	}
// 	return w, r
// }

// func Int() (func(int), func() int) {
// 	w := func(i int) {
// 		b := make([]byte, 8)
// 		binary.BigEndian.PutUint64(b, uint64(i))
// 		write(tint, b)
// 	}
// 	r := func() int {
// 		return int(binary.BigEndian.Uint64(read(tint)))
// 	}
// 	return w, r
// }
