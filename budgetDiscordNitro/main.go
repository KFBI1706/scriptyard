package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Emojis is a globally accessible struct for getting the current emojis available to the bot
type Emojis struct {
	sync.RWMutex

	List []emoji
}

type emoji struct {
	ID       string `json:"id"`
	Animated bool   `json:"animated"`
}

// Variables used for command line parameters
var (
	Token  string
	emojis Emojis
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	go func() {
		for {
			dg.State.RLock()
			emojis.Lock()
			for _, g := range dg.State.Ready.Guilds {
				g, err := dg.Guild(g.ID)
				if err != nil {
					fmt.Printf("Error getting guild %s\n", g.ID)
					continue
				}
				emojis.List = make([]emoji, 0)
				for i := range g.Emojis {
					emojis.List = append(emojis.List, emoji{ID: g.Emojis[i].ID, Animated: g.Emojis[i].Animated})
				}
			}
			dg.State.RUnlock()
			emojis.Unlock()

			time.Sleep(5 * time.Second)
		}
	}()

	startAPI()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
