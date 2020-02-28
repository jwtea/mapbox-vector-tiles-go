package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Client struct {
	token string
	URL   string
}

type VectorRequestOpts struct {
	tileset_id string
	zoom       int
	x          int
	y          int
	format     string
	style      string
}

func NewVectorRequestOpts() VectorRequestOpts {
	return VectorRequestOpts{
		tileset_id: "mapbox.mapbox-streets-v8",
		zoom:       10,
		x:          1,
		y:          1,
		format:     "vector.pbf",
		style:      "mapbox/light-v10",
	}
}

func (o VectorRequestOpts) toQuery() string {
	s := fmt.Sprintf("%s/%d/%d/%d.%s", o.tileset_id, o.zoom, o.x, o.y, o.format)
	return s
}

// NewClient returns
func NewClient(token string) *Client {
	return &Client{token, "https://api.mapbox.com/v4"}
}

//GetVectorTiles for given opts
func (c *Client) GetVectorTiles(dir string, opts VectorRequestOpts) {
	resource := fmt.Sprintf("%s/%s?access_token=%s", c.URL, opts.toQuery(), c.token)

	resp, err := http.Get(resource)

	if err != nil {
		log.Printf("Cannot get vector tiles: %s", opts.toQuery())
		return
	}

	if resp.StatusCode != http.StatusOK {
		log.Printf("Cannot get vector tiles: %s", opts.toQuery())
		return
	}
	c.CreateFile(dir, opts, resp)
}

func (c *Client) CreateFile(dir string, opts VectorRequestOpts, res *http.Response) {
	defer res.Body.Close()

	directory := fmt.Sprintf("%s/%d/%d/", dir, opts.zoom, opts.x)
	log.Printf("Dir: %s", directory)
	if err := os.MkdirAll(directory, 0777); err != nil {
		log.Fatal(err.Error())
	}
	filename := fmt.Sprintf("%d.%s", opts.y, opts.format)
	out, err := os.Create(fmt.Sprintf("%s/%s", directory, filename))
	if err != nil {
		log.Fatalf("Cannot create file: %s", err.Error())
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
}
