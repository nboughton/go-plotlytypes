package plotlytypes

import "strconv"

// Dataset is the standard dataset for most graphs
type Dataset struct {
	X           Axis   `json:"x"`
	Y           Axis   `json:"y"`
	Z           Axis   `json:"z"`
	Name        string `json:"name"`
	Mode        string `json:"mode"`
	Type        string `json:"type"`
	Line        Line   `json:"line,omitempty"`
	Marker      Marker `json:"marker,omitempty"`
	ConnectGaps bool   `json:"connectgaps"`
}

// DatasetB is specifically for bubble graphs because it requires a distinctly different
// marker object
type DatasetB struct {
	X           Axis    `json:"x"`
	Y           Axis    `json:"y"`
	Z           Axis    `json:"z"`
	Name        string  `json:"name"`
	Mode        string  `json:"mode"`
	Type        string  `json:"type"`
	Line        Line    `json:"line,omitempty"`
	Marker      MarkerB `json:"marker,omitempty"`
	ConnectGaps bool    `json:"connectgaps"`
}

// Axis implements some convenience functions for managing numerical data
// represented as strings
type Axis []string

// AddInt increments Axis[idx] by y and returns a new Axis
// Usage: dataset.Y = dataset.Y.AddInt(1, 1)
func (a Axis) AddInt(idx, y int) (Axis, error) {
	x, err := strconv.Atoi(a[idx])
	if err != nil {
		x = 0
	}

	a[idx] = strconv.Itoa(x + y)
	return a, nil
}

// SubInt decrements Axis[idx] by y and returns a new Axis
// Usage: dataset.Y = dataset.Y.AddInt(1, 1)
func (a Axis) SubInt(idx, y int) (Axis, error) {
	x, err := strconv.Atoi(a[idx])
	if err != nil {
		x = 0
	}

	a[idx] = strconv.Itoa(x - y)
	return a, nil
}

// AppendInt adds y to the end of the Axis
func (a Axis) AppendInt(y int) Axis {
	a = append(a, strconv.Itoa(y))
	return a
}

// PrependInt adds y to the begining of the Axis
func (a Axis) PrependInt(y int) Axis {
	a = append([]string{strconv.Itoa(y)}, a...)
	return a
}

// AddFloat increments Axis[idx] by y
func (a Axis) AddFloat(idx int, y float64) (Axis, error) {
	x, err := strconv.ParseFloat(a[idx], 64)
	if err != nil {
		return a, err
	}

	a[idx] = strconv.FormatFloat(x+y, 'f', -1, 64)
	return a, nil
}

// SubFloat decrements Axis[idx] by y
func (a Axis) SubFloat(idx int, y float64) (Axis, error) {
	x, err := strconv.ParseFloat(a[idx], 64)
	if err != nil {
		return a, err
	}

	a[idx] = strconv.FormatFloat(x-y, 'f', -1, 64)
	return a, nil
}

// AppendFloat adds y to the end of the Axis
func (a Axis) AppendFloat(y float64) Axis {
	a = append(a, strconv.FormatFloat(y, 'f', -1, 64))
	return a
}

// PrependFloat adds y to the beginning of the Axis
func (a Axis) PrependFloat(y float64) Axis {
	a = append([]string{strconv.FormatFloat(y, 'f', -1, 64)}, a...)
	return a
}

// AppendStr adds string y to the end of the Axis
func (a Axis) AppendStr(y string) Axis {
	a = append(a, y)
	return a
}

// PrependStr adds string y to the beginning of the Axis
func (a Axis) PrependStr(y string) Axis {
	a = append([]string{y}, a...)
	return a
}

// Line defines the properties of a line datatype in Plotlyjs
type Line struct {
	Width   float64 `json:"width,omitempty"`
	Colour  string  `json:"color,omitempty"`
	Shape   string  `json:"shape,omitempty"`
	Dash    string  `json:"dash,omitempty"`
	Opacity float64 `json:"opacity,omitempty"`
}

// Marker defines the standard marker properties for a graph in Plotlyjs
type Marker struct {
	Colour  string  `json:"color,omitempty"`
	Size    float64 `json:"size,omitempty"`
	Line    Line    `json:"line,omitempty"`
	Opacity float64 `json:"opacity,omitempty"`
	Symbol  string  `json:"symbol,omitempty"`
}

// MarkerB defines the properties of a Bubble graph marker where size values
// are a float64 array
type MarkerB struct {
	Colour   string    `json:"color,omitempty"`
	Size     []float64 `json:"size,omitempty"`
	SizeMode string    `json:"sizemode,omitempty"`
	SizeRef  float64   `json:"sizeref,omitempty"`
	Line     Line      `json:"line,omitempty"`
	Opacity  float64   `json:"opacity,omitempty"`
	Symbol   string    `json:"symbol,omitempty"`
}
