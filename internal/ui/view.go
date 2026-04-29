package ui

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"text/tabwriter"

	"github.com/ckinan/sysmon/internal"
)

func (m Model) View() string {
	header := fmt.Sprintf("RAM used: %s / %s\n\n", internal.HumanBytes(m.ram.MemUsed), internal.HumanBytes(m.ram.MemTotal))
	var buf strings.Builder
	w := tabwriter.NewWriter(&buf, 0, 0, 2, ' ', 0) // minwidth=0, tabwidth=0, padding=2
	fmt.Fprintln(w, "PID\tNAME\tRSS")
	fmt.Fprintln(w, "---\t----\t---")

	sorted := slices.Clone(m.procs)
	slices.SortFunc(sorted, func(a, b internal.Process) int {
		return cmp.Compare(b.RssKB, a.RssKB)
	})

	// Reserve lines for non-table content:
	//   Line #1: RAM header
	//   Line #2: Blank line after header
	//   Line #3: Column header (PID  NAME  RSS)
	//   Line #4: Separator row (---  ----  ---)
	//   Line #5: Blank line after table content
	//   Line #6: Footer ([q] quit ...)
	const overhead = 6

	// Step 1: how many rows fit on screen without pushing the footer off
	limit := max(1, m.height-overhead) // max(1, ...) ensures we always show at least 1 row
	// Step 2: don't try to slice more rows than we actually have
	limit = min(limit, len(sorted))
	for _, p := range sorted[:limit] {
		fmt.Fprintf(w, "%d\t%s\t%s\n", p.Pid, p.Name, internal.HumanBytes(p.RssKB))
	}

	w.Flush() // must call Flush because tabwriter only writes to buf after seeing all rows

	// Pin footer to the bottom of the terminal regardless of how many rows are shown
	// If only a few processes are visible, blank lines fill the gap so the footer
	// always appears on the last line of the screen
	blanks := m.height - overhead - limit
	for range max(0, blanks) {
		buf.WriteString("\n")
	}
	footer := "\n[↑/↓] scroll   [q] quit"
	return header + buf.String() + footer
}
