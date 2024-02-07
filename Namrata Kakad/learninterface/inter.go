package main

import(
	"fmt"
)

type MotorVehicle interface{
	mileage() float64
}

type BMW struct{
	distance float64
	fuel float64
	averagespeed string 

}

type Audi struct {
	distance float64
	fuel float64
}


func (b BMW) mileage() float64{
	return b.distance/b.fuel
}

func(a Audi) mileage() float64{
	return a.distance/a.fuel
}

func totalMileage(m []MotorVehicle){
	tm :=0.0
	for _, v:= range m{
		tm = tm + v.mileage()
	}
	fmt.Printf("Total mileage %f \n", tm)
}

func main(){
	b1 := BMW{
		distance:167.9,
		fuel:35,
		averagespeed: "58",
	}

	a1 := Audi{
		distance:169.9,
		fuel:38,
	}

	person := [] MotorVehicle{b1, a1}

	totalMileage(person)

}