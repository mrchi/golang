package lengthconv

import "fmt"

type Metre float64
type Inch float64

func (m Metre) String() string { return fmt.Sprintf("%gm", m) }
func (i Inch) String() string  { return fmt.Sprintf("%ginch", i) }

func MToI(m Metre) Inch { return Inch(m / 0.0254) }
func IToM(i Inch) Metre { return Metre(i * 0.0254) }
