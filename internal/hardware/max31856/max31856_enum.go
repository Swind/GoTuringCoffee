// Code generated by go-enum
// DO NOT EDIT!

package max31856

import (
	"fmt"
)

const (
	// ModeManual is a Mode of type Manual
	ModeManual Mode = iota
	// ModeAutomatic is a Mode of type Automatic
	ModeAutomatic
)

const _ModeName = "ManualAutomatic"

var _ModeMap = map[Mode]string{
	0: _ModeName[0:6],
	1: _ModeName[6:15],
}

func (i Mode) String() string {
	if str, ok := _ModeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Mode(%d)", i)
}

var _ModeValue = map[string]Mode{
	_ModeName[0:6]:  0,
	_ModeName[6:15]: 1,
}

// ParseMode attempts to convert a string to a Mode
func ParseMode(name string) (Mode, error) {
	if x, ok := _ModeValue[name]; ok {
		return Mode(x), nil
	}
	return Mode(0), fmt.Errorf("%s is not a valid Mode", name)
}

const (
	// SampleAvg1 is a Sample of type Avg1
	SampleAvg1 Sample = iota
	// SampleAvg2 is a Sample of type Avg2
	SampleAvg2
	// SampleAvg4 is a Sample of type Avg4
	SampleAvg4
	// SampleAvg8 is a Sample of type Avg8
	SampleAvg8
	// SampleAvg16 is a Sample of type Avg16
	SampleAvg16
)

const _SampleName = "Avg1Avg2Avg4Avg8Avg16"

var _SampleMap = map[Sample]string{
	0: _SampleName[0:4],
	1: _SampleName[4:8],
	2: _SampleName[8:12],
	3: _SampleName[12:16],
	4: _SampleName[16:21],
}

func (i Sample) String() string {
	if str, ok := _SampleMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Sample(%d)", i)
}

var _SampleValue = map[string]Sample{
	_SampleName[0:4]:   0,
	_SampleName[4:8]:   1,
	_SampleName[8:12]:  2,
	_SampleName[12:16]: 3,
	_SampleName[16:21]: 4,
}

// ParseSample attempts to convert a string to a Sample
func ParseSample(name string) (Sample, error) {
	if x, ok := _SampleValue[name]; ok {
		return Sample(x), nil
	}
	return Sample(0), fmt.Errorf("%s is not a valid Sample", name)
}

const (
	// TypeB is a Type of type B
	TypeB Type = iota
	// TypeE is a Type of type E
	TypeE
	// TypeJ is a Type of type J
	TypeJ
	// TypeK is a Type of type K
	TypeK
	// TypeN is a Type of type N
	TypeN
	// TypeR is a Type of type R
	TypeR
	// TypeS is a Type of type S
	TypeS
	// TypeT is a Type of type T
	TypeT
)

const _TypeName = "BEJKNRST"

var _TypeMap = map[Type]string{
	0: _TypeName[0:1],
	1: _TypeName[1:2],
	2: _TypeName[2:3],
	3: _TypeName[3:4],
	4: _TypeName[4:5],
	5: _TypeName[5:6],
	6: _TypeName[6:7],
	7: _TypeName[7:8],
}

func (i Type) String() string {
	if str, ok := _TypeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Type(%d)", i)
}

var _TypeValue = map[string]Type{
	_TypeName[0:1]: 0,
	_TypeName[1:2]: 1,
	_TypeName[2:3]: 2,
	_TypeName[3:4]: 3,
	_TypeName[4:5]: 4,
	_TypeName[5:6]: 5,
	_TypeName[6:7]: 6,
	_TypeName[7:8]: 7,
}

// ParseType attempts to convert a string to a Type
func ParseType(name string) (Type, error) {
	if x, ok := _TypeValue[name]; ok {
		return Type(x), nil
	}
	return Type(0), fmt.Errorf("%s is not a valid Type", name)
}
