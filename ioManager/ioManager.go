package iomanager

type IOManager interface {
	ReadLines() ([]string, error)
	WriteJSONToFile(data any) error
}
