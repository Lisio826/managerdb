package utils

import "github.com/panjf2000/ants"

var TPool *ants.Pool

func init() {
	TPool, _ := ants.NewPool(100)
	defer TPool.Release()
}
