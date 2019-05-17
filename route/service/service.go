package service

type Service interface {
	getRoutes() ([]string, error)
	getRoutesByUser(user string) ([]string, error)
}

type DefaultService struct {
	routes []string
}

func (s *DefaultService) getRoutes() ([]string, error) {
	return s.routes, nil
}

func (s *DefaultService) getRoutesByUser(user string) ([]string, error) {
	return s.routes, nil
}
