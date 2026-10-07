package grpctunnel

import "github.com/jhump/grpctunnel/tunnelpb"

// TunnelOption is an option for configuring the behavior of
// a tunnel client or tunnel server.
type TunnelOption interface {
	apply(*tunnelOpts)
}

// WithDisableFlowControl returns an option that disables the
// use of flow control, even when the tunnel peer supports it.
//
// NOTE: This should NOT be used in application code. This is
// intended for test code, to verify that the tunnels work
// without flow control, to make sure they can interop correctly
// with older versions of this package, before flow control was
// introduced.
//
// Eventually, older versions that do not use flow control will
// not be supported and this option will be removed.
func WithDisableFlowControl() TunnelOption {
	return tunnelOptFunc(func(opts *tunnelOpts) {
		opts.disableFlowControl = true
	})
}

// WithInitialWindowSize returns an option that sets the flow control window
// this side of the tunnel announces for each stream: how much the peer may
// send on a stream before this side has read it. A stream moves at most one
// window per round trip of the tunnel, so a larger window lets a stream go
// faster over a tunnel with a longer round trip, at the cost of buffering up to
// that much per stream. Zero keeps the default of 64 KiB.
//
// Each side announces its own window and the peer honours it, so the two
// sides of a tunnel need not agree, and a peer that does not set this keeps
// the default.
func WithInitialWindowSize(size uint32) TunnelOption {
	return tunnelOptFunc(func(opts *tunnelOpts) {
		opts.initialWindowSize = size
	})
}

type tunnelOpts struct {
	disableFlowControl bool
	initialWindowSize  uint32
}

// windowSize is the flow control window this side announces for each stream.
func (t *tunnelOpts) windowSize() uint32 {
	if t == nil || t.initialWindowSize == 0 {
		return initialWindowSize
	}
	return t.initialWindowSize
}

func (t *tunnelOpts) supportedRevisions() []tunnelpb.ProtocolRevision {
	if t.disableFlowControl {
		return []tunnelpb.ProtocolRevision{tunnelpb.ProtocolRevision_REVISION_ZERO}
	}
	return []tunnelpb.ProtocolRevision{
		tunnelpb.ProtocolRevision_REVISION_ZERO, tunnelpb.ProtocolRevision_REVISION_ONE,
	}
}

type tunnelOptFunc func(*tunnelOpts)

func (t tunnelOptFunc) apply(opts *tunnelOpts) {
	t(opts)
}
