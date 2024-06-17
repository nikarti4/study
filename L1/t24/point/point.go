package point

// структура точки с координатами
type Point struct {
	x float64
	y float64
}

// конструктор точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// получить абсциссу точки
func (p *Point) GetX() float64 {
	return p.x
}

// получить ординату точки
func (p *Point) GetY() float64 {
	return p.y
}

// задать абсциссу точки
func (p *Point) SetX(x float64) {
	p.x = x
}

// задать ординату точки
func (p *Point) SetY(y float64) {
	p.y = y
}
