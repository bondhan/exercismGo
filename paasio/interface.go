package paasio

import (
	"io"
	"sync"
)

// ReadCounter is an interface describing objects that can be read from,
// and that can count the number of times they have been read from.
//
// If multiple goroutines concurrently call Read, implementations are not
// required to provide any guarantees about interleaving of the Read calls.
// However, implementations MUST guarantee that calls to ReadCount always return
// correct results even in the presence of concurrent Read calls.
type ReadCounter interface {
	io.Reader
	// ReadCount returns the total number of bytes successfully read along
	// with the total number of calls to Read().
	ReadCount() (n int64, nops int)
}

type readCounter struct {
	rb  io.Reader
	mux sync.Mutex
	rc  int64
	ops int
}

func (r *readCounter) ReadCount() (int64, int) {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.rc, r.ops
}

func (r *readCounter) Read(data []byte) (n int, err error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	n, err = r.rb.Read(data)
	r.rc += int64(n)
	r.ops++

	return n, err
}

//NewReadCounter ...
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{rb: r, rc: 0, ops: 0}
}

// WriteCounter is an interface describing objects that can be written to,
// and that can count the number of times they have been written to.
//
// If multiple goroutines concurrently call Write, implementations are not
// required to provide any guarantees about interleaving of the Write calls.
// However, implementations MUST guarantee that calls to WriteCount always return
// correct results even in the presence of concurrent Write calls.
type WriteCounter interface {
	io.Writer
	// WriteCount returns the total number of bytes successfully written along
	// with the total number of calls to Write().
	WriteCount() (n int64, nops int)
}

type writeCounter struct {
	wb  io.Writer
	mux sync.Mutex
	wc  int64
	ops int
}

func (w *writeCounter) WriteCount() (n int64, nops int) {
	w.mux.Lock()
	defer w.mux.Unlock()
	return w.wc, w.ops
}

func (w *writeCounter) Write(p []byte) (n int, err error) {
	w.mux.Lock()
	defer w.mux.Unlock()

	n, err = w.wb.Write(p)
	w.wc += int64(n)
	w.ops++

	return n, err
}

// NewWriteCounter ...
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{wb: w, wc: 0, ops: 0}
}

// ReadWriteCounter is the union of ReadCounter and WriteCounter.
//
// All guarantees that apply to either of ReadCounter or WriteCounter
// also apply to ReadWriteCounter.
type ReadWriteCounter interface {
	ReadCounter
	WriteCounter
}

type readWriteCounter struct {
	readCounter
	writeCounter
}

// NewReadWriteCounter ...
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{
		readCounter:  readCounter{rb: rw, rc: 0, ops: 0},
		writeCounter: writeCounter{wb: rw, wc: 0, ops: 0},
	}
}
