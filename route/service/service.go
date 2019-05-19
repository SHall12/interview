package service

type Service interface {
	GetRoutes() ([]string, error)
	GetRoutesByUser(user string) ([]string, error)
}

type DefaultService struct {
	routes []string
}

func (s *DefaultService) GetRoutes() ([]string, error) {
	return s.routes, nil
}

func (s *DefaultService) GetRoutesByUser(user string) ([]string, error) {
	return s.routes, nil
}
