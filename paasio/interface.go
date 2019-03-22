package paasio

import "io"

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
}

func (r *readCounter) ReadCount() (int64, int) {
	return 0, 0
}

func (r *readCounter) Read(data []byte) (int, error) {
	return 0, nil
}

//NewReadCounter ...
func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{}

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
	w io.Writer
	// counter
}

func (w *writeCounter) WriteCount() (n int64, nops int) {
	return 0, 0
}

func (w *writeCounter) Write(p []byte) (n int, err error) {

	return len(p), nil
}

// NewWriteCounter ...
func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{}
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
		readCounter:  readCounter{},
		writeCounter: writeCounter{},
	}
}
