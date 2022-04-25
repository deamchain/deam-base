package eventcheck

import (
	"github.com/deamchain/deam-base/eventcheck/basiccheck"
	"github.com/deamchain/deam-base/eventcheck/epochcheck"
	"github.com/deamchain/deam-base/eventcheck/parentscheck"
	"github.com/deamchain/deam-base/inter/dag"
)

// Checkers is collection of all the checkers
type Checkers struct {
	Basiccheck   *basiccheck.Checker
	Epochcheck   *epochcheck.Checker
	Parentscheck *parentscheck.Checker
}

// Validate runs all the checks except Lachesis-related
func (v *Checkers) Validate(e dag.Event, parents dag.Events) error {
	if err := v.Basiccheck.Validate(e); err != nil {
		return err
	}
	if err := v.Epochcheck.Validate(e); err != nil {
		return err
	}
	if err := v.Parentscheck.Validate(e, parents); err != nil {
		return err
	}
	return nil
}
