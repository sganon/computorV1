package reducer

import (
	"fmt"

	"github.com/sganon/computorV1/parser"
)

func PrintEquation(e parser.Equation) {
	for _, lmember := range e.LMembers {
		fmt.Printf("%s %f * %s^%d ", lmember.Operand, lmember.Coeff, lmember.Name, lmember.Exp)
	}
	if e.RMember.Coeff == 0 {
		fmt.Printf(" = 0\n")
	} else {
		fmt.Printf(" = %f * %s^%d\n", e.RMember.Coeff, e.RMember.Name, e.RMember.Exp)
	}
}

func ValidateEquation(e *parser.Equation) error {
	if e.RMember.Coeff != 0 {
		return fmt.Errorf("Validation Error: Right Member is not equal to 0")
	}
	names := map[string]bool{}
	exp := map[int]int{}
	for i, lmember := range e.LMembers {
		exp[lmember.Exp] = i
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

func reduceRMember(e *parser.Equation) {
	for i, lmember := range e.LMembers {
		if lmember.Exp == e.RMember.Exp {
			e.LMembers[i].Coeff -= e.RMember.Coeff
			e.RMember.Coeff = 0
		}
	}
}

func ReduceEquation(e *parser.Equation) {
	exp := map[int]int{}
	toDelete := []int{}
	for i, lmember := range e.LMembers {
		if _, pres := exp[lmember.Exp]; pres {
			e.LMembers[i].Coeff += e.LMembers[exp[lmember.Exp]].Coeff
			e.LMembers[exp[lmember.Exp]].Coeff = 0
			toDelete = append(toDelete, i)
		}
		exp[lmember.Exp] = i
	}
	for _, i := range toDelete {
		e.LMembers = append(e.LMembers[:i], e.LMembers[i+1:]...)
	}
	reduceRMember(e)
}
