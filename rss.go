package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func handlerFeeds(s *state, cmd command) error {
	context := context.Background()
	//;cmd.Args[0]
	val, err := fetchFeed(context, "https://www.wagslane.dev/index.xml")

	if err != nil {
		fmt.Println(cmd.Name, " was successful.")
		os.Exit(1)
		return err
	}
	fmt.Println(val)
	os.Exit(0)
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)

	req.Header.Set("User-Agent", "gator")
	if err != nil {
		return &RSSFeed{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var issues RSSFeed
	if err := xml.Unmarshal(data, &issues); err != nil {
		fmt.Print(err)
		return nil, err
	}
	issues.Channel.Description = html.UnescapeString(issues.Channel.Description)
	issues.Channel.Title = html.UnescapeString(issues.Channel.Title)
	issues.Channel.Link = html.UnescapeString(issues.Channel.Link)

	for i, v := range issues.Channel.Item {
		issues.Channel.Item[i].Title = html.UnescapeString(v.Title)
		issues.Channel.Item[i].Description = html.UnescapeString(v.Description)
		issues.Channel.Item[i].Link = html.UnescapeString(v.Link)
	}

	return &issues, nil
}
