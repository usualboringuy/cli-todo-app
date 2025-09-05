package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	todos.add("Complete Math Assignment")
	todos.add("Revise Concepts of DBMS")
	todos.toggle(1)
	todos.print()
	storage.Sava(todos)

}
