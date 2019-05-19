package service

const StravaName string = "Strava"

type strava struct {
	*DefaultService
}

func NewStrava() *strava {
	return &strava{
		&DefaultService{routes: []string{"SRT", "CVT", "Perkiomen"}},
	}
}

func (svc *strava) GetRoutesByUser(user string) ([]string, error) {
	routes := svc.routes
	userRoutes := make([]string, len(routes))
	for index, val := range routes {
		userRoutes[index] = user + val
	}
	return userRoutes, nil
}
