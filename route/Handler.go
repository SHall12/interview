package route

import (
	"fmt"
	"interview/route/service"
)

// Handler combines multiple route services
type Handler struct {
	services map[string]service.Service
}

// NewHandler instantiates route handler for defined services
func NewHandler() *Handler {
	return &Handler{
		services: map[string]service.Service{
			service.StravaName: service.NewStrava(),
			service.RwgpsName:  service.NewRwgps(),
			service.KomootName: service.NewKomoot(),
		},
	}
}

// GetRoutes returns all routes of given services
func (h *Handler) GetRoutes(serviceNames []string) ([]string, error) {
	routes := []string{}
	for _, serviceName := range serviceNames {
		if svc, found := h.services[serviceName]; found {
			newRoutes, err := svc.GetRoutes()
			if err != nil {
				return nil, err
			}
			routes = append(routes, newRoutes...)
		} else {
			return nil, fmt.Errorf("Service '%v' not found", serviceName)
		}
	}
	return routes, nil
}

// GetRoutesByUser returns user routes of given services
func (h *Handler) GetRoutesByUser(serviceNames []string, user string) ([]string, error) {
	routes := []string{}
	for _, serviceName := range serviceNames {
		if svc, found := h.services[serviceName]; found {
			newRoutes, err := svc.GetRoutesByUser(user)
			if err != nil {
				return nil, err
			}
			routes = append(routes, newRoutes...)
		} else {
			return nil, fmt.Errorf("Service '%v' not found", serviceName)
		}
	}
	return routes, nil
}

// GetAllRoutes returns routes across all services
func (h *Handler) GetAllRoutes() ([]string, error) {
	return h.GetRoutes(h.GetServices())
}

// GetAllRoutesByUser returns user routes across all services
func (h *Handler) GetAllRoutesByUser(user string) ([]string, error) {
	return h.GetRoutesByUser(h.GetServices(), user)
}

// GetAllUniqueRoutes returns unqiue routes across all services
func (h *Handler) GetAllUniqueRoutes() ([]string, error) {
	return h.GetUniqueRoutes(h.GetServices())
}

// GetUniqueRoutes returns unique routes given services
func (h *Handler) GetUniqueRoutes(serviceNames []string) ([]string, error) {
	routes, err := h.GetRoutes(serviceNames)
	if err != nil {
		return nil, err
	}
	return uniqueElements(routes), nil
}

// GetServices returns all services being handled
func (h *Handler) GetServices() []string {
	keys := make([]string, 0, len(h.services))
	for k := range h.services {
		keys = append(keys, k)
	}
	return keys
}

// uniqueElements returns a list of string without duplicates
func uniqueElements(routes []string) []string {
	// There should be a more idiomatic way to do this
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range routes {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
