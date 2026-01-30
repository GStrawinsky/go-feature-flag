package flag

import (
	"time"
)

// ScheduledStrategy is the strategy to apply when a scheduled step is triggered.
type ScheduledStrategy = string

const (
	// ScheduledStrategyMerge merges the scheduled step configuration with the current flag configuration.
	// This is the default behavior.
	ScheduledStrategyMerge ScheduledStrategy = "merge"

	// ScheduledStrategyReset resets the flag configuration and applies only the scheduled step configuration.
	ScheduledStrategyReset ScheduledStrategy = "reset"
)

// ScheduledStep is one change of the flag.
type ScheduledStep struct {
	InternalFlag `yaml:",inline"`
	Date         *time.Time        `yaml:"date,omitempty" json:"date,omitempty" toml:"date,omitempty"`
	Strategy     ScheduledStrategy `yaml:"strategy,omitempty" json:"strategy,omitempty" toml:"strategy,omitempty"`
}

// GetStrategy returns the strategy for this scheduled step.
// If no strategy is set, it returns the default strategy (merge).
func (s *ScheduledStep) GetStrategy() ScheduledStrategy {
	if s.Strategy == "" {
		return ScheduledStrategyMerge
	}
	return s.Strategy
}
