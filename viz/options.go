package viz

type Options struct {
	Positional struct {
		Width  int    `json:"width" description:"Width of the visualization in pixels" default:"10"`
		Height int    `json:"height" long:"height" description:"Height of the visualization in pixels" default:"10"`
		Shape  string `json:"shape" default:"square"`
	} `positional-args:"yes" required:"no" description:"Positional arguments for the visualization options" json:"positional"`
}
