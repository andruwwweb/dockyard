package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
    containers []Container
    cursor     int
    selected   map[int]bool
    err        error
}

func initialModel() model {
    containers, err := getRealContainers()
    return model{
        containers: containers,
        cursor:     0,
        selected:   make(map[int]bool),
        err:        err,
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
            
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
            
        case "down", "j":
            if m.cursor < len(m.containers)-1 {
                m.cursor++
            }
            
        case " ":
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = true
            }
            
        case "enter":
            container := m.containers[m.cursor]
            if container.Status == "running" {
                stopContainer(container.Id)
            } else {
                startContainer(container.Id)
            }
            m.containers, _ = getRealContainers()
		
		case "s":
			if len(m.selected) > 0 {
				for i := range m.selected {
					if m.containers[i].Status != "running" {
						startContainer(m.containers[i].Id)
					}
				}
				m.selected = make(map[int]bool)
				m.containers, _ = getRealContainers()
			}
			
		case "x":
			if len(m.selected) > 0 {
				for i := range m.selected {
					if m.containers[i].Status == "running" {
						stopContainer(m.containers[i].Id)
					}
				}
				m.selected = make(map[int]bool)
				m.containers, _ = getRealContainers()
			}
			
		case "d":
			if len(m.selected) > 0 {
				for i := range m.selected {
					removeContainer(m.containers[i].Id)
				}
				m.selected = make(map[int]bool)
				m.containers, _ = getRealContainers()
				if m.cursor >= len(m.containers) {
					m.cursor = len(m.containers) - 1
				}
				if m.cursor < 0 {
					m.cursor = 0
				}
			}
        case "r":
            m.containers, m.err = getRealContainers()
            if m.err == nil {
                if m.cursor >= len(m.containers) {
                    m.cursor = len(m.containers) - 1
                }
                if m.cursor < 0 {
                    m.cursor = 0
                }
            }
		
		case "c":
			m.selected = make(map[int]bool)
		
		case "a":
			for i := range m.containers {
				m.selected[i] = true
			}
        }

    }
    return m, nil
}

func (m model) View() string {
    if m.err != nil {
        s := "❌ Error connecting to Docker\n\n"
        s += fmt.Sprintf("Error: %s\n\n", m.err)
        s += "Please make sure:\n"
        s += "  • Docker daemon is running\n"
        s += "  • You have permissions to access Docker\n"
        s += "  • Docker socket is available\n\n"
        s += "Press 'r' to retry or 'q' to quit\n"
        return s
    }

    s := "🐳 Dockyard - Docker Container Manager\n"
    s += "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n\n"
    
    // Статистика
    running := 0
    stopped := 0
    for _, c := range m.containers {
        if c.Status == "running" {
            running++
        } else {
            stopped++
        }
    }
    
    s += fmt.Sprintf("Total: %d | 🟢 Running: %d | 🔴 Stopped: %d\n\n", 
        len(m.containers), running, stopped)
    
    // Список контейнеров
    for i, container := range m.containers {
        cursor := "  "
        if m.cursor == i {
            cursor = "→ "
        }
        
        checkbox := "[ ]"
        if m.selected[i] {
            checkbox = "[✓]"
        }
        
        status := "🔴"
        if container.Status == "running" {
            status = "🟢"
        }
        
	s += fmt.Sprintf("%s%s %s %-35s %-35s %s\n", 
		cursor,
		checkbox,
		status,
		container.Name,
		container.Id,
		formatUptime(container.Created)) 
    }
    
    s += "\n"
    s += "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n"
    
	if len(m.selected) > 0 {
		s += fmt.Sprintf("📦 %d container(s) selected\n\n", len(m.selected))
		s += "BULK ACTIONS:\n"
		s += "  s - Start all selected\n"
		s += "  x - Stop all selected\n"
		s += "  d - Delete all selected\n"
		s += "  c - Clear selection\n\n"
	} else {
		s += "NAVIGATION:\n"
		s += "  ↑/k - Move up       ↓/j - Move down\n\n"
		s += "ACTIONS:\n"
		s += "  Space - Select/deselect    a - Select all\n"
		s += "  Enter - Start/Stop         r - Refresh       q - Quit\n\n"
	}
    
    return s
}