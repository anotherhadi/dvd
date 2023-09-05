// https://github.com/anotherhadi/dvd
package main

import (
	"fmt"
	"github.com/anotherhadi/ansi"
	"github.com/anotherhadi/getsize"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func printLogo(logo []string, x, y int, color uint8) {
	for i, line := range logo {
		ansi.CursorMove(y+i, x)
		fmt.Print("\033[3" + strconv.FormatUint(uint64(color), 10) + "m")
		fmt.Println(line)
	}
}

func cleanup() {
	ansi.ClearScreen()
	ansi.CursorVisible()
	ansi.ScreenRestore()
	ansi.CursorRestore()
}

func main() {

	w, h, err := getsize.GetSize()
	if err != nil {
		panic("Error while getting the terminal size")
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(0)
	}()

	logo := make([]string, 0)
	logo = append(logo, "  ⣸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⢀⣾⣿⣿⣿⣿⣿⣿⣿⣿⣶⣦⡀")
	logo = append(logo, "⠀⢠⣿⣿⡿⠀⠀⠈⢹⣿⣿⡿⣿⣿⣇⠀⣠⣿⣿⠟⣽⣿⣿⠇⠀⠀⢹⣿⣿⣿")
	logo = append(logo, "⠀⢸⣿⣿⡇⠀⢀⣠⣾⣿⡿⠃⢹⣿⣿⣶⣿⡿⠋⢰⣿⣿⡿⠀⠀⣠⣼⣿⣿⠏")
	logo = append(logo, "⠀⣿⣿⣿⣿⣿⣿⠿⠟⠋⠁⠀⠀⢿⣿⣿⠏⠀⠀⢸⣿⣿⣿⣿⣿⡿⠟⠋⠁⠀")
	logo = append(logo, "⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣀⣀⣸⣟⣁⣀⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀")
	logo = append(logo, "⣠⣴⣶⣾⣿⣿⣻⡟⣻⣿⢻⣿⡟⣛⢻⣿⡟⣛⣿⡿⣛⣛⢻⣿⣿⣶⣦⣄⡀⠀")
	logo = append(logo, "⠉⠛⠻⠿⠿⠿⠷⣼⣿⣿⣼⣿⣧⣭⣼⣿⣧⣭⣿⣿⣬⡭⠾⠿⠿⠿⠛⠉⠀ ")

	var logo_w int = 0
	for _, line := range logo {
		if len([]rune(line)) > logo_w {
			logo_w = len([]rune(line))
		}
	}
	logo_h := len(logo)

	var color uint8
	var max_x, max_y int = w - logo_w, h - logo_h
	var x, y int
	x = rand.Intn(max_x)
	y = rand.Intn(max_y)
	color = uint8(rand.Uint32())%7 + 1

	// Up-Left, Up-Right, Down-Left, Down-Right
	var direction string = "ur"

	ansi.CursorSave()
	ansi.ScreenSave()
	ansi.ClearScreen()
	ansi.CursorInvisible()
	for {
		ansi.ClearScreen()
		printLogo(logo, x, y, color)
		if x >= max_x {
			if direction == "ur" {
				direction = "ul"
			} else if direction == "dr" {
				direction = "dl"
			}
			color++
		} else if y <= 1 {
			if direction == "ur" {
				direction = "dr"
			} else if direction == "ul" {
				direction = "dl"
			}
			color++
		} else if x <= 1 {
			if direction == "ul" {
				direction = "ur"
			} else if direction == "dl" {
				direction = "dr"
			}
			color++
		} else if y >= max_y {
			if direction == "dr" {
				direction = "ur"
			} else if direction == "dl" {
				direction = "ul"
			}
			color++
		}

		if color > 7 {
			color = 1
		}

		if direction == "ul" {
			x -= 2
			y -= 1
		} else if direction == "ur" {
			x += 2
			y -= 1
		} else if direction == "dl" {
			x -= 2
			y += 1
		} else if direction == "dr" {
			x += 2
			y += 1
		}
		time.Sleep(150 * time.Millisecond)
	}

}
