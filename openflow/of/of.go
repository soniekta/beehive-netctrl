// Automatically generated by Packet Go code generator.
package of

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/packet/packet/src/go/packet"
)

type Versions int

const (
	OPENFLOW_1_0 Versions = 1
	OPENFLOW_1_1 Versions = 2
	OPENFLOW_1_2 Versions = 3
	OPENFLOW_1_3 Versions = 4
)

type Constants int

const (
	P_ETH_ALEN           Constants = 6
	P_MAX_PORT_NAME_LEN  Constants = 16
	P_MAX_TABLE_NAME_LEN Constants = 32
)

type Type int

const (
	PT_HELLO              Type = 0
	PT_ERROR              Type = 1
	PT_ECHO_REQUEST       Type = 2
	PT_ECHO_REPLY         Type = 3
	PT_VENDOR             Type = 4
	PT_FEATURES_REQUEST   Type = 5
	PT_FEATURES_REPLY     Type = 6
	PT_GET_CONFIG_REQUEST Type = 7
	PT_GET_CONFIG_REPLY   Type = 8
	PT_SET_CONFIG         Type = 9
)

func NewHeaderWithBuf(b []byte) Header {
	return Header{packet.Packet{Buf: b}}
}

func NewHeader() Header {
	s := 8
	b := make([]byte, s)
	p := Header{packet.Packet{Buf: b}}
	p.Init()
	return p
}

type Header struct {
	packet.Packet
}

func (this Header) minSize() int {
	return 8
}

func (this Header) Clone() (Header, error) {
	var newBuf bytes.Buffer
	_, err := io.CopyN(&newBuf, bytes.NewBuffer(this.Buf), int64(this.Size()))
	if err != nil {
		return NewHeader(), err
	}

	return NewHeaderWithBuf(newBuf.Bytes()), nil
}

type HeaderConn struct {
	net.Conn
	w      *bufio.Writer
	buf    []byte
	offset int
}

func NewHeaderConn(c net.Conn) HeaderConn {
	return HeaderConn{
		Conn: c,
		w:    bufio.NewWriter(c),
		buf:  make([]byte, packet.DefaultBufSize),
	}
}

func (c *HeaderConn) WriteHeader(pkt Header) error {
	s := pkt.Size()
	b := pkt.Buffer()[:s]
	n := 0
	for s > 0 {
		var err error
		if n, err = c.w.Write(b); err != nil {
			return fmt.Errorf("Error in write: %v", err)
		}
		s -= n
	}

	return nil
}

func (c *HeaderConn) WriteHeaders(pkts []Header) error {
	for _, p := range pkts {
		if err := c.WriteHeader(p); err != nil {
			return err
		}
	}
	return nil
}

func (c *HeaderConn) Flush() error {
	return c.w.Flush()
}

func (c *HeaderConn) ReadHeader() (Header, error) {
	pkts := make([]Header, 1)
	_, err := c.ReadHeaders(pkts)
	if err != nil {
		return NewHeader(), err
	}

	return pkts[0], nil
}

func (c *HeaderConn) ReadHeaders(pkts []Header) (int, error) {
	if len(c.buf) == c.offset {
		newSize := packet.DefaultBufSize
		if newSize < len(c.buf) {
			newSize = 2 * len(c.buf)
		}

		buf := make([]byte, newSize)
		copy(buf, c.buf[:c.offset])
		c.buf = buf
	}

	r, err := c.Conn.Read(c.buf[c.offset:])
	if err != nil {
		return 0, err
	}

	r += c.offset

	s := 0
	n := 0
	for i := range pkts {
		p := NewHeaderWithBuf(c.buf[s:])

		pSize := p.Size()
		if pSize == 0 || r < s+pSize {
			break
		}

		pkts[i] = p
		s += pSize
		n++
	}

	c.offset = r - s
	if c.offset < 0 {
		panic("Invalid value for offset")
	}

	c.buf = c.buf[s:]
	return n, nil
}

func (this *Header) Init() {
	this.SetLength(uint16(this.minSize()))
	// Invariants.
}

func (this Header) Size() int {
	if len(this.Buf) < this.minSize() {
		return 0
	}

	size := int(this.Length())
	return size
}

func ToHeader(p packet.Packet) (Header, error) {
	if !IsHeader(p) {
		return NewHeaderWithBuf(nil), errors.New("Cannot convert to of.Header")
	}

	return NewHeaderWithBuf(p.Buf), nil
}

func IsHeader(p packet.Packet) bool {
	return true
}

func (this *Header) Version() uint8 {
	offset := this.VersionOffset()
	res := uint8(this.Buf[offset])
	return res
}

func (this *Header) SetVersion(v uint8) {
	offset := this.VersionOffset()
	this.Buf[offset] = byte(v)
	offset++
}

func (this *Header) VersionOffset() int {
	offset := 0
	return offset
}

func (this *Header) Type() uint8 {
	offset := this.TypeOffset()
	res := uint8(this.Buf[offset])
	return res
}

func (this *Header) SetType(t uint8) {
	offset := this.TypeOffset()
	this.Buf[offset] = byte(t)
	offset++
}

func (this *Header) TypeOffset() int {
	offset := 1
	return offset
}

func (this *Header) Length() uint16 {
	offset := this.LengthOffset()
	res := binary.BigEndian.Uint16(this.Buf[offset:])
	return res
}

func (this *Header) SetLength(l uint16) {
	offset := this.LengthOffset()
	binary.BigEndian.PutUint16(this.Buf[offset:], l)
	offset += 2
}

func (this *Header) LengthOffset() int {
	offset := 2
	return offset
}

func (this *Header) Xid() uint32 {
	offset := this.XidOffset()
	res := binary.BigEndian.Uint32(this.Buf[offset:])
	return res
}

func (this *Header) SetXid(x uint32) {
	offset := this.XidOffset()
	binary.BigEndian.PutUint32(this.Buf[offset:], x)
	offset += 4
}

func (this *Header) XidOffset() int {
	offset := 4
	return offset
}

func NewHelloWithBuf(b []byte) Hello {
	return Hello{Header{packet.Packet{Buf: b}}}
}

func NewHello() Hello {
	s := 8
	b := make([]byte, s)
	p := Hello{Header{packet.Packet{Buf: b}}}
	p.Init()
	return p
}

type Hello struct {
	Header
}

func (this Hello) minSize() int {
	return 8
}

func (this Hello) Clone() (Hello, error) {
	var newBuf bytes.Buffer
	_, err := io.CopyN(&newBuf, bytes.NewBuffer(this.Buf), int64(this.Size()))
	if err != nil {
		return NewHello(), err
	}

	return NewHelloWithBuf(newBuf.Bytes()), nil
}

type HelloConn struct {
	net.Conn
	w      *bufio.Writer
	buf    []byte
	offset int
}

func NewHelloConn(c net.Conn) HelloConn {
	return HelloConn{
		Conn: c,
		w:    bufio.NewWriter(c),
		buf:  make([]byte, packet.DefaultBufSize),
	}
}

func (c *HelloConn) WriteHello(pkt Hello) error {
	s := pkt.Size()
	b := pkt.Buffer()[:s]
	n := 0
	for s > 0 {
		var err error
		if n, err = c.w.Write(b); err != nil {
			return fmt.Errorf("Error in write: %v", err)
		}
		s -= n
	}

	return nil
}

func (c *HelloConn) WriteHellos(pkts []Hello) error {
	for _, p := range pkts {
		if err := c.WriteHello(p); err != nil {
			return err
		}
	}
	return nil
}

func (c *HelloConn) Flush() error {
	return c.w.Flush()
}

func (c *HelloConn) ReadHello() (Hello, error) {
	pkts := make([]Hello, 1)
	_, err := c.ReadHellos(pkts)
	if err != nil {
		return NewHello(), err
	}

	return pkts[0], nil
}

func (c *HelloConn) ReadHellos(pkts []Hello) (int, error) {
	if len(c.buf) == c.offset {
		newSize := packet.DefaultBufSize
		if newSize < len(c.buf) {
			newSize = 2 * len(c.buf)
		}

		buf := make([]byte, newSize)
		copy(buf, c.buf[:c.offset])
		c.buf = buf
	}

	r, err := c.Conn.Read(c.buf[c.offset:])
	if err != nil {
		return 0, err
	}

	r += c.offset

	s := 0
	n := 0
	for i := range pkts {
		p := NewHelloWithBuf(c.buf[s:])

		pSize := p.Size()
		if pSize == 0 || r < s+pSize {
			break
		}

		pkts[i] = p
		s += pSize
		n++
	}

	c.offset = r - s
	if c.offset < 0 {
		panic("Invalid value for offset")
	}

	c.buf = c.buf[s:]
	return n, nil
}

func (this *Hello) Init() {
	this.Header.Init()
	this.SetLength(uint16(this.minSize()))
	// Invariants.
	this.SetType(uint8(0)) // type
}

func (this Hello) Size() int {
	if len(this.Buf) < this.minSize() {
		return 0
	}

	size := int(this.Length())
	return size
}

func ToHello(p Header) (Hello, error) {
	if !IsHello(p) {
		return NewHelloWithBuf(nil), errors.New("Cannot convert to of.Hello")
	}

	return NewHelloWithBuf(p.Buf), nil
}

func IsHello(p Header) bool {
	return p.Type() == 0 && true
}
