// Package cv documents the constant and variable skills
// I try to improve and the job functions I’ve performed.
package cv

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
	Infrastructure = []string{"AWS", "Kubernetes", "Linux"}
	Languages      = []string{"Go", "Bash"} // + English, Italian, Czech/Slovak
)

// Job allows me to provide for me and my family by applying skills to produce
// some value.
type Job func(skills []string) (money float64, value any)

// BackendEngineer builds and operates systems that are usually not used by end
// users.
type BackendEngineer Job

// DevOpsEngineer improves development workflows through automation, processes
// and tooling.
type DevOpsEngineer Job

// SecurityEngineer protects systems and data through hygiene, processes and
// tooling.
type SecurityEngineer Job
