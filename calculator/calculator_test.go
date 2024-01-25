package calculator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcArea(t *testing.T) {

	tc := []struct {
		name           string
		input          float64
		expectedOutput float64
		expectedError  error
	}{
		{"correct output", 2.3, 16.62, nil},
		{"negative input", -2.3, 0, errors.New("invalid input")},
		{"zero input", 0, 0, errors.New("invalid input")},
	}

	assert := assert.New(t)
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			want := test.expectedOutput
			got, err := calcArea(test.input)

			assert.Equal(test.expectedError, err)
			assert.Equal(want, got)
		})
	}
}

func TestCalcVolume(t *testing.T) {
	tc := []struct {
		name           string
		radius         float64
		height         float64
		expectedOutput float64
		expectedError  error
	}{
		{"calculate cylinder", 2.3, 4.5, 74.79, nil},
		{"negative radius", -2.3, 4.5, 0, errors.New("invalid input")},
		{"negative height", 2.3, -4.5, 0, errors.New("invalid input")},
		{"zero radius", 0, 4.5, 0, errors.New("invalid input")},
		{"zero height", 2.3, 0, 0, errors.New("invalid input")},
	}
	assert := assert.New(t)
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			want := test.expectedOutput
			got, err := calcVolume(test.radius, test.height)

			assert.Equal(test.expectedError, err)
			assert.Equal(want, got)
		})
	}
}

func TestCheckFirstLetter(t *testing.T) {
	tc := []struct {
		name           string
		person         Person
		expectedOutput bool
		expectedError  error
	}{
		{"Empty string", Person{"", "Biñan"}, false, errors.New("no input")},
		{"With A", Person{"Ambayec", "Biñan"}, true, nil},
		{"Not A", Person{"Patrick", "Biñan"}, false, nil},
		{"lowercase A", Person{"ambayec", "Biñan"}, true, nil},
	}
	assert := assert.New(t)
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			want := test.expectedOutput
			got, err := isFirstLetA(test.person)

			assert.Equal(test.expectedError, err)
			assert.Equal(want, got)
		})
	}
}

func TestSearchPersonSameCity(t *testing.T) {
	tc := []struct {
		name           string
		person         []Person
		city           string
		expectedOutput []string
		expectedError  error
	}{
		{"empty person", nil, "biñan", nil, errors.New("person is empty")},
		{"return empty person", []Person{{"Patrick", "Biñan"}, {"Ambayec", "Biñan"}}, "Taguig", nil, nil},
		{"return 2 person", []Person{{"Patrick", "Biñan"}, {"Ambayec", "Biñan"}}, "Biñan", []string{"Patrick", "Ambayec"}, nil},
	}

	assert := assert.New(t)
	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			want := test.expectedOutput
			got, err := countPersonSameCity(test.person, test.city)

			assert.Equal(test.expectedError, err)
			assert.Equal(want, got)
		})
	}
}
