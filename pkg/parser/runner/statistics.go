package runner

import (
	"fmt"

	"github.com/ilfey/hikilist-go/pkg/parser/statistic"
)

type Statistic struct {
	pages    uint
	fetched  uint
	resolved uint
	scipped  uint
}

func NewStatistic() statistic.RW {
	return &Statistic{}
}

func (s *Statistic) String() string {
	return fmt.Sprintf(
		"pages: %d, fetched: %d, resolved: %d, scipped: %d",
		s.pages,
		s.fetched,
		s.resolved,
		s.scipped,
	)
}

func (s *Statistic) Pages() uint {
	return s.pages
}

func (s *Statistic) AddPages(count uint) {
	s.pages += count
}

func (s *Statistic) Fetched() uint {
	return s.fetched
}

func (s *Statistic) AddFetched(count uint) {
	s.fetched += count
}

func (s *Statistic) Resolved() uint {
	return s.resolved
}

func (s *Statistic) AddResolved(count uint) {
	s.resolved += count
}

func (s *Statistic) Scipped() uint {
	return s.scipped
}

func (s *Statistic) AddScipped(count uint) {
	s.scipped += count
}
