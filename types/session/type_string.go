// Code generated by "stringer -type=Type"; DO NOT EDIT.

package session

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Base-1]
	_ = x[Log-2]
	_ = x[Message-4]
	_ = x[Notice-8]
	_ = x[Request-16]
	_ = x[Group-32]
	_ = x[Private-64]
}

const (
	_Type_name_0 = "BaseLog"
	_Type_name_1 = "Message"
	_Type_name_2 = "Notice"
	_Type_name_3 = "Request"
	_Type_name_4 = "Group"
	_Type_name_5 = "Private"
)

var (
	_Type_index_0 = [...]uint8{0, 4, 7}
)

func (i Type) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _Type_name_0[_Type_index_0[i]:_Type_index_0[i+1]]
	case i == 4:
		return _Type_name_1
	case i == 8:
		return _Type_name_2
	case i == 16:
		return _Type_name_3
	case i == 32:
		return _Type_name_4
	case i == 64:
		return _Type_name_5
	default:
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}