// Code generated by "stringer -type=PriorityList"; DO NOT EDIT.

package bacnet

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ManualLifeSafety1-1]
	_ = x[ManualLifeSafety2-2]
	_ = x[Available3-3]
	_ = x[Available4-4]
	_ = x[CriticalEquipementControl5-5]
	_ = x[MinimumOnOff6-6]
	_ = x[Available7-7]
	_ = x[ManualOperator8-8]
	_ = x[Available9-9]
	_ = x[Available10-10]
	_ = x[Available11-11]
	_ = x[Available12-12]
	_ = x[Available13-13]
	_ = x[Available14-14]
	_ = x[Available15-15]
	_ = x[Available16-16]
}

const _PriorityList_name = "ManualLifeSafety1ManualLifeSafety2Available3Available4CriticalEquipementControl5MinimumOnOff6Available7ManualOperator8Available9Available10Available11Available12Available13Available14Available15Available16"

var _PriorityList_index = [...]uint8{0, 17, 34, 44, 54, 80, 93, 103, 118, 128, 139, 150, 161, 172, 183, 194, 205}

func (i PriorityList) String() string {
	i -= 1
	if i >= PriorityList(len(_PriorityList_index)-1) {
		return "PriorityList(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _PriorityList_name[_PriorityList_index[i]:_PriorityList_index[i+1]]
}
