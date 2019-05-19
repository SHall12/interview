package route

// HandlerMock allows tests to control struct behavior
type HandlerMock struct {
	GetRoutesMock          func() ([]string, error)
	GetRoutesByUserMock    func() ([]string, error)
	GetAllRoutesMock       func() ([]string, error)
	GetAllRoutesByUserMock func() ([]string, error)
	GetUniqueRoutesMock    func() ([]string, error)
	GetServicesMock        func() []string
}

func (m HandlerMock) GetRoutes(serviceNames []string) ([]string, error) {
	return m.GetRoutesMock()
}

func (m HandlerMock) GetRoutesByUser(serviceNames []string, user string) ([]string, error) {
	return m.GetRoutesByUserMock()
}

func (m HandlerMock) GetAllRoutes() ([]string, error) {
	return m.GetAllRoutesMock()
}

func (m HandlerMock) GetAllRoutesByUser(user string) ([]string, error) {
	return m.GetAllRoutesByUserMock()
}

func (m HandlerMock) GetUniqueRoutes(serviceNames []string) ([]string, error) {
	return m.GetUniqueRoutesMock()
}

func (m HandlerMock) GetServices() []string {
	return m.GetServicesMock()
}
