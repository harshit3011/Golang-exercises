package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	id    int
	name  string
	price int
}

type Game struct {
	item  Item
	genre string
}

func main() {
	games := []Game{
		{
			item: Item{
				1, "god of war", 50,
			},
			genre: "action adventure",
		},
		{
			item: Item{
				2, "x-com 2", 30,
			},
			genre: "strategy",
		},
		{
			item: Item{
				3, "minecraft", 20,
			},
			genre: "sandbox",
		},
	}

	byId:=make(map[int]Game)
	for _,g:=range games{
		byId[g.item.id]=g
	}
	fmt.Printf("Harry's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > save   : exports the data to json and quits
  > quit   : quits

`)

		if !in.Scan(){
			break
		}
		fmt.Println()
		cmd:= strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.item.id, g.item.name, "("+g.genre+")", g.item.price)
			}

		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := byId[id]
			if !ok {
				fmt.Println("sorry. I don't have the game")
				continue
			}

			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.item.id, g.item.name, "("+g.genre+")", g.item.price)

		case "save":
			type jsonGame struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Genre string `json:"genre"`
				Price int    `json:"price"`
			}

			// load the data into the encodable game values
			var encodable []jsonGame
			for _, g := range games {
				encodable = append(encodable,
					jsonGame{g.item.id, g.item.name, g.genre, g.item.price})
			}

			out, err := json.MarshalIndent(encodable, "", "\t")
			if err != nil {
				fmt.Println("Sorry:", err)
				continue
			}

			fmt.Println(string(out))
			return
	}
}
}