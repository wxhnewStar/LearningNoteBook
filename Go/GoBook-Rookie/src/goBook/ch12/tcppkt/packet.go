package main

import (
	"bytes"
	"encoding/binary"
	"io"
)

/* ------- 封包格式及发送 ------- */

// Packet 自定义的二进制封包格式
type Packet struct {
	Size uint16
	Body []byte
}

// 将数据写入 dataWriter
func writePacket(dataWriter io.Writer, data []byte) error {

	// 准备一个字节数组缓冲
	var buffer bytes.Buffer // bytes 的 buffer 是自增长的

	// 将 Size 写入缓冲, 使用 binary 的可定义大小段的 Write 方法写入长度
	if err := binary.Write(&buffer, binary.LittleEndian, uint16(len(data))); err != nil {
		return err
	}

	// 写入包体
	if _, err := buffer.Write(data); err != nil {
		return err
	}

	// 获得写入的完整数据
	out := buffer.Bytes()

	// 写入 dataWriter
	if _, err := dataWriter.Write(out); err != nil {
		return err
	}

	return nil
}

/* ------- 封包读取 ------- */

// 从 dataReader 中读取封包
func readPacket(data io.Reader) (pkt Packet, err error) {

	// Size 为 unit16 类型,占两个字节
	var sizeBuffer = make([]byte, 2)

	// 持续读取 Size 直到读到为止
	_, err = io.ReadFull(data, sizeBuffer)

	// 发生错误时返回
	if err != nil {
		return
	}

	// 使用 bytes.Reader 读取 sizeBuffer 中的数据
	sizeReader := bytes.NewReader(sizeBuffer)

	// 读取小端的 unit16 作为size
	err = binary.Read(sizeReader, binary.LittleEndian, &pkt.Size)

	if err != nil {
		return
	}

	// 分配包体大小
	pkt.Body = make([]byte, pkt.Size)

	// 读取包体数据
	_, err = io.ReadFull(data, pkt.Body)

	return
}
