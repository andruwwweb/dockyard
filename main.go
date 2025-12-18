package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)


type Container struct {
	Id      string
    Name    string
    Status  string
	Image   string
    Created int64  
}

func main() {
    p := tea.NewProgram(initialModel())
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error: %v", err)
    }
}