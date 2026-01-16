// Package cv documents the constant and variable skills I hone, the job
// functions I’ve performed, and the types of projects I’ve worked on.
package cv

import "time"

// Skills that never become obsolete.
const (
	ContinuousLearning = iota
	ClearCommunication
	ProblemSolving
	SelfDiscipline
	Simplicity
	Kindness
	Humility
)

// Skills that can change over time.
var (
	Technologies = []string{"AWS", "Kubernetes", "Linux"}
	Languages    = []string{"Go", "Bash"} // + English, Italian, Czech, Slovak
)

// BackendEngineer builds and operates backend systems.
func BackendEngineer() {}

// DevOpsEngineer automates infrastructure and development workflows.
func DevOpsEngineer() {}

// SecurityEngineer protects systems and data through security tooling.
func SecurityEngineer() {}

// Observability provides insight into a system’s behavior.
type Observability struct {
	Logs    []string
	Metrics []map[time.Time]float64
	Traces  []map[string]time.Duration
}

// System represents a platform or API that provides a service.
type System struct {
	Observability Observability
	Documentation []string
}

// Build designs and constructs a system.
func Build(requirements, languages, technologies []string) *System { return &System{} }

// Operate runs and maintains the system.
func (s *System) Operate() error { return nil }

// Tool is a CLI utility that automates a task.
type Tool struct{}

// Write creates a tool using the given programming language.
func Write(language string) *Tool { return &Tool{} }

// Support helps users by communicating in a natural language.
func (t *Tool) Support(language string) error { return nil }
