package openflow

import (
	"fmt"

	"github.com/soheilhy/beehive-netctrl/nom"
	"github.com/soheilhy/beehive-netctrl/openflow/of10"
	"github.com/soheilhy/beehive-netctrl/openflow/of12"
)

func (of *of10Driver) handlePacketIn(in of10.PacketIn, c *ofConn) error {
	port, ok := of.ofPorts[in.InPort()]
	if !ok {
		return fmt.Errorf("of10driver: port not found %v", in.InPort())
	}

	nomIn := nom.PacketIn{
		Node:     c.node.UID(),
		InPort:   port.UID(),
		BufferID: nom.PacketBufferID(in.BufferId()),
	}
	nomIn.Packet = nom.Packet(in.Data())
	c.ctx.Emit(nomIn)

	//c.ctx.Emit(in)

	//buf := make([]byte, 32)
	//out := of10.NewPacketOutWithBuf(buf)
	//out.Init()
	//out.SetBufferId(in.BufferId())
	//out.SetInPort(in.InPort())

	//bcast := of10.NewActionOutput()
	//bcast.SetPort(uint16(of10.PP_FLOOD))

	//out.AddActions(bcast.ActionHeader)

	//if in.BufferId() == 0xFFFFFFFF {
	//for _, d := range in.Data() {
	//out.AddData(d)
	//}
	//} else {
	//out.SetBufferId(in.BufferId())
	//}

	//c.wCh <- out.Header
	//if err := c.WriteHeader(out.Header); err != nil {
	//return fmt.Errorf("Error in writing a packet out: %v", err)
	//}
	return nil
}

func (of *of12Driver) handlePacketIn(in of12.PacketIn, c *ofConn) error {
	return nil
}
