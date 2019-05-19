package route

import (
	"errors"
	"interview/route/service"
	"reflect"
	"sort"
	"testing"
)

var (
	firstRoutes  = []string{"SRT", "CVT", "Perkiomen"}
	secondRoutes = []string{"CVT", "Perkiomen", "Welsh Mountain"}
	firstService = service.MockService{
		GetRoutesMock: func() ([]string, error) {
			return firstRoutes, nil
		},
		GetRoutesByUserMock: func() ([]string, error) {
			return firstRoutes, nil
		},
	}

	secondService = service.MockService{
		GetRoutesMock: func() ([]string, error) {
			return secondRoutes, nil
		},
		GetRoutesByUserMock: func() ([]string, error) {
			return secondRoutes, nil
		},
	}

	errorService = service.MockService{
		GetRoutesMock: func() ([]string, error) {
			return nil, errors.New("error")
		},
		GetRoutesByUserMock: func() ([]string, error) {
			return nil, errors.New("error")
		},
	}

	dummyServices = map[string]service.Service{
		"first":  &firstService,
		"second": &secondService,
		"error":  &errorService,
	}
)

func TestNewRouteHandler(t *testing.T) {

	serviceList := []string{service.StravaName, service.RwgpsName, service.KomootName}
	handler := NewHandler()

	for _, svc := range serviceList {
		if val, found := handler.services[svc]; found {
			if val == nil {
				t.Errorf("Service not instantiated: %v", svc)
			}
		} else {
			t.Errorf("Service not found: %v", svc)
		}
	}
}

func TestGetListOfServices(t *testing.T) {
	var dummyServiceNames = []string{"error", "first", "second"}

	h := Handler{dummyServices}
	actualNames := h.GetServices()
	sort.Strings(actualNames)

	if !reflect.DeepEqual(dummyServiceNames, actualNames) {
		t.Fatalf("Service names do not match. Expected: %v, Returned: %v", dummyServiceNames, actualNames)
	}
}

func TestUniqueElements(t *testing.T) {
	tt := []struct {
		name     string
		list     []string
		expected []string
	}{
		{
			name:     "Null list",
			list:     nil,
			expected: []string{},
		},
		{
			name:     "Empty list",
			list:     []string{},
			expected: []string{},
		},
		{
			name:     "Single element",
			list:     []string{"foo"},
			expected: []string{"foo"},
		},
		{
			name:     "Duplicates",
			list:     []string{"foo", "foo", "bar", "czar", "bar"},
			expected: []string{"foo", "bar", "czar"},
		},
		{
			name:     "No Duplicates",
			list:     []string{"foo", "car", "bar", "czar", "far"},
			expected: []string{"foo", "car", "bar", "czar", "far"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			dedupped := uniqueElements(tc.list)

			if !reflect.DeepEqual(tc.expected, dedupped) {
				t.Errorf("List not unique. Expected: %v, Returned: %v", tc.expected, dedupped)
			}
		})
	}
}

func TestGetRoutes(t *testing.T) {
	tt := []struct {
		name        string
		services    map[string]service.Service
		serviceList []string
		expected    []string
	}{
		{
			name:        "No services selected",
			services:    dummyServices,
			serviceList: []string{},
			expected:    []string{},
		},
		{
			name:        "Single service available and selected",
			services:    map[string]service.Service{"first": &firstService},
			serviceList: []string{"first"},
			expected:    firstRoutes,
		},
		{
			name:        "Multiple services selected",
			services:    dummyServices,
			serviceList: []string{"first", "second"},
			expected:    append(firstRoutes, secondRoutes...),
		},
		{
			name:        "Error thrown on bad service",
			services:    dummyServices,
			serviceList: []string{"error"},
			expected:    nil,
		},
		{
			name:        "Error thrown on nonexistant service",
			services:    dummyServices,
			serviceList: []string{"dontexist"},
			expected:    nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			handler := Handler{services: tc.services}
			routes, err := handler.GetRoutes(tc.serviceList)

			if err != nil && routes != nil {
				t.Errorf("Unexpected values %v returned with error", routes)
			}

			if !reflect.DeepEqual(tc.expected, routes) {
				t.Errorf("Routes do not match. Expected: %v, Returned: %v", tc.expected, routes)
			}
		})
	}
}

func TestGetRoutesByUser(t *testing.T) {
	tt := []struct {
		name        string
		services    map[string]service.Service
		user        string
		serviceList []string
		expected    []string
	}{
		{
			name:        "No services selected",
			services:    dummyServices,
			serviceList: []string{},
			expected:    []string{},
		},
		{
			name:        "Single service available and selected",
			services:    map[string]service.Service{"first": &firstService},
			serviceList: []string{"first"},
			expected:    firstRoutes,
		},
		{
			name:        "Multiple services selected",
			services:    dummyServices,
			serviceList: []string{"first", "second"},
			expected:    append(firstRoutes, secondRoutes...),
		},
		{
			name:        "Error thrown on bad service",
			services:    dummyServices,
			serviceList: []string{"error"},
			expected:    nil,
		},
		{
			name:        "Error thrown on nonexistant service",
			services:    dummyServices,
			serviceList: []string{"dontexist"},
			expected:    nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			handler := Handler{services: tc.services}
			routes, err := handler.GetRoutesByUser(tc.serviceList, tc.user)

			if err != nil && routes != nil {
				t.Errorf("Unexpected values %v returned with error", routes)
			}

			if !reflect.DeepEqual(tc.expected, routes) {
				t.Errorf("Routes do not match. Expected: %v, Returned: %v", tc.expected, routes)
			}
		})
	}
}

/*
 Use HandlerMock to test:
	GetAllRoutes
	GetAllRoutesByUser
	GetAllUniqueRoutes
	GetUniqueRoutes
*/
