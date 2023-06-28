// Code generated by "stringer -type=ErrCode -linecomment"; DO NOT EDIT.

package berr

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrorInvalidArgument-1000]
	_ = x[ErrorNotFound-1001]
	_ = x[ErrorPermissionDeny-1002]
	_ = x[ErrorWriteOnClose-1003]
}

const _ErrCode_name = "error argument invaliderror not founderror permission denyerror write on close"

var _ErrCode_index = [...]uint8{0, 22, 37, 58, 78}

func (i ErrCode) String() string {
	i -= 1000
	if i < 0 || i >= ErrCode(len(_ErrCode_index)-1) {
		return "ErrCode(" + strconv.FormatInt(int64(i+1000), 10) + ")"
	}
	return _ErrCode_name[_ErrCode_index[i]:_ErrCode_index[i+1]]
}