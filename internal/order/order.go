package order

// Order interface defines the order to process a file.
type Order interface {
	GetInput() string
	GetOutputDir() string
}
