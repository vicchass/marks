package thestruct

type Address struct {
	City           string
	close_montreal bool
}

type Student struct {
	FirstName string
	LasttName string
	Marks     []int
	Average   float64
	Place     Address
}

var Stu1 = Student{FirstName: "victor", LasttName: "chassaigne", Marks: []int{12, 10, 11, 8}, Place: Address{City: "montreal", close_montreal: true}}

var Stu2 = Student{FirstName: "jon", LasttName: "time", Marks: []int{7, 12, 9, 4, 19}, Place: Address{City: "boston", close_montreal: false}}

var Stu3 = Student{FirstName: "amelie", LasttName: "pereira", Marks: []int{19, 6, 18, 10}, Place: Address{City: "laval", close_montreal: true}}
