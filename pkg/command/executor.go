package command

func ExecuteMultiCommands(commands []*Command) []*Result {
	var result []*Result
	for i := range commands {
		result[i] = commands[i].Execute()
	}
	return result
}
