package data

type Repository interface {
}

type Service struct {
	brokerRepo Repository
}

// NewService create new service instance.
func NewService(brokerRepo Repository) *Service {
	return &Service{brokerRepo: brokerRepo}
}
