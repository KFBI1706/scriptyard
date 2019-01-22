package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

const (
	//PAD boxes
	PAD = 2
)

var (
	weeks []week
	//AvatarURL is a link to the github avatar for global use
	AvatarURL string
)

type day struct {
	ContributionCount githubv4.Int
	Color             githubv4.String
}

type week struct {
	ContributionDays []day
}

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	var query struct {
		Viewer struct {
			ContributionsCollection struct {
				ContributionCalendar struct {
					Weeks              []week
					TotalContributions githubv4.Int
					Colors             []githubv4.String
				}
			}
			Login     githubv4.String
			CreatedAt githubv4.DateTime
			AvatarURL string `graphql:"avatarUrl(size: 72)"`
		}
	}
	err := client.Query(context.Background(), &query, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(query.Viewer.Login)
	fmt.Println("Joined", query.Viewer.CreatedAt.Format(time.RFC850))
	fmt.Println(query.Viewer.ContributionsCollection.ContributionCalendar.TotalContributions, "Contributions in the last year")
	fmt.Println("Colors ", query.Viewer.ContributionsCollection.ContributionCalendar.Colors)
	weeks = query.Viewer.ContributionsCollection.ContributionCalendar.Weeks
	lastDays := weeks[len(weeks)-1].ContributionDays
	lastDay := lastDays[len(lastDays)-1]
	fmt.Printf("You made %d contributions today!\n", lastDay.ContributionCount)

	fmt.Println("Weeks ", weeks)
	AvatarURL = query.Viewer.AvatarURL

	fmt.Println("Avatar Url: ", AvatarURL)

	termInit()

}
