package test

import (
	"fmt"
	"runtime"
	"testing"
)

// 获取正在运行的函数名
func runFuncName()string{
	pc := make([]uintptr,1)
	runtime.Callers(2,pc)
	f := runtime.FuncForPC(pc[0])
	fmt.Println(f.Name())
	return f.Name()
}

func Test1T(tt *testing.T)  {
	testFun1()
	fmt.Println("--------------")
	testFun2()
}

func testFun1()  {
	a := runFuncName()
	fmt.Println(a)
}
func testFun2()  {
	b := runFuncName()
	fmt.Println(b)
}
