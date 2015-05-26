// gopack - Golang asset pipeline
// https://github.com/wcamarao/gopack
// MIT Licensed
//
package gopack

// Open-ended interface to define a custom pipeline,
// grouping assets by patterns, and allowing multiple
// processors to be applied per group.
//
// Group      - asset group or type name (e.g. JavaScripts).
//              Use CamelCase as you may want to export this
//              out into your templates.
//
// Patterns   - file patterns to match (e.g. "*.js"). The order
//              is relevant, e.g. {"*.module.js", "*.js"} will
//              load *.module.js before any other *.js matches.
//
// Exceptions - file patterns not to match. The order is not
//              relevant. The first matched exception is skipped.
//
// Processors - processors to be applied once a pattern matches.
//              Multiple processors are piped into one another.
//
type Matcher struct {
	Group      string
	Patterns   []string
	Exceptions []string
	Processors []Processor
}
