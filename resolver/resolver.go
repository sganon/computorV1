package resolver

import (
	"fmt"
	"math"

	"github.com/sganon/computorV1/parser"
)

func ResolveEquation(e parser.Equation) {
	vars := map[string]float64{}
	for _, lmember := range e.LMembers {
		switch lmember.Exp {
		case 2:
			vars["a"] = lmember.Coeff
		case 1:
			vars["b"] = lmember.Coeff
		case 0:
			vars["c"] = lmember.Coeff
		}
	}

	if _, pres := vars["b"]; !pres {
		if vars["c"] == e.RMember.Coeff {
			fmt.Println("Every number is a solution to this equation")
		} else {
			fmt.Println("No solution to this equation")
		}
		return
	}

	if _, pres := vars["a"]; !pres {
		fmt.Println("Polynomial degree 1")
		fmt.Printf("The solution is:\n%f\n", vars["c"]/-vars["b"])
	} else if _, pres := vars["a"]; pres {
		delta := math.Pow(vars["b"], 2) - 4*vars["a"]*vars["c"]
		if delta < 0 {
			fmt.Println("Discriminant is strictly negative, no solution")
		} else if delta == 0 {
			sol := -(vars["b"] / (2 * vars["a"]))
			fmt.Println("Discriminant is strictly egal to zero, the solution is:")
			fmt.Printf("%f\n", sol)
		} else {
			sol1 := (-vars["b"] - math.Sqrt(delta)) / (2 * vars["a"])
			sol2 := (-vars["b"] + math.Sqrt(delta)) / (2 * vars["a"])
			fmt.Println("Discriminant is strictly positive, the two solutions are:")
			fmt.Printf("%f\n%f\n", sol1, sol2)
		}
	}
}
