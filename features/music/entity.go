package music

type Core struct {

}
// interface untuk Data Layer
type UserDataInterface interface {
	Insert(input Core) error
	SelectAll() ([]Core, error)
}

// interface untuk Service Layer
type UserServiceInterface interface {
	Create(input Core) error
	SelectAll() ([]Core, error)
}
