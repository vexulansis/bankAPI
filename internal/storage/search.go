package storage

import (
	"bankAPI/internal/model"
	"database/sql"
	"fmt"
)

type Searcher struct {
	Data      any
	Table     string
	RefName   string
	Reference string
}
type SearchOption func(*Searcher)

// Searcher factory
func NewSearcher(opts ...SearchOption) *Searcher {
	s := &Searcher{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Search ...
func WithData(data any) SearchOption {
	return func(s *Searcher) {
		s.Data = data
	}
}

// Search in ...
func WithTable(table string) SearchOption {
	return func(s *Searcher) {
		s.Table = table
	}
}

// Search by ...
func WithReferenceName(refn string) SearchOption {
	return func(s *Searcher) {
		s.RefName = refn
	}
}

// Search(...)
func WithReference(ref string) SearchOption {
	return func(s *Searcher) {
		s.Reference = ref
	}
}

// Optional search in database
func (storage *Storage) Search(opts ...SearchOption) (*sql.Rows, error) {
	s := NewSearcher(opts...)
	switch s.Data.(type) {
	case *model.User:
		query := fmt.Sprintf("select * from %s where %s=$1", s.Table, s.RefName)
		rows, err := storage.DB.Query(query, s.Reference)
		if err != nil {
			return nil, err
		}
		return rows, nil
	}
	return nil, nil
}
