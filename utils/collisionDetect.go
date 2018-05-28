package utils

func RectangeCollision(x1 float32, y1 float32, w1 float32, h1 float32, x2 float32, y2 float32, w2 float32, h2 float32) bool {
	if x1 >= x2+w2 || x1+w1 <= x2 || y1 >= y2+h2 || y1+h1 <= y2 {
		return false
	}
	return true
}
