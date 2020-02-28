package main

import (
	"math"

	"github.com/buckhx/tiles"
)

func searchLatLngs(zoom int, lon1 float64, lat1 float64, lon2 float64, lat2 float64) [][]float64 {
	interval := 360 / math.Exp2(float64(zoom))
	interval = interval / 2
	coords := [][]float64{}
	for i := lon1; i < lon2+interval; i += interval {
		for j := lat1; j < lat2+interval; j += interval {
			coord := []float64{i, j}
			coords = append(coords, coord)
		}
	}
	return coords
}

func searchTiles(tile1 tiles.Tile, tile2 tiles.Tile) (tileSlice []tiles.Tile) {
	for x := tile1.X; x < tile2.X; x++ {
		for y := tile1.Y; y < tile2.Y; y++ {
			tileSlice = append(tileSlice, tiles.Tile{x, y, tile1.Z})
		}
	}
	return tileSlice
}
