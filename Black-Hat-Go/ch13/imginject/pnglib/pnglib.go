package pnglib

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"io"
	"log"
	"os"
)

const PNGEndChunkType = "IEND"

type PNGHeader [8]byte

type PNGChunk struct {
	Size uint32
	Type [4]byte
	Data []byte
	CRC  [4]byte
}

func (c *PNGChunk) marshal() *bytes.Buffer {
	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, c.Size); err != nil {
		log.Fatalln(err)
	}
	buf.Write(c.Type[:])
	buf.Write(c.Data)
	buf.Write(c.CRC[:])
	return buf
}

type PNGInject struct {
	byteReader *bytes.Reader
}

// 验证 header，确定为 PNG 文件
func (r *PNGInject) ValidateHeader() (bool, error) {
	var header PNGHeader
	_, err := r.byteReader.Read(header[:])
	if err != nil {
		return false, err
	}
	return string(header[1:4]) == "PNG", nil
}

// 读取 chunk
func (r *PNGInject) ReadChunk(c *PNGChunk) {
	r.ReadChunkSize(c)
	r.ReadChunkType(c)
	r.ReadChunkData(c, c.Size)
	r.ReadChunkCRC(c)
}

// 读取 chunk 的 4 字节 size
func (r *PNGInject) ReadChunkSize(c *PNGChunk) {
	if err := binary.Read(r.byteReader, binary.BigEndian, &(c.Size)); err != nil {
		log.Fatalln(err)
	}
}

// 读取 chunk 的 4 字节 type
func (r *PNGInject) ReadChunkType(c *PNGChunk) {
	if _, err := r.byteReader.Read(c.Type[:]); err != nil {
		log.Fatalln(err)
	}
}

// 读取 chunk 的 size 字节的 data
func (r *PNGInject) ReadChunkData(c *PNGChunk, size uint32) {
	c.Data = make([]byte, size)
	if _, err := r.byteReader.Read(c.Data); err != nil {
		log.Fatalln(err)
	}
}

// 读取 chunk 的 4 字节 CRC
func (r *PNGInject) ReadChunkCRC(c *PNGChunk) {
	if _, err := r.byteReader.Read(c.CRC[:]); err != nil {
		log.Fatalln(err)
	}
}

// 获取 offset
func (r *PNGInject) GetOffset() int {
	offset, _ := r.byteReader.Seek(0, 1)
	return int(offset)
}

// 获取 IEND chunk 和 offset 位置
func (r *PNGInject) ReadIENDChunk(c *PNGChunk) (offset int) {
	// 读取 header 并验证
	valid, err := r.ValidateHeader()
	if err != nil {
		log.Fatalln(err)
	}
	if !valid {
		log.Fatalln("Provided file is not a valid PNG format")
	}

	for string(c.Type[:]) != PNGEndChunkType {
		offset = r.GetOffset()
		r.ReadChunk(c)
	}
	return offset
}

// 在 offset 位置写入 IEND chunk，并另存为新文件
func (r *PNGInject) WriteIENDChunk(offset int, data *[]byte) *bytes.Buffer {
	chunk, err := NewPNGChunk(data, PNGEndChunkType)
	if err != nil {
		log.Fatalln(err)
	}

	originalContent := make([]byte, offset)
	r.byteReader.Seek(0, 0)
	r.byteReader.Read(originalContent)

	buf := new(bytes.Buffer)
	buf.Write(originalContent)
	buf.Write(chunk.marshal().Bytes())

	// 将源文件 IEND 块之后的内容写入新文件
	_, err = io.Copy(buf, r.byteReader)
	if err != nil {
		log.Fatalln(err)
	}

	return buf
}

func NewPNGInject(file *os.File) (*PNGInject, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileContent := make([]byte, fileInfo.Size())
	_, err = file.Read(fileContent)
	if err != nil {
		return nil, err
	}

	return &PNGInject{bytes.NewReader(fileContent)}, nil
}

func NewPNGChunk(data *[]byte, t string) (*PNGChunk, error) {
	chunk := PNGChunk{
		Size: uint32(len(*data)),
		Data: *data,
	}
	copy(chunk.Type[:], t)

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, chunk.Type); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, chunk.Data); err != nil {
		return nil, err
	}
	binary.BigEndian.PutUint32(chunk.CRC[:], crc32.ChecksumIEEE(buf.Bytes()))

	return &chunk, nil
}
