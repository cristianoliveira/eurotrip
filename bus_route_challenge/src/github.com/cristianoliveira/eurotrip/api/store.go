package api

import (
	"io/ioutil"
	"strings"
)

/*
* QuickUnion implementation
*
* This algorithm takes more time preparing but when its done
* it enable us to find an union between two points with O(1)
*
 */
type Route struct {
	unions map[string]string
}

func NewRoute(stops []string) *Route {
	ro := &Route{unions: make(map[string]string)}

	var ini string
	for _, ini = range stops {
		ro.unions[ini] = ini
	}

	for _, dest := range stops {
		ro.Union(ini, dest)
		ini = dest
	}

	return ro
}

func (ro *Route) Connected(p, r string) bool {
	pid := ro.unions[p]
	rid := ro.unions[r]

	return len(rid) != 0 && rid == pid
}

func (ro *Route) Union(p, r string) {
	pid := ro.unions[p]
	rid := ro.unions[r]

	if pid == rid {
		return
	}

	for k, _ := range ro.unions {
		if ro.unions[k] == pid {
			ro.unions[k] = rid
		}
	}
}

func (ro *Route) ToString() string {
	routeAsString := ""

	for k, _ := range ro.unions {
		routeAsString += "->" + k
	}

	return routeAsString
}

type Itinerary struct {
	routes map[string]*Route
}

func (it *Itinerary) Connected(p, r string) bool {
	for _, route := range it.routes {
		if route.Connected(p, r) {
			return true
		}
	}

	return false
}

func (it *Itinerary) Routes() map[string]string {
	lines := make(map[string]string)

	for id, route := range it.routes {
		lines[id] = route.ToString()
	}

	return lines
}

func LoadItineraries(path string) (*Itinerary, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	routes, err := extractRoutes(string(data))
	if err != nil {
		return nil, err
	}

	itineraries := &Itinerary{routes: routes}

	return itineraries, nil
}

func extractRoutes(data string) (map[string]*Route, error) {
	routes := make(map[string]*Route)
	lines := strings.Split(data, "\n")

	for _, line := range lines[1:] {
		routeData := strings.Split(line, " ")
		routeId := routeData[0]
		routes[routeId] = NewRoute(routeData[1:])
	}

	return routes, nil
}
