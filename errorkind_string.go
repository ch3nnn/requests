// Code generated by "stringer -type=ErrorKind"; DO NOT EDIT.

package requests

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[KindNone-0]
	_ = x[KindUnknown-1]
	_ = x[KindURLErr-2]
	_ = x[KindBodyGet-3]
	_ = x[KindBadMethod-4]
	_ = x[KindNilContext-5]
	_ = x[KindConnectErr-6]
	_ = x[KindInvalid-7]
	_ = x[KindHandlerErr-8]
}

const _ErrorKind_name = "KindNoneKindUnknownKindURLErrKindBodyGetKindBadMethodKindNilContextKindConnectErrKindInvalidKindHandlerErr"

var _ErrorKind_index = [...]uint8{0, 8, 19, 29, 40, 53, 67, 81, 92, 106}

func (i ErrorKind) String() string {
	if i < 0 || i >= ErrorKind(len(_ErrorKind_index)-1) {
		return "ErrorKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorKind_name[_ErrorKind_index[i]:_ErrorKind_index[i+1]]
}