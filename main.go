package main

import (
	"log"
	"math"
	"os"

	"github.com/buckhx/tiles"
)

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	token := getenv("MAPBOX_API_TOKEN", "")
	c := NewClient(token)

	for zoom := 7; zoom <= 7; zoom++ {
		bee := int(math.Pow(float64(2), float64(zoom)))
		tile1 := tiles.Tile{0, 0, zoom}
		tile2 := tiles.Tile{bee, bee, zoom}
		log.Printf(" %#v %#v", tile1, tile2)
		tiles := searchTiles(tile1, tile2)

		log.Printf("Tiles: %d", len(tiles))

		for _, tile := range tiles {
			r := NewVectorRequestOpts()
			r.zoom = tile.Z
			r.x = tile.X
			r.y = tile.Y

			c.GetVectorTiles("data", r)
		}
	}

}
