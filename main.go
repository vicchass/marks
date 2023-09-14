package main

import (
	"fmt"

	st "github.com/vicchass/marks/theStruct"
	"github.com/vicchass/tools/utisli"
)

// VARIABLE
var all_students []st.Student

func init() {
	//assign all students to a slice of students
	all_students = append(all_students, st.Stu1, st.Stu2, st.Stu3)

}

func main() {

	//START LOGIC

	all_students = Get_average(all_students)
	Print_stu(all_students)

	// END LOGIC
}

// print name and average
func Print_stu(sli_stu []st.Student) {
	for _, v := range sli_stu {
		fmt.Println("the average of", v.FirstName, "is ", v.Average)
	}
}

// return the students with theire average
func Get_average(sli_stu []st.Student) []st.Student {
	for k, v := range sli_stu {
		new := utisli.Average_slice_int(v.Marks)
		sli_stu[k].Average = new
	}

	fmt.Println(sli_stu[0].Average)
	return sli_stu

}

// add marks to a student
func Add_mark(sli_stu []st.Student) []st.Student {

}
