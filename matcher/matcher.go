package matcher

import (
	"fmt"
	"net/http"
)

// Matcher is the interface that determines if a request should be mirrored by a prism Server versus only proxied to the upstream.
type Matcher interface {
	Match(http.Request) bool
}

// NewMatcher creates a Matcher given its registered name and parameters.
func NewMatcher(name string, parameters string) (Matcher, error) {
	if mk, ok := matcherFuncs[name]; ok {
		return mk(parameters), nil
	}
	return nil, fmt.Errorf("no matcher registered with the name '%s'", name)
}

var matcherFuncs map[string]func(string) Matcher

// RegisterMatcherFunc associates a Matcher creating function with a name.
func RegisterMatcherFunc(name string, mkfunc func(string) Matcher) (alreadyRegistered bool) {
	_, alreadyRegistered = matcherFuncs[name]
	matcherFuncs[name] = mkfunc
	return alreadyRegistered
}
