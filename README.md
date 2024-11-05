# Snake Game in Go

A classic Snake game implementation using Go and the Fyne toolkit. The game features a modern, clean interface with multiple food items for enhanced gameplay.

## Features

- Clean, minimalist graphics
- 20 food items simultaneously on the field
- Score tracking
- Collision detection (walls and self)
- Automatic restart after game over
- Smooth controls using arrow keys

## Requirements

- Go 1.16 or higher
- Fyne v2.x

## Installation

```bash
git clone https://github.com/yourusername/snake-game.git
cd snake-game
go mod tidy
go run .
```

## Controls

- ↑ Up Arrow: Move up
- ↓ Down Arrow: Move down
- ← Left Arrow: Move left
- → Right Arrow: Move right

## Game Rules

- Control the snake to eat the red food items
- Each food item eaten increases your score by 1
- The snake grows longer with each food item eaten
- Game ends if the snake hits the wall or itself
- The game automatically restarts 2 seconds after game over

## Technical Details

- Written in Go
- Uses Fyne GUI toolkit
- Grid size: 20x20
- Refresh rate: 200ms

## Contributing

Feel free to fork the repository and submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)