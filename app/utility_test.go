package app

import (
	"testing"
)

type inferSchemaDatumCaseStruct struct {
	arg1 string
	want bool
}

type inferSchemaCasesStruct struct {
	arg1 string
	want string
}

func TestIsInt(t *testing.T) {
	cases := []inferSchemaDatumCaseStruct{
		{"123", true},
		{"123.456", false},
		{"test string", false},
	}

	for _, element := range cases {
		datumStruct := datum{Input: element.arg1}
		got := datumStruct.isInt()
		if got != element.want {
			t.Errorf("wanted: %v, got: %v", element.want, got)
		}

	}
}

func TestIsFloat(t *testing.T) {
	cases := []inferSchemaDatumCaseStruct{
		{"123", true},
		{"456.008", true},
		{"test string", false},
	}

	for _, element := range cases {
		datumStruct := datum{Input: element.arg1}
		got := datumStruct.isFloat()
		if got != element.want {
			t.Errorf("wanted: %v, got: %v", element.want, got)
		}
	}

}

func TestIsBool(t *testing.T) {
	cases := []inferSchemaDatumCaseStruct{
		{"T", true},
		{"True", true},
		{"False", true},
		{"F", true},
		{"Tru", false},
		{"true", true},
		{"fals", false},
	}

	for _, element := range cases {
		datumStruct := datum{Input: element.arg1}
		got := datumStruct.isBool()
		if got != element.want {
			t.Errorf("wanted: %v, got: %v", element.want, got)
		}
	}
}

func TestIsDate(t *testing.T) {
	cases := []inferSchemaDatumCaseStruct{
		{"02-02-2022 15:00", true},
		{"02/02/2022 15:30:33", true},
		{"2022/03/01", true},
		{"Fri, 25 Feb 2022 15:24:42 GMT", false},
		{"test string", false},
	}

	for _, element := range cases {
		datumStruct := datum{Input: element.arg1}
		got := datumStruct.isDateTime()
		if got != element.want {
			t.Errorf("wanted: %v, got: %v", element.want, got)
		}
	}
}

func TestInferDatumSchema(t *testing.T) {
	cases := []inferSchemaCasesStruct{
		{"t", "bool"},
		{"test string", "string"},
		{"01-01-2001", "datetime"},
		{"123", "int"},
		{"45.008", "float"},
	}
	for _, element := range cases {
		got := inferDatumSchema(element.arg1)
		if got != element.want {
			t.Errorf("wanted: %v, got: %v", element.want, got)
		}
	}
}
