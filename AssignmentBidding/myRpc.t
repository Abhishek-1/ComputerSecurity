MyRpc interface {
PrintAssignment () (list Assignmentlist)
BookAssignment (id uint64) (cnf uint64)

}

Assignmentlist struct {
	Field1 Assignment
	Field2 Assignment
	Field3 Assignment
	Field4 Assignment
	Field5 Assignment
	Field6 Assignment
	Field7 Assignment
	
}

Assignment struct {
	Field0 uint64
	Field1 string
	Field2 uint64
}

Strm struct {
	Field0 uint64
}
