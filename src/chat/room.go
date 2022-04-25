package main

import "net"

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {
	for addr, m := range r.members {
		// remoteAddress()用于绑定地址和端口的，用于客户端
		if addr != sender.conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}
