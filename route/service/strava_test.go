package service

import (
	"reflect"
	"testing"
)

func TestNewStrava(t *testing.T) {
	var defaultRoutes = []string{"SRT", "CVT", "Perkiomen"}

	svc := NewStrava()
	routes, _ := svc.getRoutes()

	if !reflect.DeepEqual(defaultRoutes, routes) {
		t.Fatalf("Routes do not match. Expected: %v, Returned: %v", defaultRoutes, routes)
	}
}

func TestStravaGetRoutesByUser(t *testing.T) {
	var testRoutes = []string{"SRT", "CVT", "Perkiomen"}
	var userName = "dude"
	var namedRoutes = []string{"dudeSRT", "dudeCVT", "dudePerkiomen"}
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
			routes:   testRoutes[0:0],
			expected: namedRoutes[0:0],
		},
		{
			name:     "Multiple routes",
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
			svc := strava{&DefaultService{routes: tc.routes}}
			routes, _ := svc.getRoutesByUser(tc.user)

			if !reflect.DeepEqual(tc.expected, routes) {
				t.Errorf("Routes do not match. Expected: %v, Returned: %v", tc.expected, routes)
			}
		})
	}
}
