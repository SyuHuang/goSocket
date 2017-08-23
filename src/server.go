package main

import (
	"fmt"
	"net"
	"strconv"

	network "./lib/network"
	util "./lib/utility"
	proto "./proto"
)

//ServerStart :
func ServerStart(ip string, port int) {
	go StartListen(ip, port)
}

//StartListen :
func StartListen(ip string, port int) {

	setting := fmt.Sprintf("%s:%d", ip, port)
	server, err := net.Listen("tcp", setting)
	CheckError(err)

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println("[Server Info] Accept error:", err)
			continue
		}

		fmt.Println("[Server Info] A new connection is established from " + conn.RemoteAddr().String() + ".")

		go handleConnection(conn)
	}
}

//handleConnection
func handleConnection(conn net.Conn) {

	helper := network.NewPacketHelper(conn)

	for {
		recvData, size, err := helper.Read()

		if err != nil {

			conn.Close()
			fmt.Println("[Server Info] Read error:", err)
			break

		} else {

			fmt.Println("[Server Info] Recieved a packet size = " + strconv.Itoa(size))

			protos := proto.TransferRawDataToProtocol(recvData)
			for _, proto := range protos {
				//Implement protocol's game logic
				proto.Todo()
			}

			count := util.RandInteger(1, 2000)
			bs := make([]proto.IProtocol, count)
			for i := 0; i < count; i++ {
				pro := new(proto.HeartBeatProto)
				bs[i] = pro
			}

			sendData := proto.TransferProtocolToRawData(bs)
			helper.Write(sendData)

			fmt.Println("[Server Info] Sent a packet size = " + strconv.Itoa(len(sendData)))
		}
	}

	defer conn.Close()
}
