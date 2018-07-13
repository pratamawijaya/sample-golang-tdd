package main

func main() {
	a := App{}

	a.Initialize("db_username", "db_password", "db_name")

	a.Run(":8080")
}
