package order

// StripeOrder is the order to stripe a file.
type StripeOrder struct {
	Name   string
	Output string
}

// GetInput returns the input file name.
func (o *StripeOrder) GetInput() string {
	return o.Name
}

// GetOutputDir returns the output directory.
func (o *StripeOrder) GetOutputDir() string {
	return o.Output
}
