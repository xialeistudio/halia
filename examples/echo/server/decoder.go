package main

import (
	"bufio"
	"halia/channel"
	"io"
)

type LineDelimiterFrameDecoder struct{}

func (d *LineDelimiterFrameDecoder) OnError(c channel.HandlerContext, err error) {
	c.FireOnError(err)
}

func (d *LineDelimiterFrameDecoder) ChannelActive(c channel.HandlerContext) {
	c.FireChannelActive()
}

func (d *LineDelimiterFrameDecoder) ChannelInActive(c channel.HandlerContext) {
	c.FireChannelInActive()
}

func (d *LineDelimiterFrameDecoder) ChannelRead(c channel.HandlerContext, _ interface{}) {
	var br = bufio.NewReader(c.Channel())
	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			return
		}
		if err != nil {
			d.OnError(c, err)
			return
		}
		c.FireChannelRead(string(line))
	}
}
