package main

import (
	"os"
	"sync"
	"errors"
	"fmt"
	"path/filepath"
	"math/rand"
	"sync/atomic"
)

type Data []byte

type DataFile interface {
	Read() (rsn int64, d Data, err error)
	Write(d Data) (wsn int64, err error)
	RSN() int64
	WSN() int64
	DataLen() uint32
	Close() error
}

type myDataFile struct {
	f       *os.File
	fmutex  sync.RWMutex
	woffset int64
	roffset int64
	wmutex  sync.Mutex
	rmutex  sync.Mutex
	dataLen uint32
}

func newDataFile(path string, dataLen uint32) (DataFile, error) {

	f, err := os.Create(path)

	if err != nil {
		return nil, err
	}

	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}

	df := &myDataFile{f: f, dataLen: dataLen}

	return df, nil
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {

	var offset int64

	df.rmutex.Lock()

	offset = df.roffset
	df.roffset += int64(df.dataLen)

	df.rmutex.Unlock()

	rsn = offset / int64(df.dataLen)
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()

	bytes := make([]byte, df.dataLen)
	_, err = df.f.ReadAt(bytes, offset)

	if err != nil {
		return
	}

	d = bytes
	return
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {

	var offset int64

	df.wmutex.Lock()

	offset = df.woffset
	df.woffset += int64(df.dataLen)

	df.wmutex.Unlock()

	wsn = offset / int64(df.dataLen)

	var bytes []byte

	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}

	df.fmutex.Lock()
	defer df.fmutex.Unlock()

	_, err = df.f.Write(bytes)

	return
}

func (df *myDataFile) RSN() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) WSN() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func (df *myDataFile) Close() error {

	if df.f == nil {
		return nil
	}

	return df.f.Close()
}

func main() {

	path := filepath.Join(os.TempDir(), "data_file.txt")

	fmt.Println(path, 3)

	df, err := newDataFile(path, 3)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(df.WSN(), df.RSN(), df.DataLen())

	data := Data{byte(rand.Int31n(256)), byte(rand.Int31n(256)), byte(rand.Int31n(256)), }

	df.Write(data)

	fmt.Println(df.WSN(), df.RSN(), df.DataLen())

	//df.Close()

	rsn, data, err := df.Read()

	fmt.Println(rsn, data, err)

	var i32 uint32

	atomic.AddUint32(&i32, 3)

	fmt.Println(i32)

	//一个负整数的补码可以通过对它按位求反并加一得到
	atomic.AddUint32(&i32, ^uint32(1 - 1))

	fmt.Println(i32)

	//fmt.Println(i32)

	//fmt.Println(uint32(int32(10)))
	//fmt.Println(^uint32(-10 - 1))

	var a int32

	a = 1000

	addValue(&a, 100)

	fmt.Println( a )

}

func addValue(value *int32, delta int32) {

	for {
		v := &value
		if atomic.CompareAndSwapInt32(value, v, v+delta) {
			break;
		}
	}
}
