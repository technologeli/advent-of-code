package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Set struct {
    Red int
    Green int
    Blue int
}

type Game struct {
    Id int
    Sets []Set
}

func extract(text string) Game {
    var g Game

    strs := strings.Split(text, ": ")
    sid := strings.Split(strs[0], " ")[1]
    id, _ := strconv.Atoi(sid)

    g.Id = id

    sets := strings.Split(strs[1], "; ")
    for _, set := range sets {
        var s Set
        scubes := strings.Split(set, ", ")
        for _, scube := range scubes {
            cube := strings.Split(scube, " ")
            cubeCount, _ := strconv.Atoi(cube[0])
            if (cube[1][0] == 'r') {
                s.Red = cubeCount
            } else if (cube[1][0] == 'g') {
                s.Green = cubeCount
            } else {
                s.Blue = cubeCount
            }
        }

        g.Sets = append(g.Sets, s)
    }

    return g
}

func main() {
    reader := bufio.NewReader(os.Stdin)	
    sum := 0
    for {
        text, err := reader.ReadString('\n')
        if err == io.EOF {
            break
        }

        game := extract(text)
        // if game.legal() {
        //     sum += game.Id
        // }
        sum += game.power()
    }
    fmt.Printf("Sum is %d\n", sum)
}

func (g Game) legal() (bool) {
    const red int = 12
    const green int = 13
    const blue int = 14

    for _, set := range g.Sets {
        if (set.Red > red || set.Green > green || set.Blue > blue) {
            return false
        }
    }

    return true
}

func (g Game) power() (int) {
    red := 0
    green := 0
    blue := 0

    for _, set := range g.Sets {
        if (set.Red > red) {
            red = set.Red
        }

        if (set.Green > green) {
            green = set.Green
        }

        if (set.Blue > blue) {
            blue = set.Blue
        }
    }
    
    return red * green * blue
}
