package main

var global_number_of_threads = 8

func SetGlobalNumberOfThreads(number_of_threads int) {
	global_number_of_threads = number_of_threads
}

func GetGlobalNumberOfThreads() int {
	return global_number_of_threads
}
