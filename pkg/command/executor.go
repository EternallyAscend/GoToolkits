package command

func ExecuteMultiCommands(commands []*Command) []*Result {
	var result []*Result
	for i := range commands {
		result = append(result, commands[i].Execute())
	}
	return result
}
