package ui

import (
	"github.com/charmbracelet/bubbles/table"
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
	table  table.Model
}

func New(ch <-chan collector.Snapshot) Model {
	// height: 24 is a safe fallback
	// frame is painted right after startup, so this default is almost never actually visible
	cols := []table.Column{
		{Title: "PID", Width: 8},
		{Title: "NAME", Width: 25},
		{Title: "RSS", Width: 10},
	}
	t := table.New(
		table.WithColumns(cols),
		table.WithFocused(true), // focused = keyboard nav (↑/↓) is active
	)
	return Model{snapCh: ch, height: 24, table: t}
}

func (m Model) Init() tea.Cmd {
	return waitForSnapshot(m.snapCh)
}
