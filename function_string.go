// Code generated by "stringer -type=Function"; DO NOT EDIT.

package bacnet

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BacFuncResult-0]
	_ = x[BacFuncWriteBroadcastDistributionTable-1]
	_ = x[BacFuncBroadcastDistributionTable-2]
	_ = x[BacFuncBroadcastDistributionTableAck-3]
	_ = x[BacFuncForwardedNPDU-4]
	_ = x[BacFuncUnicast-10]
	_ = x[BacFuncBroadcast-11]
}

const (
	_Function_name_0 = "BacFuncResultBacFuncWriteBroadcastDistributionTableBacFuncBroadcastDistributionTableBacFuncBroadcastDistributionTableAckBacFuncForwardedNPDU"
	_Function_name_1 = "BacFuncUnicastBacFuncBroadcast"
)

var (
	_Function_index_0 = [...]uint8{0, 13, 51, 84, 120, 140}
	_Function_index_1 = [...]uint8{0, 14, 30}
)

func (i Function) String() string {
	switch {
	case i <= 4:
		return _Function_name_0[_Function_index_0[i]:_Function_index_0[i+1]]
	case 10 <= i && i <= 11:
		i -= 10
		return _Function_name_1[_Function_index_1[i]:_Function_index_1[i+1]]
	default:
		return "Function(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
