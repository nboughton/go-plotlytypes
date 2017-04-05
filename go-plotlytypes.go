package plotlytypes

/*
   Why are all axes strings? Because it's more flexible. You're passing this data to
	 Plotly.js which interprets it accordingly and where strconv has no value it returns
	 an empty string which plotly ignores but maintains the correct number of elements
	 in the set. This removes useless visual noise from the graph and allows you to make
	 your axes anything. Yes, there's a performance overhead to the conversion but it is
	 on the scale of a fraction of a fraction of a second, which I see as a reasonable
	 compromise
*/

// Dataset is the standard dataset for most graphs
type Dataset struct {
	X           []string `json:"x"`
	Y           []string `json:"y"`
	Z           []string `json:"z"`
	Name        string   `json:"name"`
	Mode        string   `json:"mode"`
	Type        string   `json:"type"`
	Line        Line     `json:"line"`
	Marker      Marker   `json:"marker"`
	ConnectGaps bool     `json:"connectgaps"`
}

// DatasetB is specifically for bubble graphs because it requires a distinctly different
// marker object
type DatasetB struct {
	X           []string `json:"x"`
	Y           []string `json:"y"`
	Z           []string `json:"z"`
	Name        string   `json:"name"`
	Mode        string   `json:"mode"`
	Type        string   `json:"type"`
	Line        Line     `json:"line"`
	Marker      MarkerB  `json:"marker"`
	ConnectGaps bool     `json:"connectgaps"`
}

// Marker defines the standard marker properties for a graph in Plotlyjs
type Marker struct {
	Colour  string  `json:"color"`
	Size    float64 `json:"size"`
	Line    Line    `json:"line"`
	Opacity float64 `json:"opacity"`
	Symbol  string  `json:"symbol"`
}

// MarkerB defines the properties of a Bubble graph marker where size values
// are a float64 array
type MarkerB struct {
	Colour   string    `json:"color"`
	Size     []float64 `json:"size"`
	SizeMode string    `json:"sizemode"`
	SizeRef  float64   `json:"sizeref"`
	Line     Line      `json:"line"`
	Opacity  float64   `json:"opacity"`
	Symbol   string    `json:"symbol"`
}

// Line defines the properties of a line datatype in Plotlyjs
type Line struct {
	Width   float64 `json:"width"`
	Colour  string  `json:"color"`
	Shape   string  `json:"shape"`
	Dash    string  `json:"dash"`
	Opacity float64 `json:"opacity"`
}
