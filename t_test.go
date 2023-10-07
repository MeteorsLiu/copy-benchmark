package ttt

import (
	"testing"
	"time"
)


const (
	GB = 1024*1024*1024
)

func copy_rep(src, dst *byte, n int) 

func copy_req(src, dst *byte, n int) 


func copy_avx(src, dst *byte, n int) 


func getLog(t *testing.T , name string, written uint64, last time.Time) {
	b := written / uint64(time.Since(last).Seconds()) 
	t.Log(name + " Output: ",b / GB, "Gb/s ", b, "B/s")
}

func TestRep(t *testing.T) {
	src := make([]byte, 32768)
	for i := 0; i < len(src); i++ {
		src[i] = 1
	}
	dst := make([]byte, 32768)

	copy_rep(&dst[0], &src[0], 32768)

	t.Log(dst)
}

func TestReq(t *testing.T) {
	src := make([]byte, 32768)
	for i := 0; i < len(src); i++ {
		src[i] = 1
	}
	dst := make([]byte, 32768)

	copy_req(&dst[0], &src[0], 32768)

	t.Log(dst)
}

func TestOutput(t *testing.T) {
	src := make([]byte, 32768)
	for i := 0; i < len(src); i++ {
		src[i] = 1
	}
	dst := make([]byte, 32768)
	zero :=  make([]byte, 32768)
	written := uint64(0)
	now := time.Now()
	
	for time.Since(now).Seconds() < 11 {
		copy_rep(&dst[0], &src[0], 32768)
		written += 32768

		copy_rep(&dst[0], &zero[0], 32768)
		written += 32768

	}

	getLog(t, "REP SB", written, now)

	clear(dst)
	written = uint64(0)
	now = time.Now()
	
	for time.Since(now).Seconds() < 11 {
		copy_req(&dst[0], &src[0], 32768)
		written += 32768

		copy_req(&dst[0], &zero[0], 32768)
		written += 32768
	}

	getLog(t, "REP SQ", written, now)

	clear(dst)
	written = uint64(0)
	now = time.Now()
	
	for time.Since(now).Seconds() < 11 {
		written += uint64(copy(dst, src))

		written += uint64(copy(dst, zero))
	}

	getLog(t, "copy", written, now)

}

func BenchmarkREPSB(b *testing.B) {
	src := make([]byte, 32768)
	for i := 0; i < len(src); i++ {
		src[i] = 1
	}
	dst := make([]byte, 32768)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		copy_rep(&dst[0], &src[0], 32768)
	}
}

func BenchmarkREPSQ(b *testing.B) {
	src := make([]byte, 32768)
	for i := 0; i < len(src); i++ {
		src[i] = 1
	}
	dst := make([]byte, 32768)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		copy_req(&dst[0], &src[0], 32768)
	}
}


func BenchmarkCopy(b *testing.B) {
	src := make([]byte, 32768)
	for i := 0; i < len(src); i++ {
		src[i] = 1
	}
	dst := make([]byte, 32768)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		copy(dst, src)
	}
}