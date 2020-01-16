// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SynchronizationStatusCode undocumented
type SynchronizationStatusCode int

const (
	// SynchronizationStatusCodeVNotConfigured undocumented
	SynchronizationStatusCodeVNotConfigured SynchronizationStatusCode = 0
	// SynchronizationStatusCodeVNotRun undocumented
	SynchronizationStatusCodeVNotRun SynchronizationStatusCode = 1
	// SynchronizationStatusCodeVActive undocumented
	SynchronizationStatusCodeVActive SynchronizationStatusCode = 2
	// SynchronizationStatusCodeVPaused undocumented
	SynchronizationStatusCodeVPaused SynchronizationStatusCode = 3
	// SynchronizationStatusCodeVQuarantine undocumented
	SynchronizationStatusCodeVQuarantine SynchronizationStatusCode = 4
)

// SynchronizationStatusCodePNotConfigured returns a pointer to SynchronizationStatusCodeVNotConfigured
func SynchronizationStatusCodePNotConfigured() *SynchronizationStatusCode {
	v := SynchronizationStatusCodeVNotConfigured
	return &v
}

// SynchronizationStatusCodePNotRun returns a pointer to SynchronizationStatusCodeVNotRun
func SynchronizationStatusCodePNotRun() *SynchronizationStatusCode {
	v := SynchronizationStatusCodeVNotRun
	return &v
}

// SynchronizationStatusCodePActive returns a pointer to SynchronizationStatusCodeVActive
func SynchronizationStatusCodePActive() *SynchronizationStatusCode {
	v := SynchronizationStatusCodeVActive
	return &v
}

// SynchronizationStatusCodePPaused returns a pointer to SynchronizationStatusCodeVPaused
func SynchronizationStatusCodePPaused() *SynchronizationStatusCode {
	v := SynchronizationStatusCodeVPaused
	return &v
}

// SynchronizationStatusCodePQuarantine returns a pointer to SynchronizationStatusCodeVQuarantine
func SynchronizationStatusCodePQuarantine() *SynchronizationStatusCode {
	v := SynchronizationStatusCodeVQuarantine
	return &v
}