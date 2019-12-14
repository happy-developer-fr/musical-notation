package pitch

type Pitch string

type Octave int8

func (o Octave) Print() string {
	oPrint := ""
	oPrintSymbol := "'"
	var oAbs = int(o)
	if oAbs < 0 {
		oPrintSymbol = ","
		oAbs = -oAbs
	}

	for i := 0; i <= oAbs; i++ {
		if i > 0 {
			oPrint += oPrintSymbol
		}
	}
	return oPrint
}

const (
	C Pitch = "C"
	D Pitch = "D"
	E Pitch = "E"
	F Pitch = "F"
	G Pitch = "G"
	A Pitch = "A"
	B Pitch = "B"
)
