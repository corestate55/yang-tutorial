package turing_machine

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

// ToString convert Output to string
func (output *TuringMachine_TransitionFunction_Delta_Output) ToString(currentState uint32) string {
	var moveString string
	switch output.GetHeadMove() {
	case "left":
		moveString = "<= "
	case "right":
		moveString = " =>"
	default:
		moveString = " -:" // default (if empty)
	}

	stateString := fmt.Sprintf("[S%d]", output.GetState())
	if currentState > 0 && output.GetState() == 0 {
		stateString = "----" // keep current status (no need to change)
	}
	return fmt.Sprintf("%s %5s %4s", stateString, output.GetSymbol(), moveString)
}

// ToString convert TuringMachine to string
func (tm *TuringMachine) ToString(step uint32) string {
	var stateString string
	stateString = fmt.Sprintf("%4d  [S%d] | ", step, tm.GetState())
	for i, cell := range tm.GetTape().GetCell() {
		if i == int(tm.GetHeadPosition()) {
			stateString += fmt.Sprintf("<%s>|", cell.GetSymbol())
		} else {
			stateString += fmt.Sprintf(" %s |", cell.GetSymbol())
		}
	}
	return stateString
}

// ChangeState sets TuringMachine states
func (tm *TuringMachine) ChangeState(output *TuringMachine_TransitionFunction_Delta_Output) {
	// change to next state
	// Notice: protobuf-defined GetStaet() returns 0 when "state element is empty",
	//   same as initial state 0. But, it is different of "State 0" (existing state 0).
	//   This turing-machine assumes that its state starts 0 and DOES NOT back to 0 again.
	if output.GetState() > 0 {
		tm.State = output.GetState()
	}
	// write symbol to tape under head-position
	if output.GetSymbol() != "" {
		cellList := tm.GetTape().GetCell()
		cellList[tm.GetHeadPosition()].Symbol = output.GetSymbol()
	}
	// move to next head position
	switch output.GetHeadMove() {
	case "left":
		tm.HeadPosition--
	default:
		tm.HeadPosition++
	}
}

// ToXMLString converts TuringMachine to XML string
func (tm *TuringMachine) ToXMLString() string {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(tm, "", "  ")
	if err != nil {
		log.Fatalf("Error: Turing Machine XML Marshal err: %v\n", err)
	}
	return string(xmlBuf)
}

// ToJSONString converts TuringMachine to JSON string
func (tm *TuringMachine) ToJSONString() string {
	var jsonBuf, err = json.MarshalIndent(tm, "", "  ")
	if err != nil {
		log.Fatalf("Error: Turing Machine JSON Marshal err: %v\n", err)
	}
	return string(jsonBuf)
}

// ToXMLString convert notification notification to XML string
func (notification *Notification) ToXMLString() string {
	// marshal (returns []byte)
	var xmlBuf, err = xml.MarshalIndent(notification, "", "  ")
	if err != nil {
		log.Fatalf("Error: Notification XML Marshal err: %v\n", err)
	}
	return string(xmlBuf)
}

// ToJSONString convert notification notification to JSON string
func (notification *Notification) ToJSONString() string {
	var jsonBuf, err = json.MarshalIndent(notification, "", "  ")
	if err != nil {
		log.Fatalf("Error: Notification JSON Marshal err: %v\n", err)
	}
	return string(jsonBuf)
}
