package diarium

// Output represents all logging outputs
// e.g: File, Terminal, ElasticSearch, Redis, etc.
type Output interface {
	output(LogBody) error
}

// output logs the given log into Output interface
func output(o Output, log LogBody) error {
	return o.output(log)
}
