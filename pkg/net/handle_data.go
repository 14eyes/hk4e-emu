package net

import (
	"encoding/binary"
	"log"
	"net"
)

func (l *KCPConn) handleDataPacket(addr *net.UDPAddr, buf []byte) error {
	id := binary.LittleEndian.Uint32(buf[:4])
	token := binary.LittleEndian.Uint32(buf[4:8])
	session, err := l.getSession(addr, id, token)
	if err != nil {
		log.Println("[net.KCPConn] Failed to get session, error:", err)
		return l.sendCtrlDisconnect(addr, id, token, 5) // ENET_SERVER_KICK
	}
	return session.OnPacket(buf, l.packetCh)
}

func (l *KCPConn) sendDataPacket(addr *net.UDPAddr, buf []byte) (err error) {
	_, err = l.conn.WriteToUDP(buf, addr)
	return err
}
