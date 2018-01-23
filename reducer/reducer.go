package reducer

import (
	"fmt"

	"github.com/sganon/computorV1/parser"
)

func ValidateEquation(e parser.Equation) error {
	if e.RMember.Coeff != 0 || e.RMember.Exp != 0 {
		return fmt.Errorf("Validation Error: Right Member is not equal to 0")
	}
	names := map[string]bool{}
	for _, lmember := range e.LMembers {
		if !names[lmember.Name] {
			names[lmember.Name] = true
		}
		if lmember.Exp > 2 && lmember.Coeff != 0 {
			return fmt.Errorf("Validation Error: Equation is superior to degree 2")
		}
	}
	if len(names) > 1 {
		return fmt.Errorf("Validation Error: Equation has more then one unknown")
	}
	return nil
}

func ReduceEquation(e *parser.Equation) {
	for i, lmember := range e.LMembers {
		if lmember.Exp == e.RMember.Exp {
			e.LMembers[i].Coeff -= e.RMember.Coeff
			e.RMember.Coeff = 0
		}
	}
}
