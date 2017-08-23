package network

import (
	"fmt"
	"hash/crc32"
	"net"

	util "../utility"
)

const (
	waitHead    = 1 //wait for the packet header
	waitPayload = 2 //wait for the packet payload
	waitCRC     = 3 //wait for the packet crc checksum
)

//HeaderSize packet header size
const HeaderSize = 4

//TailCrcSize packet crc checksum size
const TailCrcSize = 4

//PacketHelper :Promise the packet doesn't miss it's payload by header and crc check.
type PacketHelper struct {
	_conn        net.Conn
	_payloadSize int //
}

//Read :It's stepping from Head to Crc. Before it's done, this function will stay in a blocking status.
func (p *PacketHelper) Read() ([]byte, int, error) {

	recvBuff := []byte{}
	curState := waitHead

	for {
		switch curState {
		case waitHead:
			{
				buf, err := ReadLimitedLength(p._conn, HeaderSize)
				if err != nil {
					return nil, 0, err
				}

				p._payloadSize = util.BytesToInt(buf)
				curState = waitPayload
			}
		case waitPayload:
			{
				buf, err := ReadLimitedLength(p._conn, p._payloadSize)
				if err != nil {
					return nil, 0, err
				}

				recvBuff = append(recvBuff, buf...)
				curState = waitCRC
			}
		case waitCRC:
			{
				buf, err := ReadLimitedLength(p._conn, TailCrcSize)
				if err != nil {
					return nil, 0, err
				}

				clientCrcCheckSum := util.BytesToUInt(buf)
				ServerCrcCheckSum := crc32.ChecksumIEEE(recvBuff)

				if clientCrcCheckSum != ServerCrcCheckSum {
					return nil, 0, fmt.Errorf("CRC CheskSum Error!")
				}

				return recvBuff, p._payloadSize, nil
			}
		}
	}
}

//Write :Write network packet with header and crc checksum
func (p *PacketHelper) Write(payload []byte) {

	writer := util.NewBinaryWriter()

	//WriteBytes has been done header parts. Doesn't need to call WriteInt for the header again.
	//writer.WriteInt(len(payload))
	writer.WriteBytes(payload)
	writer.WriteUInt(crc32.ChecksumIEEE(payload))

	packet := writer.ToBytes()
	p._conn.Write(packet)
}

//NewPacketHelper : Constructor for PacketHelper
func NewPacketHelper(conn net.Conn) *PacketHelper {
	reader := &PacketHelper{
		_conn:        conn,
		_payloadSize: 0,
	}
	return reader
}
