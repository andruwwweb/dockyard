# 🐳 Dockyard

A blazing fast, keyboard-driven Terminal UI for managing Docker containers. Built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Docker](https://img.shields.io/badge/Docker-Required-2496ED?style=flat&logo=docker)

## ✨ Features

- 🚀 **Blazing Fast** - Written in Go, single binary with zero dependencies
- ⌨️ **Keyboard-Driven** - Vim-style navigation (j/k) and intuitive shortcuts
- 📦 **Bulk Operations** - Select multiple containers and manage them at once
- 🎯 **Smart Actions** - Start, stop, restart, and delete containers effortlessly
- 📊 **Real-time Info** - Container status, IDs, and creation times at a glance
- 🎨 **Beautiful Interface** - Clean, minimal TUI that works everywhere
- 🔄 **Auto-Refresh** - Keep your view up-to-date with one keystroke

## 🎬 Demo

```
🐳 Dockyard - Docker Container Manager
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Total: 3 | 🟢 Running: 2 | 🔴 Stopped: 1

      NAME                                 CONTAINER ID    CREATED
────────────────────────────────────────────────────────────────────
→  [✓] 🟢 app-redis                         5061f3af6d33    17d 11h
   [✓] 🟢 app-postgres                      98636666137b    12d 18h
   [ ] 🔴 app-nginx                         3aafe2acaf56    4h 57m

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
📦 2 container(s) selected

BULK ACTIONS:
  s - Start all selected
  x - Stop all selected
  d - Delete all selected
  c - Clear selection
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21 or higher
- Docker daemon running

### Installation

**Option 1: Install from source**

```bash
git clone https://github.com/YOUR_USERNAME/dockyard.git
cd dockyard
go build -o dockyard
./dockyard
```

**Option 2: Install with go install**

```bash
go install github.com/YOUR_USERNAME/dockyard@latest
dockyard
```

**Option 3: Download binary** (coming soon)

## ⌨️ Keybindings

### Navigation

| Key        | Action           |
| ---------- | ---------------- |
| `↑` or `k` | Move cursor up   |
| `↓` or `j` | Move cursor down |

### Container Actions

| Key     | Action                            |
| ------- | --------------------------------- |
| `Space` | Select/deselect container         |
| `a`     | Select all containers             |
| `c`     | Clear all selections              |
| `Enter` | Start/Stop container under cursor |
| `r`     | Refresh container list            |

### Bulk Operations (when containers are selected)

| Key | Action                         |
| --- | ------------------------------ |
| `s` | Start all selected containers  |
| `x` | Stop all selected containers   |
| `d` | Delete all selected containers |

### General

| Key             | Action           |
| --------------- | ---------------- |
| `q` or `Ctrl+C` | Quit application |

## 🎯 Usage Examples

### Start a stopped container

1. Navigate to the container with `↑/↓` or `j/k`
2. Press `Enter` to start it

### Bulk stop multiple running containers

1. Select containers with `Space`
2. Press `x` to stop all selected

### Clean up stopped containers

1. Press `a` to select all
2. Navigate through and deselect running ones with `Space`
3. Press `d` to delete all selected stopped containers

## 🏗️ Built With

- [Go](https://golang.org/) - Fast, compiled language
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Powerful TUI framework based on The Elm Architecture
- [Docker Engine API](https://github.com/docker/docker) - Official Docker Go SDK

## 🛠️ Development

### Setup

```bash
# Clone the repository
git clone https://github.com/YOUR_USERNAME/dockyard.git
cd dockyard

# Download dependencies
go mod download

# Run in development mode
go run *.go
```

### Project Structure

```
dockyard/
├── main.go       # Entry point
├── docker.go     # Docker API interactions
├── tui.go        # Terminal UI logic
├── actions.go    # Container operations (start/stop/delete)
└── go.mod        # Dependencies
```

### Building

```bash
# Build for current platform
go build -o dockyard

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o dockyard-linux-amd64
GOOS=darwin GOARCH=amd64 go build -o dockyard-darwin-amd64
GOOS=windows GOARCH=amd64 go build -o dockyard-windows-amd64.exe
```

## 🗺️ Roadmap

- [ ] Container logs viewer
- [ ] Real-time CPU/Memory/Network stats
- [ ] Docker Compose stack management
- [ ] Detailed container inspection view
- [ ] Search and filter containers
- [ ] Custom color themes
- [ ] Configuration file support
- [ ] Container restart action
- [ ] Image management
- [ ] Volume management

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is [MIT](LICENSE) licensed.

## 🙏 Acknowledgments

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Amazing TUI framework
- [lazydocker](https://github.com/jesseduffield/lazydocker) - Inspiration for Docker TUI
- Docker community for the excellent Go SDK

## 👨‍💻 Author

**Andrei Monchenko**

## ⭐ Show your support

Give a ⭐️ if you find this project useful!

---

<p align="center">
  <i>Thanks for paying attention on</i>
</p>
