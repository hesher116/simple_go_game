package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	totalView       = 20                         // Загальна кількість рядків дороги, які гравець може бачити
	viewBehind      = 2                          // Кількість рядків дороги за машиною
	viewAhead       = totalView - viewBehind - 1 // Кількість рядків дороги перед машиною
	roadWidth       = 50                         // Ширина дороги
	carPosition     = roadWidth / 2
	road            []string
	obstacleChance  = 30
	tickDuration    = 2 * time.Second
	distance        = 0
	difficulties    = []string{"Very Easy", "Easy", "Medium", "Hard", "Very Hard", "Extreme", "Impossible"}
	difficultyIndex = 2
	ticker          *time.Ticker
)

func initializeRoad() {
	road = make([]string, totalView)
	for i := range road {
		if i == 0 {
			road[i] = strings.Repeat("#", roadWidth) // Upper boundary
		} else if i == 1 {
			road[i] = "|" + strings.Repeat(" ", (roadWidth-6)/2) + "START" + strings.Repeat(" ", ((roadWidth-6)/2)-1) + "|" // Start line
		} else {
			road[i] = "|" + strings.Repeat(" ", roadWidth-2) + "|"
		}
	}
}

func display() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for i, line := range road {
		tbprint(0, i, termbox.ColorDefault, termbox.ColorDefault, line)
	}
	// Draw the car as a larger block
	car := []string{"╔══╗", "║██║", "╚══╝"}
	carY := totalView / 2
	for i, part := range car {
		tbprint(carPosition-2, carY+i, termbox.ColorRed, termbox.ColorDefault, part)
	}

	// Display control information and score
	info := fmt.Sprintf("Controls: Left/Right arrows to move, ESC to exit | Difficulty: %s | Distance: %d", difficulties[difficultyIndex], distance)
	tbprint(0, totalView, termbox.ColorYellow, termbox.ColorDefault, info)
	tbprint(0, totalView+1, termbox.ColorYellow, termbox.ColorDefault, "Use Up/Down arrows to change difficulty")

	termbox.Flush()
}

func tbprint(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func generateObstacles() {
	newLine := "|" + strings.Repeat(" ", roadWidth-2) + "|"
	if rand.Intn(100) < obstacleChance {
		obstacleLength := rand.Intn(3) + 1
		obstaclePosition := rand.Intn(roadWidth-2-obstacleLength) + 1
		for i := 0; i < obstacleLength; i++ {
			newLine = newLine[:obstaclePosition+i] + "X" + newLine[obstaclePosition+i+1:]
		}
	}
	road = append([]string{newLine}, road[:len(road)-1]...)
	distance++
}

func checkCollision() bool {
	carY := totalView / 2
	for i := 0; i < 4; i++ {
		if road[carY][carPosition-2+i] == 'X' || road[carY+1][carPosition-2+i] == 'X' || road[carY+2][carPosition-2+i] == 'X' {
			return true
		}
	}
	return false
}

func adjustDifficulty() {
	switch difficultyIndex {
	case 0:
		obstacleChance = 10
		tickDuration = 3000 * time.Millisecond
	case 1:
		obstacleChance = 20
		tickDuration = 2500 * time.Millisecond
	case 2:
		obstacleChance = 30
		tickDuration = 2000 * time.Millisecond
	case 3:
		obstacleChance = 40
		tickDuration = 1500 * time.Millisecond
	case 4:
		obstacleChance = 50
		tickDuration = 1000 * time.Millisecond
	case 5:
		obstacleChance = 60
		tickDuration = 800 * time.Millisecond
	case 6:
		obstacleChance = 70
		tickDuration = 200 * time.Millisecond
	}
	if ticker != nil {
		ticker.Stop()
	}
	ticker = time.NewTicker(tickDuration)
	go func() {
		for range ticker.C {
			generateObstacles()
			if checkCollision() {
				termbox.Close()
				fmt.Println("\nGame Over! You hit an obstacle. Total distance: ", distance)
				os.Exit(0)
			}
			display()
		}
	}()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	adjustDifficulty()
	initializeRoad()
	display()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyEsc:
					return
				case termbox.KeyArrowLeft:
					if carPosition > 3 {
						carPosition--
					}
				case termbox.KeyArrowRight:
					if carPosition < roadWidth-4 {
						carPosition++
					}
				case termbox.KeyArrowUp:
					if difficultyIndex < len(difficulties)-1 {
						difficultyIndex++
						adjustDifficulty()
					}
				case termbox.KeyArrowDown:
					if difficultyIndex > 0 {
						difficultyIndex--
						adjustDifficulty()
					}
				}
				display()
			}
		}
	}
}
