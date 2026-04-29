package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ckinan/sysmon/internal"
	"github.com/ckinan/sysmon/internal/collector"
)

// Model is the bubbletea model. It holds all UI state
type Model struct {
	snapCh <-chan collector.Snapshot // read-only channel from the collect
	ram    internal.Ram
	procs  []internal.Process
	height int // terminal height
}

func New(ch <-chan collector.Snapshot) Model {
	// height: 24 is a safe fallback
	// frame is painted right after startup, so this default is almost never actually visible
	return Model{snapCh: ch, height: 24}
}

func (m Model) Init() tea.Cmd {
	return waitForSnapshot(m.snapCh)
}
