package endpoint

type Machine struct {
	Mode    string `json:"mode"`
	Offset  Offset `json:"offset"`
	Control string `json:"control"`
}

type Offset struct {
	X string `json:"x"`
	Y string `json:"y"`
	Z string `json:"z"`
}

var RemotePath string
