package order

// StripOrder is the order to strip a file.
type StripOrder struct {
	Name   string
	Output string
}

// GetInput returns the input file name.
func (o *StripOrder) GetInput() string {
	return o.Name
}

// GetOutputDir returns the output directory.
func (o *StripOrder) GetOutputDir() string {
	return o.Output
}
