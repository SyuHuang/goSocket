package utility

import (
	"bytes"
)

//BinaryReader :
type BinaryReader struct {
	_reader *bytes.Reader
	_offset int
}

//Init :
func (p *BinaryReader) Init(bs []byte) {
	p._reader = bytes.NewReader(bs)
	p._offset = 0
}

//ReadInt16 :
func (p *BinaryReader) ReadInt16() int {
	bs := make([]byte, 2)
	_, err := p._reader.ReadAt(bs, int64(p._offset))
	if err != nil {
		return 0
	}
	p._offset += 2
	return BytesToInt16(bs)
}

//ReadInt :
func (p *BinaryReader) ReadInt() int {
	bs := make([]byte, 4)
	_, err := p._reader.ReadAt(bs, int64(p._offset))
	if err != nil {
		return 0
	}
	p._offset += 4
	return BytesToInt(bs)
}

//ReadUInt :
func (p *BinaryReader) ReadUInt() uint32 {
	bs := make([]byte, 4)
	_, err := p._reader.ReadAt(bs, int64(p._offset))
	if err != nil {
		return 0
	}
	p._offset += 4
	return BytesToUInt(bs)
}

//ReadBytes :
func (p *BinaryReader) ReadBytes(length int) []byte {
	if length <= 0 {
		return nil
	}
	bs := make([]byte, length)
	_, err := p._reader.ReadAt(bs, int64(p._offset))
	if err != nil {
		return nil
	}
	p._offset += length
	return bs
}

//ReadString :
func (p *BinaryReader) ReadString() string {
	length := p.ReadInt()
	if length <= 0 {
		return ""
	}

	bs := p.ReadBytes(length)
	return BytesToString(bs)
}

//NewBinaryReader :Construtor for BinaryReader
func NewBinaryReader(bs []byte) *BinaryReader {
	reader := &BinaryReader{}
	reader.Init(bs)
	return reader
}
