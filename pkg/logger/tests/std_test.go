package tests

import (
	"context"
	"github.com/ilfey/hikilist-go/pkg/logger"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
	"testing"
)

type StdSuite struct {
	suite.Suite

	log *logger.File

	close func()
}

func (s *StdSuite) SetupSuite() {
	log, closeFn, err := logger.NewFile(context.Background(), 1, 1)
	if err != nil {
		panic(err)
	}

	s.log = log
	s.close = closeFn
}

func (s *StdSuite) TearDownSuite() {
	s.close()
}

func (s *StdSuite) TestLog() {
	s.NotPanics(func() {
		s.log.Log(errors.New("test Log"))
	})
}

func (s *StdSuite) TestLogPropagate() {
	s.NotPanics(func() {
		s.Error(
			s.log.Propagate(
				errors.New("test Propagate"),
			),
		)
	})
}

func (s *StdSuite) TestLogData() {
	s.NotPanics(func() {
		s.log.Object("test Object")
	})
}

func (s *StdSuite) TestInfo() {
	s.NotPanics(func() {
		s.log.Info("test Info")
	})
}

func (s *StdSuite) TestInfoPropagate() {
	s.NotPanics(func() {
		s.Error(
			s.log.InfoPropagate("test InfoPropagate"),
		)
	})
}

func (s *StdSuite) TestDebug() {
	s.NotPanics(func() {
		s.log.Debug("test Debug")
	})
}

func (s *StdSuite) TestDebugPropagate() {
	s.NotPanics(func() {
		s.Error(
			s.log.DebugPropagate("test DebugPropagate"),
		)
	})
}

func (s *StdSuite) TestWarn() {
	s.NotPanics(func() {
		s.log.Warn("test Warn")
	})
}

func (s *StdSuite) TestWarnPropagate() {
	s.NotPanics(func() {
		s.Error(
			s.log.WarnPropagate("test WarnPropagate"),
		)
	})
}

func (s *StdSuite) TestCritical() {
	s.NotPanics(func() {
		s.log.Critical("test Critical")
	})
}

func (s *StdSuite) TestCriticalPropagate() {
	s.NotPanics(func() {
		s.Error(
			s.log.CriticalPropagate("test CriticalPropagate"),
		)
	})
}

func TestStdSuite(t *testing.T) {
	suite.Run(t, new(StdSuite))
}
