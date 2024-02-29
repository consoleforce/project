package adapter

type ContactRepository interface {
	// Define your contact repository interface methods here
}

type contactRepository struct {
	// Define any dependencies or database connections here
}

func NewContactRepository() ContactRepository {
	// Implement your constructor function here
	return &contactRepository{}
}

// Implement the ContactRepository interface methods here
