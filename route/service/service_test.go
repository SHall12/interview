package service

import (
	"reflect"
	"testing"
)

func TestGetRoutes(t *testing.T) {
	var testRoutes = []string{"SRT", "CVT", "Perkiomen"}
	tt := []struct {
		name     string
		service  DefaultService
		expected []string
	}{
		{
			name:     "No routes",
			service:  DefaultService{},
			expected: nil,
		},
		{
			name:     "Empty routes",
			service:  DefaultService{routes: []string{}},
			expected: []string{},
		},
		{
			name:     "Single route",
			service:  DefaultService{routes: testRoutes[0:0]},
			expected: testRoutes[0:0],
		},
		{
			name:     "Multiple routes",
			service:  DefaultService{routes: testRoutes},
			expected: testRoutes,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			routes, _ := tc.service.getRoutes()

			if !reflect.DeepEqual(tc.expected, routes) {
				t.Errorf("Routes do not match. Expected: %v, Returned: %v", tc.expected, routes)
			}
		})
	}
}
