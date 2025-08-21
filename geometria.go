package geometria

type Triangulo struct {
	Base   float64
	Altura float64
}

type Cuadrado struct {
	Lado float64
}

type Pentagono struct {
	Lado    float64
	Apotema float64
}

func (t *Triangulo) Area() float64 {
	return t.Base * t.Altura / 2
}

func (c *Cuadrado) Area() float64 {
	return c.Lado * c.Lado
}

func (p *Pentagono) Area() float64 {
	peri := p.perimetro()
	return peri * p.Apotema / 2
}

func (p *Pentagono) perimetro() float64 {
	return p.Lado * 5
}

func Figura(lados string) string {
	switch lados {
	case "3":
		return "Triangulo"
	case "4":
		return "Cuadrado"
	case "5":
		return "Pentagono"
	default:
		return "No valido"
	}
}

// func main() {
// 	pentagono := Pentagono{Lado: 2, Apotema: 5}
// 	area := pentagono.Area()
// 	fmt.Println("El area es", area)
// 	fmt.Println("Figura", Figura("5"))
// }
