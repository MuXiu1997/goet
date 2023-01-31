package main

const projectName = "goet"

//goland:noinspection GoUnusedGlobalVariable
var (
	version = "dev"
	commit  string
	date    string
	builtBy string
)

func init() {
	initCmd()
	initConfig()
}

func main() {
	_ = cmd.Execute()
}
