package conf

import "regexp"

var (
	BuiltAt    string
	GoVersion  string
	GitAuthor  string
	GitCommit  string
	Version    string = "dev"
	WebVersion string
)

var (
	Conf *Config
)

var TypesMap = make(map[string][]string)
var PrivacyReg []*regexp.Regexp
