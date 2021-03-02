package fsm

import (
	"github.com/looplab/fsm"
	"github.com/po1yb1ank/FSMOrchestrator/internal/rest/endpoint"
	"log"
	"strings"
)

var machine *fsm.FSM

func ProcessMachine(state endpoint.Machine) bool {
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

func PushMachine(ch chan endpoint.Machine) bool {
	current := <-ch
	return ProcessMachine(current)
}

func InitMachine() {
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
func eventBuilder(state endpoint.Machine) string {
	return strings.Join([]string{strings.ToUpper(state.Mode), strings.ToUpper(state.Control)}, ";")
}
