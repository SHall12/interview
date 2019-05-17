package service

// TODO: Make other services
// TODO: make tests

const StravaName string = "Strava"

type strava struct {
	*DefaultService
}

func NewStrava() *strava {
	return &strava{
		&DefaultService{routes: []string{"SRT", "CVT", "Perkiomen"}},
	}
}

func (s *strava) getRoutesByUser(user string) ([]string, error) {
	routes := s.routes
	userRoutes := make([]string, len(routes))
	for index, val := range routes {
		userRoutes[index] = user + val
	}
	return userRoutes, nil
}
