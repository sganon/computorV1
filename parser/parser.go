package parser

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Member struct {
	Name    string
	Coeff   float64
	Exp     int
	Operand string
}

type Equation struct {
	LMembers []Member
	RMember  Member
}

func parseMember(str string, memberPos int) (Member, error) {
	var member Member
	var err error
	str = strings.Trim(str, " ")
	fSplit := strings.Split(str, "^")
	if len(fSplit) == 2 {
		member.Name = string(fSplit[0][len(fSplit[0])-1])
		Exp, err := strconv.Atoi(string(fSplit[1][0]))
		if err != nil {
			return member, fmt.Errorf("Error parsing member n°%d: Exp is not a int", memberPos)
		}
		member.Exp = Exp
	}

	switch string(str[0]) {
	case "-":
		member.Operand = "-"
	default:
		member.Operand = "+"
	}

	i := 0
	for unicode.IsDigit(rune(str[i])) || string(str[i]) == "+" || string(str[i]) == "-" || string(str[i]) == " " || string(str[i]) == "." {
		i++
	}

	Coeff, err := strconv.ParseFloat(strings.Replace(string(str[0:i]), " ", "", -1), 64)
	if err != nil {
		return member, fmt.Errorf("Error parsing member n°%d: Coeff is not a int", memberPos)
	}
	member.Coeff = Coeff
	return member, nil
}

func getOperandPos(side string) []int {
	var positons []int

	for i, c := range side {
		if string(c) == "+" || string(c) == "-" {
			positons = append(positons, i)
		}
	}
	return positons
}

func ParseEquation(equation string) (Equation, error) {
	var members []Member

	sides := strings.Split(equation, "=")
	if len(sides) > 2 || len(sides) == 1 {
		return Equation{}, fmt.Errorf("Equation malformatted")
	}
	positions := getOperandPos(sides[0])
	i := 0
	for i < 3 {
		var err error
		var member Member
		switch i {
		case 0:
			member, err = parseMember(string(sides[0][:positions[i]]), i+1)
		case 1:
			member, err = parseMember(string(sides[0][positions[i-1]:positions[i]]), i+1)
		case 2:
			member, err = parseMember(string(sides[0][positions[i-1]:]), i+1)
		}
		if err != nil {
			return Equation{}, err
		}
		members = append(members, member)
		i++
	}
	rightMember, _ := parseMember(string(sides[1]), -1)
	return Equation{LMembers: members, RMember: rightMember}, nil
}
