package main



func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos);
	CmdFlags := NewCmdFlags();
	CmdFlags.Execute(&todos)
	storage.Save(todos);
}