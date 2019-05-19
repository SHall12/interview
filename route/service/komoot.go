package service

const KomootName string = "Komoot"

type komoot struct {
	*DefaultService
}

func NewKomoot() *komoot {
	return &komoot{
		&DefaultService{routes: []string{"SRT", "Welsh Mountain", "Oaks to Philly"}},
	}
}

func (svc *komoot) GetRoutesByUser(user string) ([]string, error) {
	routes := svc.routes
	userRoutes := make([]string, len(routes))
	for index, val := range routes {
		userRoutes[index] = user + val + user
	}
	return userRoutes, nil
}
