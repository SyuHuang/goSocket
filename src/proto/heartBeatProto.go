//Highly recommand to make a proto generator to gen this file.

package proto

import (
	util "../lib/utility"
)

//HeartBeatProto :
type HeartBeatProto struct {
	_id   int
	_name string
}

//FromBytes :
func (p *HeartBeatProto) FromBytes(reader *util.BinaryReader) {
	p._id = reader.ReadInt()
	p._name = reader.ReadString()
}

//ToBytes :
func (p *HeartBeatProto) ToBytes(writer *util.BinaryWriter) {
	writer.WriteInt(p._id)
	writer.WriteString(p._name)
}
