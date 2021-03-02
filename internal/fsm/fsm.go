package fsm

import (
	"github.com/looplab/fsm"
	"github.com/po1yb1ank/FSMOrchestrator/internal/rest/endpoint"
	"strings"
)
var machine *fsm.FSM
func ProcessMachine(state endpoint.Machine) (result bool){
	event := eventBuilder(state)
	err := machine.Event(event)
	if err != nil {
		return false
	}
	return true
}

func PushMachine(ch chan endpoint.Machine){
	current := <- ch
	ProcessMachine(current)
}

func InitMachine(){
	machine = fsm.NewFSM(
		"MANUAL;IDLE",
		fsm.Events{
			{Name: "AUTO;IDLE", Src: []string{"AUTO;STOP"}, Dst: "AUTO;STOP"},
			{Name: "AUTO;IDLE", Src: []string{"MANUAL;IDLE"}, Dst: "MANUAL;IDLE"},

			{Name: "AUTO;CONTINUE", Src: []string{"AUTO;PAUSE"}, Dst: "AUTO;PAUSE"},

			{Name: "AUTO;PAUSE", Src: []string{"AUTO;IDLE"}, Dst: "AUTO;IDLE"},
			{Name: "AUTO;PAUSE", Src: []string{"AUTO;CONTINUE"}, Dst: "AUTO;CONTINUE"},

			{Name: "AUTO;STOP", Src: []string{"AUTO;IDLE"}, Dst: "AUTO;IDLE"},
			{Name: "AUTO;STOP", Src: []string{"AUTO;CONTINUE"}, Dst: "AUTO;CONTINUE"},
			{Name: "AUTO;STOP", Src: []string{"AUTO;PAUSE"}, Dst: "AUTO;PAUSE"},

			{Name: "MANUAL;IDLE", Src: []string{"AUTO;IDLE"}, Dst: "AUTO;IDLE"},
			{Name: "MANUAL;IDLE", Src: []string{"AUTO;CONTINUE"}, Dst: "AUTO;CONTINUE"},
			{Name: "MANUAL;IDLE", Src: []string{"AUTO;PAUSE"}, Dst: "AUTO;PAUSE"},
			{Name: "MANUAL;IDLE", Src: []string{"AUTO;STOP"}, Dst: "AUTO;STOP"},
			{Name: "MANUAL;IDLE", Src: []string{"MANUAL;IDLE"}, Dst: "MANUAL;IDLE"},

		},
		fsm.Callbacks{},
	)
}
func eventBuilder(state endpoint.Machine)string{
	return strings.Join([]string{strings.ToUpper(state.Mode), strings.ToUpper(state.Control)}, ";")
}
