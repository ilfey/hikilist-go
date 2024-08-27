package parser

import (
	"context"

	"github.com/ilfey/hikilist-go/pkg/parser/extractor"
	"github.com/ilfey/hikilist-go/pkg/parser/runner"
	"github.com/sirupsen/logrus"
)

type Parser struct {
	Logger *logrus.Logger

	Extractors []extractor.Extractor
}

func (parser *Parser) runExtractor(ctx context.Context, ex extractor.Extractor) error {
	r := runner.New(parser.Logger, ex)

	_, err := r.Run(ctx)

	return err
}

// Run parser
func (parser *Parser) Run(ctx context.Context) error {
	for _, ex := range parser.Extractors {
		err := parser.runExtractor(ctx, ex)
		if err != nil {
			return err
		}
	}

	return nil
}
