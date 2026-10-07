package grpctunnel

import "testing"

func TestWindowSize(t *testing.T) {
	var unset *tunnelOpts
	if got := unset.windowSize(); got != initialWindowSize {
		t.Errorf("nil options: %d, want the default %d", got, initialWindowSize)
	}

	var opts tunnelOpts
	if got := opts.windowSize(); got != initialWindowSize {
		t.Errorf("unset: %d, want the default %d", got, initialWindowSize)
	}

	WithInitialWindowSize(0).apply(&opts)
	if got := opts.windowSize(); got != initialWindowSize {
		t.Errorf("zero: %d, want the default %d", got, initialWindowSize)
	}

	WithInitialWindowSize(4 << 20).apply(&opts)
	if got := opts.windowSize(); got != 4<<20 {
		t.Errorf("set: %d, want %d", got, 4<<20)
	}

	handler := NewTunnelServiceHandler(TunnelServiceHandlerOptions{InitialWindowSize: 1 << 20})
	if got := handler.tunnelOpts.windowSize(); got != 1<<20 {
		t.Errorf("handler: %d, want %d", got, 1<<20)
	}
}
