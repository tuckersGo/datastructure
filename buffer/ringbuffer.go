package buffer

type RingBuffer[T any] struct {
	buf     []T
	readPt  int
	writePt int
	isFull  bool
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buf: make([]T, size),
	}
}

func (r *RingBuffer[T]) Write(data []T) int {
	if len(data) == 0 || r.Writable() == 0 {
		return 0
	}

	var writed int
	if r.writePt >= r.readPt {
		writableToEnd := len(r.buf) - r.writePt
		writed = min(writableToEnd, len(data))
	} else {
		writed = min(r.Writable(), len(data))
	}

	copy(r.buf[r.writePt:], data[:writed])
	r.writePt += writed
	r.writePt %= len(r.buf)

	// isFull?
	if writed > 0 && r.writePt == r.readPt {
		r.isFull = true
	}

	remain := len(data) - writed
	if remain > 0 && r.Writable() > 0 {
		writed += r.Write(data[writed:])
	}

	return writed
}

func (r *RingBuffer[T]) Readable() int {
	if r.isFull {
		return len(r.buf)
	}
	if r.writePt < r.readPt {
		return len(r.buf) - r.readPt + r.writePt
	}
	return r.writePt - r.readPt
}

func (r *RingBuffer[T]) Writable() int {
	return len(r.buf) - r.Readable()
}

func (r *RingBuffer[T]) Read(count int) []T {
	if r.Readable() == 0 || count <= 0 {
		return []T{}
	}
	readCnt := min(count, r.Readable())
	rst := make([]T, readCnt)

	if r.readPt+readCnt >= len(r.buf) {
		remainUntilEnd := len(r.buf) - r.readPt
		copy(rst, r.buf[r.readPt:])
		r.readPt = 0

		remain := readCnt - remainUntilEnd
		copy(rst[remainUntilEnd:], r.buf[:remain])
		r.readPt += remain
	} else {
		copy(rst, r.buf[r.readPt:r.readPt+readCnt])
		r.readPt += readCnt
	}
	r.isFull = false
	return rst
}
