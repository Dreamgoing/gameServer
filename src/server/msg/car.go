package msg


///用户的数据,car
type Car struct {
	CarID int `json:"car_id"`
	X float32 `json:"x"`
	Y float32 `json:"y"`
	A float32 `json:"a"`
	V float32 `json:"v"`
}

/***@brief
目前设定汽车的简单运动,汽车的初始位置位于(0,0) 按照如下的坐标系,进行运动
y
^
|
|
|
|
| ---------->x
 */

///汽车的运动的函数
func (s *Car)Up()  {
	s.Y++
}

func (s *Car) Left() {
	s.X--
}

func (s *Car) Right()  {
	s.X++
}

func (s *Car) Down() {
	s.Y--
}