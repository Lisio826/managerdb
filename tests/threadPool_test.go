package test

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
	"testing"
	"time"
)

func a(i interface{}) {
	fmt.Println("===============")
	fmt.Println(i)
	fmt.Println("---------------")
	time.Sleep(1 * time.Second)
}
func TestPools(t *testing.T){
	pool,_ := ants.NewPool(10)
	var wg sync.WaitGroup
	syncCalculateSum := func(i interface{}) {
		a(i)
		wg.Done()
		fmt.Println("++++++++++")
	}

	syncs:= func() {
		wg.Done()
		fmt.Println("++++++++++")
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		fmt.Println("^^^^^^^^^^^^^^")
		pool.Submit(syncs)
		fmt.Println(ants.Cap())
		fmt.Println(ants.Running())
	}


	p,_ := ants.NewPoolWithFunc(2,syncCalculateSum)
	defer p.Release()
	for i := 0; i < 20; i++ {
		wg.Add(1)
		fmt.Println("^^^^^^^^^^^^^^")
		p.Invoke(i)
		fmt.Println(ants.Cap())
		fmt.Println(ants.Running())
	}
	wg.Wait()

	fmt.Println("end")
}
