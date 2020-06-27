package syncgrp

import (
	"sync"
)

// SyncGrp wraps pointers to a sync.Mutex, sync.RWMutex
// and a sync.WaitGroup
type SyncGrp struct {
	mu *sync.Mutex
	rw *sync.RWMutex
	wg *sync.WaitGroup
}

// New returns a pointer to a SyncGrp
func New() *SyncGrp {
	var sg SyncGrp
	sg.mu = new(sync.Mutex)
	sg.rw = new(sync.RWMutex)
	sg.wg = new(sync.WaitGroup)
	return &sg
}

// Lock locks the SyncGrp's underlying Mutex
func (sg *SyncGrp) Lock() {
	sg.mu.Lock()
}

// LockCndtnl calls Lock iff the bool is true
func (sg *SyncGrp) LockCndtnl(doit bool) {
	if doit {
		sg.mu.Lock()
	}
}

// RWLock allows a SyncGrp to Lock based on its
// underlying mutex
func (sg *SyncGrp) RWLock() {
	sg.rw.Lock()
}

// RWLockCndtnl calls Lock iff the bool is true
func (sg *SyncGrp) RWLockCndtnl(doit bool) {
	if doit {
		sg.rw.Lock()
	}
}

// RLock allows a SyncGrp to Lock based on its
// underlying mutex
func (sg *SyncGrp) RLock() {
	sg.rw.RLock()
}

// Unlock allows a SyncGrp to Lock based on its
// underlying mutex
func (sg *SyncGrp) Unlock() {
	sg.mu.Unlock()
}

// RUnlock unlocks the SyncGrp's underlying RWMutex
func (sg *SyncGrp) RUnlock() {
	sg.rw.RUnlock()
}

// UnlockCndtnl calls Unlock iff the bool is true
func (sg *SyncGrp) UnlockCndtnl(doit bool) {
	if doit {
		sg.mu.Unlock()
	}
}

// RWUnlock allows a SyncGrp to Lock based on its
// underlying mutex
func (sg *SyncGrp) RWUnlock() {
	sg.rw.Unlock()
}

// RWUnlockCndtnl calls Unlock iff the bool is true
func (sg *SyncGrp) RWUnlockCndtnl(doit bool) {
	if doit {
		sg.rw.Unlock()
	}
}

// Add allows a SyncGrp to directly access its
// underlying Waitgroup's function
func (sg *SyncGrp) Add(delta int) {
	sg.wg.Add(delta)
}

// AddCndtnl calls Add iff the bool argument is true
func (sg *SyncGrp) AddCndtnl(delta int, doit bool) {
	if doit {
		sg.wg.Add(delta)
	}
}

// Done allows a SyncGrp to directly access its
// underlying Waitgroup's function
func (sg *SyncGrp) Done() {
	sg.wg.Done()
}

// DoneCndtnl calls Done iff the bool argument is true
func (sg *SyncGrp) DoneCndtnl(doit bool) {
	if doit {
		sg.wg.Done()
	}
}

// Wait allows a SyncGrp to directly access its
// underlying Waitgroup's function
func (sg *SyncGrp) Wait() {
	sg.wg.Wait()
}

// WaitCndtnl calls Wait iff the bool argument is true
func (sg *SyncGrp) WaitCndtnl(doit bool) {
	if doit {
		sg.wg.Wait()
	}
}
