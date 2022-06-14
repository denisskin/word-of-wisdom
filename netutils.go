package wow

import (
	"encoding/binary"
	"io"
	"net"
	"time"
)

const tcpReadWriteTimeout = 15 * time.Second

func writeBytes(conn net.Conn, data []byte) error {
	if err := conn.SetWriteDeadline(time.Now().Add(tcpReadWriteTimeout)); err != nil {
		return err
	}
	buf := make([]byte, 2+len(data))
	binary.BigEndian.PutUint16(buf, uint16(len(data)))
	copy(buf[2:], data)
	_, err := conn.Write(buf)
	return err
}

func readBytes(conn net.Conn) (data []byte, err error) {
	if err = conn.SetReadDeadline(time.Now().Add(tcpReadWriteTimeout)); err != nil {
		return
	}
	var buf [2]byte
	if _, err = io.ReadFull(conn, buf[:]); err != nil {
		return
	}
	n := binary.BigEndian.Uint16(buf[:])
	data = make([]byte, n)
	_, err = io.ReadFull(conn, data)
	return
}
