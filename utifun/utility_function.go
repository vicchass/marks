package utifun

import (
	"fmt"

	st "github.com/vicchass/marks/theStruct"
)

// add a single mark to add student
func Add_single_mark(the_student st.Student) st.Student {
	var new_mark int
	fmt.Println("type the new mark for ", the_student.FirstName)
	fmt.Scan(&new_mark)
	the_student.Marks = append(the_student.Marks, new_mark)
	return the_student

}
