package calculator

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Person struct {
	name string
	city string
}

func calcArea(r float64) (float64, error) {
	if r <= 0 {
		return 0, errors.New("invalid input")
	}
	twoDec := fmt.Sprintf("%.2f", math.Pi*(math.Pow(r, 2)))
	return strconv.ParseFloat(twoDec, 64)
}

func calcVolume(r float64, h float64) (float64, error) {

	if r <= 0 || h <= 0 {
		return 0, errors.New("invalid input")
	}
	twoDec := fmt.Sprintf("%.2f", math.Pi*(math.Pow(r, 2)*h))
	return strconv.ParseFloat(twoDec, 64)
}

func isFirstLetA(p Person) (bool, error) {
	p.name = strings.ToLower(strings.Trim(p.name, " "))
	if len(p.name) > 0 {
		if p.name[0] == 'a' {
			return true, nil
		} else {
			return false, nil
		}
	}
	return false, errors.New("no input")
}

func searchPersonSameCity(p []Person, qCity string) (matchPerson []string, err error) {
	if len(p) < 1 {
		err = errors.New("person is empty")
	}

	for _, person := range p {
		personCity := person.city
		if strings.EqualFold(qCity, personCity) {
			matchPerson = append(matchPerson, person.name)
		}
	}
	return
}
