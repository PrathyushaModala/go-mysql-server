// Copyright 2022 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/dolthub/vitess/go/sqltypes"
	"github.com/dolthub/vitess/go/vt/proto/query"
	"github.com/shopspring/decimal"

	"github.com/dolthub/go-mysql-server/sql"
)

var systemBoolValueType = reflect.TypeOf(int8(0))

// SystemBoolType_ is an internal boolean type ONLY for system variables.
type SystemBoolType_ struct {
	varName string
}

var _ sql.SystemVariableType = SystemBoolType_{}
var _ sql.CollationCoercible = SystemBoolType_{}

// NewSystemBoolType returns a new systemBoolType.
func NewSystemBoolType(varName string) sql.SystemVariableType {
	return SystemBoolType_{varName}
}

// Compare implements Type interface.
func (t SystemBoolType_) Compare(a interface{}, b interface{}) (int, error) {
	as, err := t.Convert(a)
	if err != nil {
		return 0, err
	}
	bs, err := t.Convert(b)
	if err != nil {
		return 0, err
	}
	ai := as.(int8)
	bi := bs.(int8)

	if ai == bi {
		return 0, nil
	}
	if ai < bi {
		return -1, nil
	}
	return 1, nil
}

// Convert implements Type interface.
func (t SystemBoolType_) Convert(v interface{}) (interface{}, error) {
	// Nil values are not accepted
	switch value := v.(type) {
	case bool:
		if value {
			return int8(1), nil
		}
		return int8(0), nil
	case int:
		return t.Convert(int64(value))
	case uint:
		return t.Convert(int64(value))
	case int8:
		return t.Convert(int64(value))
	case uint8:
		return t.Convert(int64(value))
	case int16:
		return t.Convert(int64(value))
	case uint16:
		return t.Convert(int64(value))
	case int32:
		return t.Convert(int64(value))
	case uint32:
		return t.Convert(int64(value))
	case int64:
		if value == 0 || value == 1 {
			return int8(value), nil
		}
	case uint64:
		return t.Convert(int64(value))
	case float32:
		return t.Convert(float64(value))
	case float64:
		// Float values aren't truly accepted, but the engine will give them when it should give ints.
		// Therefore, if the float doesn't have a fractional portion, we treat it as an int.
		if value == float64(int64(value)) {
			return t.Convert(int64(value))
		}
	case decimal.Decimal:
		f, _ := value.Float64()
		return t.Convert(f)
	case decimal.NullDecimal:
		if value.Valid {
			f, _ := value.Decimal.Float64()
			return t.Convert(f)
		}
	case string:
		switch strings.ToLower(value) {
		case "on", "true":
			return int8(1), nil
		case "off", "false":
			return int8(0), nil
		}
	}

	return nil, sql.ErrInvalidSystemVariableValue.New(t.varName, v)
}

// MustConvert implements the Type interface.
func (t SystemBoolType_) MustConvert(v interface{}) interface{} {
	value, err := t.Convert(v)
	if err != nil {
		panic(err)
	}
	return value
}

// Equals implements the Type interface.
func (t SystemBoolType_) Equals(otherType sql.Type) bool {
	if ot, ok := otherType.(SystemBoolType_); ok {
		return t.varName == ot.varName
	}
	return false
}

// MaxTextResponseByteLength implements the Type interface
func (t SystemBoolType_) MaxTextResponseByteLength() uint32 {
	// system types are not sent directly across the wire
	return 0
}

// Promote implements the Type interface.
func (t SystemBoolType_) Promote() sql.Type {
	return t
}

// SQL implements Type interface.
func (t SystemBoolType_) SQL(ctx *sql.Context, dest []byte, v interface{}) (sqltypes.Value, error) {
	if v == nil {
		return sqltypes.NULL, nil
	}

	v, err := t.Convert(v)
	if err != nil {
		return sqltypes.Value{}, err
	}

	stop := len(dest)
	dest = strconv.AppendInt(dest, int64(v.(int8)), 10)
	val := dest[stop:]

	return sqltypes.MakeTrusted(t.Type(), val), nil
}

// String implements Type interface.
func (t SystemBoolType_) String() string {
	return "system_bool"
}

// Type implements Type interface.
func (t SystemBoolType_) Type() query.Type {
	return sqltypes.Int8
}

// CollationCoercibility implements sql.CollationCoercible interface.
func (SystemBoolType_) CollationCoercibility(ctx *sql.Context) (collation sql.CollationID, coercibility byte) {
	return sql.Collation_binary, 5
}

// ValueType implements Type interface.
func (t SystemBoolType_) ValueType() reflect.Type {
	return systemBoolValueType
}

// Zero implements Type interface.
func (t SystemBoolType_) Zero() interface{} {
	return int8(0)
}

// EncodeValue implements SystemVariableType interface.
func (t SystemBoolType_) EncodeValue(val interface{}) (string, error) {
	expectedVal, ok := val.(int8)
	if !ok {
		return "", sql.ErrSystemVariableCodeFail.New(val, t.String())
	}
	if expectedVal == 0 {
		return "0", nil
	}
	return "1", nil
}

// DecodeValue implements SystemVariableType interface.
func (t SystemBoolType_) DecodeValue(val string) (interface{}, error) {
	if val == "0" {
		return int8(0), nil
	} else if val == "1" {
		return int8(1), nil
	}
	return nil, sql.ErrSystemVariableCodeFail.New(val, t.String())
}
