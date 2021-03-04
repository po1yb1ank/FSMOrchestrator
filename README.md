# FSMOrchestrator
  Listens on :8080.
  
  Parses json input on localhost:8080/request and sends it to remote only when the input is valid state in the FSM.
## REST:
  - /request - json structure in internal/rest/endpoint/types.go
