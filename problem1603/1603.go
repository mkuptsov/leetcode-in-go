package problem1603

type ParkingSystem struct {
	Spaces [4]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Spaces: [4]int{0, big, medium, small},
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.Spaces[carType] > 0 {
		this.Spaces[carType]--
		return true
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
