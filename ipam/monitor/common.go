package monitor

import (
	"github.com/weaveworks/weave/net/address"
)

// AllocMonitor is an interface for tracking changes in ring allocations.
type AllocMonitor interface {
	// HandleUpdate is called whenever an address ring gets updated.
	//
	// {old,new}Ranges correspond to address ranges owned by a peer which
	// executes the method.
	HandleUpdate(oldRanges, newRanges []address.Range)
}
