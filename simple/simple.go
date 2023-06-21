package simple

type SimpleRepo struct {
}

func NewSimpleRepo() *SimpleRepo {
	return &SimpleRepo{}
}

type SimpleService struct {
	*SimpleRepo
}

func NewSimpleService(repo *SimpleRepo) *SimpleService {
	return &SimpleService{SimpleRepo: repo}
}
