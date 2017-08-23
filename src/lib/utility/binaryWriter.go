package utility

import (
	"bufio"
	"bytes"
	"fmt"
)

//BinaryWriter :
type BinaryWriter struct {
	_buffer bytes.Buffer
	_writer *bufio.Writer
}

//Init :
func (p *BinaryWriter) Init() {
	p._writer = bufio.NewWriter(&p._buffer)
}

//ToBytes :
func (p *BinaryWriter) ToBytes() []byte {
	p._writer.Flush()
	return p._buffer.Bytes()
}

//WriteInt16 :
func (p *BinaryWriter) WriteInt16(value int16) {
	bs := Int16ToBytes(value)
	p._writer.Write(bs)
}

//WriteInt :
func (p *BinaryWriter) WriteInt(value int) {
	bs := IntToBytes(value)
	p._writer.Write(bs)
}

//WriteUInt :
func (p *BinaryWriter) WriteUInt(value uint32) {
	bs := UIntToBytes(value)
	p._writer.Write(bs)
}

//WriteBytes :
func (p *BinaryWriter) WriteBytes(bytes []byte) {
	length := len(bytes)
	p.WriteInt(length)
	p._writer.Write(bytes)
}

//WriteString :
func (p *BinaryWriter) WriteString(str string) {
	tmp := StringToBytes(str)
	p.WriteBytes(tmp)
}

//Write :
func (p *BinaryWriter) Write(value interface{}) {
	switch value.(type) {
	default:
		fmt.Println("Not supported type.")
	case int16:
		p.WriteInt16(value.(int16))
	case int:
		p.WriteInt(value.(int))
	case uint32:
		p.WriteUInt(value.(uint32))
	case []byte:
		p.WriteBytes(value.([]byte))
	case string:
		p.WriteString(value.(string))
	}
}

//NewBinaryWriter :Construtor for BinaryWriter
func NewBinaryWriter() *BinaryWriter {
	writer := &BinaryWriter{}
	writer.Init()
	return writer
}
