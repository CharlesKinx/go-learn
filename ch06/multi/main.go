package main

/*
* 多接口
 */
type MyWriter interface {
	WriteInter()
}

type MyClose interface {
	CloseInter()
}

type myWriter struct {
	name string
}

func (mw *myWriter) WriteInter() {

}
func (mw *myWriter) CloseInter() {

}
