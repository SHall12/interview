package main

import (
	"fmt"
	"interview/route"
	"interview/route/service"
)

func main() {

	h := route.NewHandler()

	if routes, err := h.GetAllRoutes(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("All routes: %+q\n", routes)
	}

	if routes, err := h.GetAllUniqueRoutes(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Unique routes: %+q\n", routes)
	}

	user := "42"
	if routes, err := h.GetAllRoutesByUser(user); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("For user %v: %+q\n", user, routes)
	}

	services := []string{service.KomootName, service.RwgpsName}
	if routes, err := h.GetRoutesByUser(services, user); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("For user %v services %+q: %+q\n", user, services, routes)

	}
}
