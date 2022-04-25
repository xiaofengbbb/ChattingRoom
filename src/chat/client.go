package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn       //连接
	nick     string         //名称
	room     *room          //房间
	commands chan<- command //命令
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		//去掉字符串msg中首部以及尾部与字符串"\r\n"中每个相匹配的字符
		msg = strings.Trim(msg, "\r\n")

		//按照" ",分割成两个字符串，组成一个字符串切片
		args := strings.Split(msg, " ")

		//取出字符串切片中第一个字符串
		cmd := strings.Trim(args[0], " ")

		switch cmd {
		case "/nick":
			c.commands <- command{
				id:     CMD_NICK,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CMD_ROOMS,
				client: c,
				args:   args,
			}
		case "/msg":
			c.commands <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		case "/quit":
			c.commands <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
			}
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("ERROR: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	//Write 将数据写入连接。
	c.conn.Write([]byte("> " + msg + "\n"))
}
