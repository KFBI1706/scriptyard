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

	List []Emoji
}

// Emoji is a small struct with the basic information about an emoji
type Emoji struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	Animated bool   `json:"animated"`
}

// Variables used for command line parameters
var (
	Token     string
	GitCommit string
	emojis    Emojis
)

func init() {
	discordAPIToken := os.Getenv("DISCORD_API_TOKEN")
	flag.StringVar(&Token, "t", discordAPIToken, "Bot Token")
	flag.Parse()

	fmt.Printf("Token: %s\n", Token)

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
				emojis.List = make([]Emoji, 0)
				for i := range g.Emojis {
					emojis.List = append(emojis.List, Emoji{Name: g.Emojis[i].Name, URL: fmt.Sprintf("https://cdn.discordapp.com/emojis/%s.png", g.Emojis[0].ID), Animated: g.Emojis[i].Animated})
				}
			}
			dg.State.RUnlock()
			emojis.Unlock()

			time.Sleep(5 * time.Second)
		}
	}()

	go startAPI()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
