// Package types provides shared types and structs.
package types

// Lock structure to hold terraform lock details
type Lock struct {
	ID        string `json:"ID" yaml:"ID"`
	Operation string `json:"Operation" yaml:"Operation"`
	Info      string `json:"Info" yaml:"Info"`
	Who       string `json:"Who" yaml:"Who"`
	Version   string `json:"Version" yaml:"Version"`
	Created   string `json:"Vreated" yaml:"Vreated"`
	Path      string `json:"Path" yaml:"Path"`
}
