package comparator

import "github.com/tmc/prism/httputils"

type Reporter interface {
	ReportDifference(different bool, responses ...httputils.RequestResponse)
}

type ResponseComparator interface {
	CompareResponses(reporter Reporter, responses ...httputils.RequestResponse)
}
