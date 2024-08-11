package main

import "fmt"

func main() {
	// boolean operation = && (AND), || (OR), ! (Reverse)

	// && - IF BOTH TRUE = TRUE
	// || - IF ANY TRUE = TRUE

	var subjectScore = 90
	var attendanceScore = 80
	fmt.Println("Subject Score: ", subjectScore)
	fmt.Println("Attendance Score: ", attendanceScore)

	var subjectPassed bool = subjectScore > 80
	fmt.Println("Subject Passed: Subject Score > 80 = ", subjectPassed)

	var attendancePassed bool = attendanceScore > 80
	fmt.Println("Attendance Passed: Attendance Score > 80 = ", attendancePassed)

	var passed bool = subjectPassed && attendancePassed
	fmt.Println("Student Passed: Subject Passed && Attendance Passed ?")
	fmt.Println("Student Passed: ", passed)
}
