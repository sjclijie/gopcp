package main

import (
	"sync"
	"io"
	"errors"
	"log"
	"sync/atomic"
	"time"
	"math/rand"
)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrPoolClosed = errors.New("Pool has been closed.")

func NewPool(fn func() (io.Closer, error), size int) (*Pool, error) {

	if size <= 0 {
		return nil, errors.New("Size value too small.")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Acquire: ", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire: ", "New Resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {

	p.m.Lock()

	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Release: ", "In Queue")
	default:
		//队列已满
		log.Println("Release: ", "Closing")
		r.Close()
	}
}

func (p *Pool) Close() {

	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}

// ========= test

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Printf("Close: Connection %d", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{idCounter }, nil
}

func main() {

	maxGoroutines := 2500
	pooledResources := 500

	//test
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	p, err := NewPool(createConnection, pooledResources)

	if err != nil {
		log.Println(err)
	}

	for query := 0; query < maxGoroutines; query++ {
		go func() {

			conn, err := p.Acquire()

			if err != nil {
				log.Println(err)
				return
			}

			defer p.Release(conn)

			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

			wg.Done()

		}()
	}

	wg.Wait()

	p.Close()
}
