package Test

type Service interface {
	GetUserTest() (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetUserTest() (User, error){
	user, err := s.repository.FindById()
	if err != nil{
		return user, err
	}
	return user, err
}