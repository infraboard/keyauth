//go:generate  mcube enum -m

package user

const (
	// Unknown (unknown) todo
	Unknown Gender = iota
	// Male (male) 男
	Male
	// Female (female) 女
	Female
)

// Gender 性别
type Gender uint
