package zone

import "github.com/umahmood/haversine"

const (
	ZoneThreshold  float64 = 20.0
	ScoreThreshold float32 = 80.0
)

var SchoolPos = Coord{0, 0}

func SchoolDistance(pos Coord) float64 {
	school := haversine.Coord{Lat: SchoolPos.Lat, Lon: SchoolPos.Long}
	home := haversine.Coord{Lat: pos.Lat, Lon: pos.Long}
	_, km := haversine.Distance(school, home)
	return km
}

func PassScore(score float32) bool {
	return score >= ScoreThreshold
}
