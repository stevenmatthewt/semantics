package bump

// Map represents a mapping of commit message descriptors
// to what they should increment. I.e. Major, Minor, Patch
type Map struct {
	Major string
	Minor string
	Patch string
}
