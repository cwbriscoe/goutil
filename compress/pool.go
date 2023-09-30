package compress

// Pool is the Compress interface for compression pools of different types.
type Pool interface {
	Compress([]byte) ([]byte, error)
}
