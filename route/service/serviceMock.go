package service

type MockService struct {
	GetRoutesMock       func() ([]string, error)
	GetRoutesByUserMock func() ([]string, error)
}

func (m MockService) GetRoutes() ([]string, error)                  { return m.GetRoutesMock() }
func (m MockService) GetRoutesByUser(user string) ([]string, error) { return m.GetRoutesByUserMock() }
