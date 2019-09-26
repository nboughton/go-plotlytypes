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
	Line        Line   `json:"line"`
	Marker      Marker `json:"marker"`
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
	Line        Line    `json:"line"`
	Marker      MarkerB `json:"marker"`
	ConnectGaps bool    `json:"connectgaps"`
}

// Axis implements some convenience functions for managing numerical data
// represented as strings
type Axis []string

// AddInt increments Axis[idx] by y
func (a Axis) AddInt(idx, y int) error {
	x, err := strconv.Atoi(a[idx])
	if err != nil {
		return err
	}

	a[idx] = strconv.Itoa(x + y)
	return nil
}

// SubInt decrements Axis[idx] by y
func (a Axis) SubInt(idx, y int) error {
	x, err := strconv.Atoi(a[idx])
	if err != nil {
		return err
	}

	a[idx] = strconv.Itoa(x - y)
	return nil
}

// AppendInt adds y to the end of the Axis
func (a Axis) AppendInt(y int) {
	a = append(a, strconv.Itoa(y))
}

// PrependInt adds y to the begining of the Axis
func (a Axis) PrependInt(y int) {
	a = append([]string{strconv.Itoa(y)}, a...)
}

// AddFloat increments Axis[idx] by y
func (a Axis) AddFloat(idx int, y float64) error {
	x, err := strconv.ParseFloat(a[idx], 64)
	if err != nil {
		return err
	}

	a[idx] = strconv.FormatFloat(x+y, 'f', -1, 64)
	return nil
}

// SubFloat decrements Axis[idx] by y
func (a Axis) SubFloat(idx int, y float64) error {
	x, err := strconv.ParseFloat(a[idx], 64)
	if err != nil {
		return err
	}

	a[idx] = strconv.FormatFloat(x-y, 'f', -1, 64)
	return nil
}

// AppendFloat adds y to the end of the Axis
func (a Axis) AppendFloat(y float64) {
	a = append(a, strconv.FormatFloat(y, 'f', -1, 64))
}

// PrependFloat adds y to the beginning of the Axis
func (a Axis) PrependFloat(y float64) {
	a = append([]string{strconv.FormatFloat(y, 'f', -1, 64)}, a...)
}

// AppendStr adds string y to the end of the Axis
func (a Axis) AppendStr(y string) {
	a = append(a, y)
}

// PrependStr adds string y to the beginning of the Axis
func (a Axis) PrependStr(y string) {
	a = append([]string{y}, a...)
}

// Line defines the properties of a line datatype in Plotlyjs
type Line struct {
	Width   float64 `json:"width"`
	Colour  string  `json:"color"`
	Shape   string  `json:"shape"`
	Dash    string  `json:"dash"`
	Opacity float64 `json:"opacity"`
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
