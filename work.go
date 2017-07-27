package main

import (
	"sync"
	"log"
	"time"
)

//必须实现的接口
type Worker interface {
	Task()
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func NewPool(maxGoroutines int) *Pool {

	p := &Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return p
}

//提交工作到工作池
func (p *Pool) Run(w Worker) {
	p.work <- w
}

//停止所有工作
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}

// ================ test

var names = []string{
	"aaa", "bbb", "ccc", "dddl",
}

type namePrinter struct {
	name string
}

func (p *namePrinter) Task() {
	log.Println("name: ", p.name)
	time.Sleep(time.Second)
}

func main() {

	//工作池数量为2
	p := NewPool(100)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {

		for _, name := range names {

			nameprint := &namePrinter{name: name }

			go func() {
				p.Run(nameprint)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	//关闭工作池
	p.Shutdown()
}
