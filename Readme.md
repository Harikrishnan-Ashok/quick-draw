# Quick Draw

Quick Draw is a simple command-line utility for system management, providing a terminal user interface for shutdown and reboot operations.

## Features

- Terminal-based user interface
- Options for system shutdown and reboot
- Easy-to-use arrow key navigation and selection

## Installation

### Prerequisites

- Go 1.16 or higher
- Git

### Building from source

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/quick-draw.git
   cd quick-draw
   ```

2. Build the binary:
   ```
   go build -o quick-draw
   ```

3. Move the binary to a directory in your PATH:
   ```
   sudo mv quick-draw /usr/local/bin/
   ```

## Usage

After installation, you can run the program by typing:

```
quick-draw
```

Use arrow keys to navigate and Enter to select an option.

### Integration with i3 window manager

To use Quick Draw with i3, add the following line to your i3 config file (`~/.config/i3/config`):

```
bindsym $mod+p exec --no-startup-id alacritty -e quick-draw
```

Replace `alacritty` with your preferred terminal emulator if different.


## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
