package model

type commandResult struct {
	message string
	err     bool
}

func handleCommand(m Model, cmd string) (Model, commandResult) {
	switch cmd {
	case "clear":
		m.CpuHistory = []float64{}
		m.DlHistory = []float64{}
		m.UlHistory = []float64{}
		return m, commandResult{"graphs cleared", false}
	case "quit", "q":
		return m, commandResult{"", false}
	case "help":
		return m, commandResult{"commands: clear, quit, help", false}
	default:
		return m, commandResult{"unknown command: " + cmd, true}
	}
}
