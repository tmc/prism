package matcher

import (
	"fmt"
	"net/http"
)

type Matcher interface {
	Match(http.Request) bool
}

func NewMatcher(name string, parameters string) (Matcher, error) {
	if mk, ok := matcherFuncs[name]; ok {
		return mk(parameters), nil
	}
	return nil, fmt.Errorf("no matcher registered with the name '%s'", name)
}

var matcherFuncs map[string]func(string) Matcher

func RegisterMatcherFunc(name string, mkfunc func(string) Matcher) (alreadyRegistered bool) {
	_, alreadyRegistered = matcherFuncs[name]
	matcherFuncs[name] = mkfunc
	return alreadyRegistered
}
