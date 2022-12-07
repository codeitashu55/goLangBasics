package gettingStartedWithGo

func init() {
	println("calling variables")
}

func Variables() {

	//declaring variable
	var age int
	age = 32
	println("My age is", age)
	age = 42
	println("My age after 10 years", age)

	//declaring variables with initial value (-,-,-) -> where - represents an data type
	var roll int = 10
	println("Roll no is", roll, "for so long")

	//type inference -> variable can decide there type from the initial value
	var x = 24
	println(x)

	//multiple variable declaration
	var x1, x2 = 10, "what"
	println(x1, x2)

	//when we want to have some variables which have initial value and some do not have them
	var (
		x3 int
		x4 = 24
		x5 int
	)
	println(x3, x4, x5)

	//sort hand declaration :=  can be used to declare new variable and also for re-assignment of existing variable
	//limitation -> command must contain an new variable [using assigment operator solves that limitation]
	score, player := 10, "Ronaldo"
	score, player, year := 20, "Messi", 24
	score, player, year = 50, "Ganabri", 22
	println(player, "=", "Score =", score, "age", year)
}
