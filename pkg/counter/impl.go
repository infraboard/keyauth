package counter

// Service todo
type Service interface {
	GetNextSequenceValue(sequenceName string) (*Count, error)
}

// NewCount todo
func NewCount() *Count {
	return &Count{}
}

// Count todo
type Count struct {
	SequenceName string `bson:"_id"`
	Value        uint64 `bson:"value"`
}
