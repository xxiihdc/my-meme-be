package repository

type Repo[T interface{}] interface {
	GetAll() ([]T, error)
	GetByID(id int) (*T, error)
	Create(t T) (*T, error)
	Update(t T) (*T, error)
	Delete(id int) error
}
