package pitch

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
