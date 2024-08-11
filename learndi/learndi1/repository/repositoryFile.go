package repository

type IRepository interface {
	GetUser() string
}

type Repository struct {

}

func New() IRepository {
	return &Repository{}
}

func (r *Repository) GetUser() string {
	return "value from repo"
}