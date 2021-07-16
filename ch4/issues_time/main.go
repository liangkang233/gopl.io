package main

import (
	"fmt"
	"gopl/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	format := "#%-5d %9.9s %.55s\n"
	now := time.Now()

	pastDay := make([]*github.Issue, 0)
	pastMonth := make([]*github.Issue, 0)
	pastYear := make([]*github.Issue, 0)
	pastMoreYear := make([]*github.Issue, 0)
	a, b, c, d := 0, 0, 0, 0

	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)
	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(day):
			pastDay = append(pastDay, item)
			a++
		case item.CreatedAt.After(month) && item.CreatedAt.Before(day):
			pastMonth = append(pastMonth, item)
			b++
		case item.CreatedAt.After(year) && item.CreatedAt.Before(month):
			pastYear = append(pastYear, item)
			c++
		default:
			pastMoreYear = append(pastMoreYear, item)
			d++
		}

	}

	if len(pastDay) > 0 {
		fmt.Printf("\nCount:%d\tPast day:\n", a)
		for _, item := range pastDay {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
	if len(pastMonth) > 0 {
		fmt.Printf("\nCount:%d\tPast month:\n", b)
		for _, item := range pastMonth {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
	if len(pastYear) > 0 {
		fmt.Printf("\nCount:%d\tPast year:\n", c)
		for _, item := range pastYear {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
	if len(pastMoreYear) > 0 {
		fmt.Printf("\nCount:%d\tPast more year:\n", d)
		for _, item := range pastMoreYear {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
}
