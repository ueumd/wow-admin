package utils

import "sync"

type WaitGroup struct {
	sync.WaitGroup
}

func (w *WaitGroup) Wrap(f func()){
	w.Add(1)
	go func() {
		defer w.Done()
		f()
	}()
}
