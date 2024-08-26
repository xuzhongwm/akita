package tlb

import (
	"github.com/sarchlab/akita/v4/mem/vm"
	"github.com/sarchlab/akita/v4/sim"
)

// A FlushReq asks the TLB to invalidate certain entries. It will also not block all incoming and outgoing ports
type FlushReq struct {
	sim.MsgMeta
	VAddr []uint64
	PID   vm.PID
}

// Meta returns the meta data associated with the message.
func (r *FlushReq) Meta() *sim.MsgMeta {
	return &r.MsgMeta
}

// Clone returns cloned FlushReq with different ID
func (r *FlushReq) Clone() sim.Msg {
	cloneMsg := *r
	cloneMsg.ID = sim.GetIDGenerator().Generate()

	return &cloneMsg
}

// FlushReqBuilder can build AT flush requests
type FlushReqBuilder struct {
	src, dst sim.Port
	vAddrs   []uint64
	pid      vm.PID
}

// WithSrc sets the source of the request to build.
func (b FlushReqBuilder) WithSrc(src sim.Port) FlushReqBuilder {
	b.src = src
	return b
}

// WithDst sets the destination of the request to build.
func (b FlushReqBuilder) WithDst(dst sim.Port) FlushReqBuilder {
	b.dst = dst
	return b
}

// WithVAddrs sets the Vaddr of the pages to be flushed
func (b FlushReqBuilder) WithVAddrs(vAddrs []uint64) FlushReqBuilder {
	b.vAddrs = vAddrs
	return b
}

// WithPID sets the pid whose entries are to be flushed
func (b FlushReqBuilder) WithPID(pid vm.PID) FlushReqBuilder {
	b.pid = pid
	return b
}

// Build creates a new TLBFlushReq
func (b FlushReqBuilder) Build() *FlushReq {
	r := &FlushReq{}
	r.ID = sim.GetIDGenerator().Generate()
	r.Src = b.src
	r.Dst = b.dst
	r.VAddr = b.vAddrs
	r.PID = b.pid
	return r
}

// A FlushRsp is a response from AT indicating flush is complete
type FlushRsp struct {
	sim.MsgMeta
}

// Meta returns the meta data associated with the message.
func (r *FlushRsp) Meta() *sim.MsgMeta {
	return &r.MsgMeta
}

// Clone returns cloned FlushRsp with different ID
func (r *FlushRsp) Clone() sim.Msg {
	cloneMsg := *r
	cloneMsg.ID = sim.GetIDGenerator().Generate()

	return &cloneMsg
}

// FlushRspBuilder can build AT flush rsp
type FlushRspBuilder struct {
	src, dst sim.Port
}

// WithSrc sets the source of the request to build.
func (b FlushRspBuilder) WithSrc(src sim.Port) FlushRspBuilder {
	b.src = src
	return b
}

// WithDst sets the destination of the request to build.
func (b FlushRspBuilder) WithDst(dst sim.Port) FlushRspBuilder {
	b.dst = dst
	return b
}

// Build creates a new TLBFlushRsps.
func (b FlushRspBuilder) Build() *FlushRsp {
	r := &FlushRsp{}
	r.ID = sim.GetIDGenerator().Generate()
	r.Src = b.src
	r.Dst = b.dst

	return r
}

// A RestartReq is a request to TLB to start accepting requests and resume operations
type RestartReq struct {
	sim.MsgMeta
}

// Meta returns the meta data associated with the message.
func (r *RestartReq) Meta() *sim.MsgMeta {
	return &r.MsgMeta
}

// Clone returns cloned RestartReq with different ID
func (r *RestartReq) Clone() sim.Msg {
	cloneMsg := *r
	cloneMsg.ID = sim.GetIDGenerator().Generate()

	return &cloneMsg
}

// RestartReqBuilder can build TLB restart requests.
type RestartReqBuilder struct {
	src, dst sim.Port
}

// WithSrc sets the source of the request to build.
func (b RestartReqBuilder) WithSrc(src sim.Port) RestartReqBuilder {
	b.src = src
	return b
}

// WithDst sets the destination of the request to build.
func (b RestartReqBuilder) WithDst(dst sim.Port) RestartReqBuilder {
	b.dst = dst
	return b
}

// Build creates a new TLBRestartReq.
func (b RestartReqBuilder) Build() *RestartReq {
	r := &RestartReq{}
	r.ID = sim.GetIDGenerator().Generate()
	r.Src = b.src
	r.Dst = b.dst

	return r
}

// A RestartRsp is a response from AT indicating it has resumed working
type RestartRsp struct {
	sim.MsgMeta
}

// Meta returns the meta data associated with the message.
func (r *RestartRsp) Meta() *sim.MsgMeta {
	return &r.MsgMeta
}

// Clone returns cloned RestartRsp with different ID
func (r *RestartRsp) Clone() sim.Msg {
	cloneMsg := *r
	cloneMsg.ID = sim.GetIDGenerator().Generate()

	return &cloneMsg
}

// RestartRspBuilder can build AT flush rsp
type RestartRspBuilder struct {
	src, dst sim.Port
}

// WithSrc sets the source of the request to build.
func (b RestartRspBuilder) WithSrc(src sim.Port) RestartRspBuilder {
	b.src = src
	return b
}

// WithDst sets the destination of the request to build.
func (b RestartRspBuilder) WithDst(dst sim.Port) RestartRspBuilder {
	b.dst = dst
	return b
}

// Build creates a new TLBRestartRsp
func (b RestartRspBuilder) Build() *RestartRsp {
	r := &RestartRsp{}
	r.ID = sim.GetIDGenerator().Generate()
	r.Src = b.src
	r.Dst = b.dst

	return r
}
