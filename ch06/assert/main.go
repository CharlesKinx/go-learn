package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func add32(a, b int32) int32 {
	return a + b
}

func add64(a, b int64) int64 {
	return a + b
}

func addInter(a, b interface{}) int {
	ai, _ := a.(int)
	bi, _ := b.(int)
	return ai + bi
}

// 接口嵌套

type MyWriter interface {
	Writer(string)
}

type MyReader interface {
	Reader() string
}

type MyWriterAndReader interface {
	MyWriter
	MyReader
	ReadWriter()
}

type SreadWriter struct {
}

func (s SreadWriter) Writer(s2 string) {
	//TODO implement me
	fmt.Println("implement me")
}

func (s SreadWriter) Reader() string {
	//TODO implement me
	panic("implement me")
}

func (s SreadWriter) ReadWriter() {
	//TODO implement me
	panic("implement me")
}

func main() {

	var srw MyWriterAndReader = SreadWriter{}
	srw.Writer("qaz")
}
