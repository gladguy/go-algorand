// Code generated by "stringer -type=stateMachineTag"; DO NOT EDIT.

package agreement

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[demultiplexer-0]
	_ = x[playerMachine-1]
	_ = x[voteMachine-2]
	_ = x[voteMachineRound-3]
	_ = x[voteMachinePeriod-4]
	_ = x[voteMachineStep-5]
	_ = x[proposalMachine-6]
	_ = x[proposalMachineRound-7]
	_ = x[proposalMachinePeriod-8]
}

const _stateMachineTag_name = "demultiplexerplayerMachinevoteMachinevoteMachineRoundvoteMachinePeriodvoteMachineStepproposalMachineproposalMachineRoundproposalMachinePeriod"

var _stateMachineTag_index = [...]uint8{0, 13, 26, 37, 53, 70, 85, 100, 120, 141}

func (i stateMachineTag) String() string {
	if i < 0 || i >= stateMachineTag(len(_stateMachineTag_index)-1) {
		return "stateMachineTag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _stateMachineTag_name[_stateMachineTag_index[i]:_stateMachineTag_index[i+1]]
}
