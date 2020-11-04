package solver

//Solver func for all problems
type Solver func(chan string, chan string)

//Problem func
type Problem struct {
	Key       string
	Day       string
	Part      string
	Solve     Solver
	InputFile string
}
