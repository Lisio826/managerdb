package utils

import (
	"runtime"
)

// 获取正在运行的函数名
func RunFuncName()string{
	pc := make([]uintptr,1)
	runtime.Callers(2,pc)
	f := runtime.FuncForPC(pc[0])
	//fmt.Println(f.Name())
	return f.Name()
}
