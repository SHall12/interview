package service

const RwgpsName string = "RWGPS"

type rwgps struct {
	*DefaultService
}

func NewRwgps() *rwgps {
	return &rwgps{
		&DefaultService{routes: []string{"CVT", "Perkiomen", "Welsh Mountain"}},
	}
}

func (svc *rwgps) GetRoutesByUser(user string) ([]string, error) {
	routes := svc.routes
	userRoutes := make([]string, len(routes))
	for index, val := range routes {
		userRoutes[index] = val + user
	}
	return userRoutes, nil
}
