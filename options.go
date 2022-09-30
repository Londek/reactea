package reactea

import (
	tea "github.com/charmbracelet/bubbletea"
)

type nilReader struct{}

func (nilReader) Read([]byte) (int, error) {
	return 0, nil
}

// Useful for testing on Github Actions, by default Bubble Tea
// would try reading from /dev/tty, but on Github Actions
// it's restricted resulting in error
func WithoutInput() func(*tea.Program) {
	return tea.WithInput(nilReader{})
}

func WithRoute(route string) func(*tea.Program) {
	return func(*tea.Program) {
		currentRoute = route
	}
}
