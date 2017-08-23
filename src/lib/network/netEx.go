package network

import (
	"fmt"
	"io"
	"net"
)

//maximunBufferSize :Maximum packet size for the game.
const maximunBufferSize = 65535

//ReadLimitedLength :Read network data from conn util required size
func ReadLimitedLength(conn net.Conn, length int) ([]byte, error) {

	if length > maximunBufferSize {
		//return nil, 0, errors.New(fmt.Sprintf("Require size(%d) too long.", length))
		return nil, fmt.Errorf("Require size(%d) too long.", length)
	}

	buffer := []byte{}
	frag := make([]byte, length)
	totalSize := 0
	for {
		size, err := conn.Read(frag)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			return nil, err
		}

		//fmt.Println(size)

		buffer = append(buffer, frag[:size]...)

		totalSize += size
		if totalSize == length {
			return buffer[:totalSize], nil
		}

		frag = make([]byte, length-totalSize)
	}
}
