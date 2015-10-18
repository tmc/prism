package comparator

import "github.com/tmc/prism/httputils"

// Reporter is the interface that is invoked when a response is considered different.
type Reporter interface {
	ReportDifference(different bool, responses ...httputils.RequestResponse)
}

// ResponseComparator considers a set of responses and if a difference is detected it
// should be reported to the provided Reporter.
type ResponseComparator interface {
	CompareResponses(reporter Reporter, responses ...httputils.RequestResponse)
}
