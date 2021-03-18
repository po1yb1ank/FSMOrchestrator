package fsm

import (
	"github.com/looplab/fsm"
	"github.com/po1yb1ank/FSMOrchestrator/internal"
	"log"
	"strings"
)

var machine *fsm.FSM

func ProcessMachine(state internal.Machine) bool {
	event := eventBuilder(state)
	log.Println("got event:", event)
	err := machine.Event(event)
	if err != nil {
		log.Println("event", event, "is prohibited, returning false")
		return false
	}
	log.Println("event", event, "is ok, returning true")
	return true
}

func PushMachine(ch chan internal.Machine) bool {
	current := <-ch
	return ProcessMachine(current)
}

func InitMachine() {
	machine = fsm.NewFSM(
		"MANUAL;IDLE",
		fsm.Events{
			{Name: "AUTO;IDLE", Src: []string{"AUTO;STOP"}, Dst: "AUTO;IDLE"},
			{Name: "AUTO;IDLE", Src: []string{"MANUAL;IDLE"}, Dst: "AUTO;IDLE"},

			{Name: "AUTO;CONTINUE", Src: []string{"AUTO;PAUSE"}, Dst: "AUTO;CONTINUE"},

			{Name: "AUTO;PAUSE", Src: []string{"AUTO;IDLE"}, Dst: "AUTO;PAUSE"},
			{Name: "AUTO;PAUSE", Src: []string{"AUTO;CONTINUE"}, Dst: "AUTO;PAUSE"},

			{Name: "AUTO;STOP", Src: []string{"AUTO;IDLE"}, Dst: "AUTO;STOP"},
			{Name: "AUTO;STOP", Src: []string{"AUTO;CONTINUE"}, Dst: "AUTO;STOP"},
			{Name: "AUTO;STOP", Src: []string{"AUTO;PAUSE"}, Dst: "AUTO;STOP"},

			{Name: "MANUAL;IDLE", Src: []string{"AUTO;IDLE"}, Dst: "MANUAL;IDLE"},
			{Name: "MANUAL;IDLE", Src: []string{"AUTO;CONTINUE"}, Dst: "MANUAL;IDLE"},
			{Name: "MANUAL;IDLE", Src: []string{"AUTO;PAUSE"}, Dst: "MANUAL;IDLE"},
			{Name: "MANUAL;IDLE", Src: []string{"AUTO;STOP"}, Dst: "MANUAL;IDLE"},
			{Name: "MANUAL;IDLE", Src: []string{"MANUAL;IDLE"}, Dst: "MANUAL;IDLE"},
		},
		fsm.Callbacks{},
	)
}
func eventBuilder(state internal.Machine) string {
	return strings.Join([]string{strings.ToUpper(state.Mode), strings.ToUpper(state.Control)}, ";")
}
