package main

func main() {
	a := App{}

	a.Initialize("root", "", "db_golang")

	a.Run(":8080")
}
