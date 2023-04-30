package flags

import "strconv"

type stringFlag struct {
	value string
	isSet bool
}

func NewStringFlag(value string) *stringFlag {
	return &stringFlag{value: value}
}

func (sf *stringFlag) String() string {
	return sf.value
}

func (sf *stringFlag) Set(value string) error {
	sf.value = value
	sf.isSet = true
	return nil
}

func (sf *stringFlag) IsSet() bool {
	return sf.isSet
}

type intFlag struct {
	value int
	isSet bool
}

func NewIntFlag(value int) *intFlag {
	return &intFlag{value: value}
}

func (ifl *intFlag) Int() int {
	return ifl.value
}

func (ifl *intFlag) Set(value string) error {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	ifl.value = intValue
	ifl.isSet = true
	return nil
}

func (ifl *intFlag) String() string {
	return strconv.Itoa(ifl.value)
}

func (ifl *intFlag) IsSet() bool {
	return ifl.isSet
}

type float64Flag struct {
	value float64
	isSet bool
}

func NewFloat64Flag(value float64) *float64Flag {
	return &float64Flag{value: value}
}

func (ff *float64Flag) Float64() float64 {
	return ff.value
}

func (ff *float64Flag) Set(value string) error {
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return err
	}
	ff.value = floatValue
	ff.isSet = true
	return nil
}

func (ff *float64Flag) String() string {
	return strconv.FormatFloat(ff.value, 'f', -1, 64)
}

func (ff *float64Flag) IsSet() bool {
	return ff.isSet
}
