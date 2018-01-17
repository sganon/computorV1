package parser

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Member struct {
	name    string
	coeff   float64
	exp     int
	operand string
}

func parseMember(str string, memberPos int) (Member, error) {
	var member Member
	var err error
	str = strings.Trim(str, " ")
	fmt.Println(str)
	fSplit := strings.Split(str, "^")
	if len(fSplit) == 2 {
		member.name = string(fSplit[0][len(fSplit[0])-1])
		exp, err := strconv.Atoi(string(fSplit[1][0]))
		if err != nil {
			return member, fmt.Errorf("Error parsing member n°%d: Exp is not a int", memberPos)
		}
		member.exp = exp
	}

	switch string(str[0]) {
	case "-":
		member.operand = "-"
	default:
		member.operand = "+"
	}

	i := 0
	for unicode.IsDigit(rune(str[i])) || string(str[i]) == "+" || string(str[i]) == "-" || string(str[i]) == " " || string(str[i]) == "." {
		i++
	}

	coeff, err := strconv.ParseFloat(strings.Replace(string(str[0:i]), " ", "", -1), 64)
	if err != nil {
		return member, fmt.Errorf("Error parsing member n°%d: Coeff is not a int", memberPos)
	}
	member.coeff = coeff
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

func ParseEquation(equation string) error {
	var members []Member

	sides := strings.Split(equation, "=")
	if len(sides) > 2 || len(sides) == 1 {
		return fmt.Errorf("Equation malformatted")
	}
	positions := getOperandPos(sides[0])
	fmt.Println(positions)
	i := 0
	for i < 3 {
		var err error
		var member Member
		switch i {
		case 0:
			member, err = parseMember(string(sides[0][:positions[i]]), i+1)
			fmt.Println(member)
		case 1:
			member, err = parseMember(string(sides[0][positions[i-1]:positions[i]]), i+1)
		case 2:
			member, err = parseMember(string(sides[0][positions[i-1]:]), i+1)
		}
		if err != nil {
			return err
		}
		members = append(members, member)
		i++
	}
	rightMember, _ := parseMember(string(sides[1]), -1)
	fmt.Println(rightMember)
	return nil
}
