package client

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	network "../lib/network"
	util "../lib/utility"
	proto "../proto"
)

type Player struct {
	_conn   *net.TCPConn
	_helper *network.PacketHelper
}

func (p *Player) Init_tcp() {
	laddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:123")
	if err != nil {
		log.Fatal(err)
	}
	tcp, err := net.DialTCP("tcp", nil, laddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("[Client Info] Server connected.")

	p._conn = tcp
	p._helper = network.NewPacketHelper(p._conn)
}

func (p *Player) Start() {

	p.Init_tcp()

	defer p._conn.Close()

	go p.handleConnection()

	for {
		time.Sleep(16 * time.Millisecond)
	}
}

//handleConnection
func (p *Player) handleConnection() {

	for {
		time.Sleep(1000 * time.Millisecond)

		count := util.RandInteger(1, 2000)
		bs := make([]proto.IProtocol, count)
		for i := 0; i < count; i++ {
			pro := new(proto.HeartBeatProto)
			bs[i] = pro
		}

		sendData := proto.TransferProtocolToRawData(bs)
		p._helper.Write(sendData)
		fmt.Println("[Client Info] Sent a packet size = " + strconv.Itoa(len(sendData)))

		recvData, size, err := p._helper.Read()

		if err != nil {
			p._conn.Close()
			fmt.Println("[Client Info] Read error:", err)
			break
		} else {

			fmt.Println("[Client Info] Recieved a packet size = " + strconv.Itoa(size))

			protos := proto.TransferRawDataToProtocol(recvData)
			for _, proto := range protos {
				proto.Todo()
			}
		}
	}
}

func NewPlayer() *Player {
	player := &Player{}
	return player
}
