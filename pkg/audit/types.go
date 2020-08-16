//go:generate  mcube enum -m

package audit

// Result todo
type Result uint

const (
	// Success (success) todo
	Success Result = iota
	// Failed (failed) todo
	Failed
)
