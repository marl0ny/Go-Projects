package main

var global_number_of_routines = 8

func SetGlobalNumberOfRoutines(n int) {
	global_number_of_routines = n
}

func GetGlobalNumberOfRoutines() int {
	return global_number_of_routines
}
