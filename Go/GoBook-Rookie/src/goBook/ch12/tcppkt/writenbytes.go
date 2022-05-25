package main

import "io"

func ReadAtLeast(r io.Reader, buf []byte, min int) (n int, err error) {
	if len(buf) < min {
		return 0, io.ErrShortBuffer
	}

	// 当没读够且没出错时一直读取
	for n < min && err == nil {
		var nn int
		nn, err = r.Read(buf[n:]) // n 记录着当前读取的位置
		n += nn
	}

	// 如果 buf 读取不下了,但是大于 min 退出咋整?

	if n >= min {
		err = nil
	} else if n > 0 && err == io.EOF {
		err = io.ErrUnexpectedEOF
	}
	return
}

func ReadFull(r io.Reader, buf []byte) (n int, err error) {
	return ReadAtLeast(r, buf, len(buf))
}
