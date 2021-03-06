package bootstrap

import (
	"github.com/halia-group/halia/channel"
	"net"
)

type ServerOptions struct {
	ChannelFactory func(conn net.Conn) channel.Channel
}

type Server struct {
	listener net.Listener
	options  *ServerOptions
}

func NewServer(options *ServerOptions) *Server {
	return &Server{options: options}
}

func (server *Server) Listen(network, addr string) error {
	var err error
	server.listener, err = net.Listen(network, addr)
	if err != nil {
		return err
	}
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			return err
		}
		go server.onConnect(conn)
	}
}

func (server *Server) onConnect(conn net.Conn) {
	c := server.options.ChannelFactory(conn)
	defer server.onDisconnect(c)

	c.Pipeline().FireChannelActive()
	// 数据包读取由入站handler进行轮询读取
	c.Pipeline().FireChannelRead(nil)
}

// 断开连接
func (server *Server) onDisconnect(c channel.Channel) {
	c.Pipeline().FireChannelInActive()
}
