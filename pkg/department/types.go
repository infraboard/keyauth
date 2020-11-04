//go:generate  mcube enum -m

package department

const (
	// Pending (pending) todo
	Pending ApplicationFormStatus = iota
	// Passed (passed)
	Passed
	// Deny (deny)
	Deny
)

// ApplicationFormStatus 性别
type ApplicationFormStatus uint
