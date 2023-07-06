package simple

import "errors"

type SimpleRepo struct {
	Error bool
}

func NewSimpleRepo(isError bool) *SimpleRepo {
	return &SimpleRepo{
		Error: isError,
	}
}

type SimpleService struct {
	*SimpleRepo
}

func NewSimpleService(repo *SimpleRepo) (*SimpleService, error) {
	if repo.Error {
		return nil, errors.New("Failed create service")
	} else {
		return &SimpleService{SimpleRepo: repo}, nil
	}
}
