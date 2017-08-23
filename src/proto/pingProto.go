//Highly recommand to make a proto generator to gen this file.

package proto

import (
	util "../lib/utility"
)

//PingProto :
type PingProto struct {
	_id int
}

//FromBytes :
func (p *PingProto) FromBytes(reader *util.BinaryReader) {
	p._id = reader.ReadInt()
}

//ToBytes :
func (p *PingProto) ToBytes(writer *util.BinaryWriter) {
	writer.WriteInt(p._id)
}
