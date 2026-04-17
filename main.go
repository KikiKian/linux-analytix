package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cpu        float64
	ram        float64
	download   string
	upload     string
	network    string
	width      int
	height     int
	prevRecv   uint64
	prevSent   uint64
	dlRaw      uint64
	ulRaw      uint64
	cpuHistory []float64
	dlHistory  []float64
	ulHistory  []float64
}

func (m model) Init() tea.Cmd { return tick() }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tickMsg:
		m.cpu = getCPU()
		m.ram = getRAM()
		var dl, ul uint64
		dl, ul, m.prevRecv, m.prevSent = getNetworkSpeed(m.prevRecv, m.prevSent)
		m.dlRaw = dl
		m.ulRaw = ul
		m.download = formatSpeed(dl)
		m.upload = formatSpeed(ul)
		m.cpuHistory = appendHistory(m.cpuHistory, m.cpu)
		m.dlHistory = appendHistory(m.dlHistory, float64(m.dlRaw/1024))
		m.ulHistory = appendHistory(m.ulHistory, float64(m.ulRaw/1024))
		return m, tick()
	}
	return m, nil
}

func (m model) View() string {
	graphWidth := m.width - 12
	if graphWidth < 10 {
		graphWidth = 10
	}
	return fmt.Sprintf(
		"CPU: %.1f%%\n%s\n\nDownload: %s\n%s\n\nUpload: %s\n%s\n\npress q to quit",
		m.cpu,
		renderCPUGraph(m.cpuHistory, graphWidth, 8),
		m.download,
		renderDownloadGraph(m.dlHistory, graphWidth, 6),
		m.upload,
		renderUploadGraph(m.ulHistory, graphWidth, 6),
	)
}

func main() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
