package extractor

type Action string

const (
	// ActionResolved notifies that anime was resolved.
	ActionResolved Action = "resolved"

	// ActionRetry notifies that fetch will been retring.
	ActionRetry Action = "retry"

	// ActionSkip notifies that anime or list will be skipped.
	ActionSkip Action = "skip"

	// ActionStopExtractors notifies that parsing will be stopping.
	ActionStopExtractors Action = "stop_extractors"

	// ActionStopExtractor notifies that extracting will be stopping.
	ActionStopExtractor Action = "stop_extractor"

	// ActionStopParser notifies that all extractors will be stopping.
	ActionStopParser Action = "stop_parser"
)
