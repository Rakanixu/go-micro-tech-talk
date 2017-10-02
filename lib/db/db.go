package db

// Db ...
type Db interface {
	Init(url string) error
	Index(index string, docType string, id string, data string) error
	BulkIndex(index string, docType string, id string, data interface{})
	Read(index string, docType string, id string) (interface{}, error)
	Search(index string, docType string, query interface{}) ([]interface{}, error)
}

var s Db

// Register ..
func Register(storage Db) {
	s = storage
}

// Init ..
func Init(url string) error {
	return s.Init(url)
}

// Index ...
func Index(index, docType, id string, data string) error {
	return s.Index(index, docType, id, data)
}

// BulkIndex ...
func BulkIndex(index, docType, id string, data interface{}) {
	s.BulkIndex(index, docType, id, data)
}

// Read ...
func Read(index, docType, id string) (interface{}, error) {
	return s.Read(index, docType, id)
}

// Search ...
func Search(index, docType string, query interface{}) ([]interface{}, error) {
	return s.Search(index, docType, query)
}
