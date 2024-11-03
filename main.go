package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image"
	"image/color"
	"math/rand"
	"time"
)

const (
	gridSize = 20
	cellSize = 20
)

type SnakeGame struct {
	snake     []fyne.Position
	direction fyne.Position
	food      []fyne.Position
	score     int
	gameOver  bool
	canvas    *canvas.Raster
}

func main() {
	rand.Seed(time.Now().UnixNano())

	myApp := app.New()
	myWindow := myApp.NewWindow("Snake Game")
	myWindow.Resize(fyne.NewSize(gridSize*cellSize, gridSize*cellSize))

	game := &SnakeGame{
		snake:     []fyne.Position{{X: 5, Y: 5}},
		direction: fyne.Position{X: 1, Y: 0},
		food:      generateFood(20),
		score:     0,
		gameOver:  false,
	}

	game.canvas = canvas.NewRaster(game.drawGame)
	game.canvas.Resize(fyne.NewSize(float32(gridSize*cellSize), float32(gridSize*cellSize)))
	game.canvas.SetMinSize(fyne.NewSize(float32(gridSize*cellSize), float32(gridSize*cellSize)))

	content := container.New(layout.NewVBoxLayout())
	scoreLabel := widget.NewLabel("Score: 0")
	content.Add(scoreLabel)
	content.Add(game.canvas)

	myWindow.SetContent(content)

	// Tastatureingaben
	myWindow.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
		switch ev.Name {
		case fyne.KeyUp:
			if game.direction.Y != 1 {
				game.direction = fyne.Position{X: 0, Y: -1}
			}
		case fyne.KeyDown:
			if game.direction.Y != -1 {
				game.direction = fyne.Position{X: 0, Y: 1}
			}
		case fyne.KeyLeft:
			if game.direction.X != 1 {
				game.direction = fyne.Position{X: -1, Y: 0}
			}
		case fyne.KeyRight:
			if game.direction.X != -1 {
				game.direction = fyne.Position{X: 1, Y: 0}
			}
		}
	})

	// Spiel-Loop starten
	go func() {
		game.run(scoreLabel)
	}()

	myWindow.ShowAndRun()
}

func (g *SnakeGame) run(scoreLabel *widget.Label) {
	ticker := time.NewTicker(200 * time.Millisecond)
	for range ticker.C {
		if g.gameOver {
			scoreLabel.SetText("Game Over! Score: " + string(rune(g.score+'0')))
			// Optional: Spiel nach kurzer Verzögerung neu starten
			time.Sleep(2 * time.Second)
			g.reset()
			scoreLabel.SetText("Score: 0")
		}
		g.update()
		scoreLabel.SetText("Score: " + string(rune(g.score+'0')))
		g.canvas.Refresh()
	}
}

// Neue Methode zum Zurücksetzen des Spiels
func (g *SnakeGame) reset() {
	g.snake = []fyne.Position{{X: 5, Y: 5}}
	g.direction = fyne.Position{X: 1, Y: 0}
	g.food = generateFood(20)
	g.score = 0
	g.gameOver = false
}

func (g *SnakeGame) update() {
	head := g.snake[0]
	newHead := fyne.Position{
		X: head.X + g.direction.X,
		Y: head.Y + g.direction.Y,
	}

	// Kollisionserkennung mit Wänden
	if newHead.X < 0 || newHead.X >= float32(gridSize) ||
		newHead.Y < 0 || newHead.Y >= float32(gridSize) {
		g.gameOver = true
		return
	}

	// Kollisionserkennung mit sich selbst (nur wenn Schlange länger als 1 ist)
	if len(g.snake) > 1 {
		for _, segment := range g.snake {
			if newHead == segment {
				g.gameOver = true
				return
			}
		}
	}

	// Essen aufsammeln
	foodEaten := false
	for i, f := range g.food {
		if newHead == f {
			g.snake = append([]fyne.Position{newHead}, g.snake...)
			g.score++
			// Entferne das gegessene Futter und generiere ein neues
			g.food = append(g.food[:i], g.food[i+1:]...)
			g.food = append(g.food, generateFood(1)...)
			foodEaten = true
			break
		}
	}

	if !foodEaten {
		g.snake = append([]fyne.Position{newHead}, g.snake[:len(g.snake)-1]...)
	}
}

func generateFood(count int) []fyne.Position {
	food := make([]fyne.Position, 0, count)
	occupied := make(map[fyne.Position]bool)

	// Füge die aktuelle Schlangenposition zur occupied-Map hinzu
	for i := 0; i < count; i++ {
		var newFood fyne.Position
		for {
			newFood = fyne.Position{
				X: float32(rand.Intn(gridSize)),
				Y: float32(rand.Intn(gridSize)),
			}
			if !occupied[newFood] {
				occupied[newFood] = true
				food = append(food, newFood)
				break
			}
		}
	}
	return food
}

func (g *SnakeGame) drawGame(w, h int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))

	// Hintergrund in hellgrau zeichnen
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			img.Set(x, y, color.RGBA{200, 200, 200, 255})
		}
	}

	cellWidth := float32(w) / float32(gridSize)
	cellHeight := float32(h) / float32(gridSize)

	// Schlange zeichnen
	for _, pos := range g.snake {
		x := int(pos.X * cellWidth)
		y := int(pos.Y * cellHeight)
		for i := 1; i < int(cellWidth)-1; i++ {
			for j := 1; j < int(cellHeight)-1; j++ {
				img.Set(x+i, y+j, color.RGBA{0, 200, 0, 255})
			}
		}
	}

	// Alle Futtersteine zeichnen
	for _, food := range g.food {
		foodX := int(food.X * cellWidth)
		foodY := int(food.Y * cellHeight)
		for i := 1; i < int(cellWidth)-1; i++ {
			for j := 1; j < int(cellHeight)-1; j++ {
				img.Set(foodX+i, foodY+j, color.RGBA{200, 0, 0, 255})
			}
		}
	}

	return img
}
