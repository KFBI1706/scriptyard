package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

func main() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewClient(httpClient)

	type day struct {
		ContributionCount githubv4.Int
	}
	type week struct {
		ContributionDays []day
	}

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
	weeks := query.Viewer.ContributionsCollection.ContributionCalendar.Weeks
	lastDays := weeks[len(weeks)-1].ContributionDays
	lastDay := lastDays[len(lastDays)-1]
	fmt.Printf("You made %d contributions today!\n", lastDay.ContributionCount)

	fmt.Println("Weeks ", query.Viewer.ContributionsCollection.ContributionCalendar.Weeks)

}
