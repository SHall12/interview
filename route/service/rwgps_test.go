package service

import (
	"reflect"
	"testing"
)

func TestNewRwgps(t *testing.T) {
	var defaultRoutes = []string{"CVT", "Perkiomen", "Welsh Mountain"}

	svc := NewRwgps()
	routes, _ := svc.GetRoutes()

	if !reflect.DeepEqual(defaultRoutes, routes) {
		t.Fatalf("Routes do not match. Expected: %v, Returned: %v", defaultRoutes, routes)
	}
}

func TestRwgpsGetRoutesByUser(t *testing.T) {
	var testRoutes = []string{"SRT", "CVT", "Perkiomen"}
	var userName = "dude"
	var namedRoutes = []string{"SRTdude", "CVTdude", "Perkiomendude"}
	tt := []struct {
		name     string
		routes   []string
		user     string
		expected []string
	}{
		{
			name:     "No routes",
			user:     userName,
			routes:   nil,
			expected: []string{},
		},
		{
			name:     "Empty routes",
			user:     userName,
			routes:   []string{},
			expected: []string{},
		},
		{
			name:     "Single route",
			user:     userName,
			routes:   testRoutes[:1],
			expected: namedRoutes[:1],
		},
		{
			name:     "Multiple routes",
			user:     userName,
			routes:   testRoutes,
			expected: namedRoutes,
		},
		{
			name:     "Multiple routesAgain",
			user:     userName,
			routes:   testRoutes,
			expected: namedRoutes,
		},
		{
			name:     "No user",
			user:     "",
			routes:   testRoutes,
			expected: testRoutes,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			svc := rwgps{&DefaultService{routes: tc.routes}}
			routes, _ := svc.GetRoutesByUser(tc.user)

			if !reflect.DeepEqual(tc.expected, routes) {
				t.Errorf("Routes do not match. Expected: %v, Returned: %v", tc.expected, routes)
			}
		})
	}
}
