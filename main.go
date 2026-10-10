package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jjgmckenzie/podcastindex"
)

func main() {

	// Get podcast search from cli
	if len(os.Args) < 2 {
		log.Fatalln("Usage: podcastindex <name>")
	}
	cmdLineArg := os.Args[1]

	var options podcastindex.NewClientOptions
	options = envSetup()
	client := podcastindex.NewClient(options)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	searchParams := podcastindex.SearchPodcastsByTermParams{
		Max:      40,
		FullText: true,
		Similar:  true,
	}

	podcasts, err := client.SearchPodcastsByTerm(ctx, cmdLineArg, &searchParams)
	if err != nil {
		fmt.Printf("failed to search podcasts: %v", err)
	}
	for _, podcast := range podcasts {
		fmt.Printf("%s\n", podcast.Title)
		episodes, err := client.GetEpisodes(ctx, *podcast, nil)
		if err != nil {
			fmt.Printf("Could not find episodes: %v", err)
		}
		for _, episode := range *episodes {
			fmt.Printf("%s: %s\n", podcast.Title, episode.Title)
		}
	}
}

func envSetup() podcastindex.NewClientOptions {
	key := os.Getenv("PODCASTINDEX_API_KEY")
	if key == "" {
		log.Fatal("Could not load API key")
	}
	secret := os.Getenv("PODCASTINDEX_API_SECRET")
	if secret == "" {
		log.Fatal("Could not load secret")
	}

	// Guess this will work for now
	userAgent := "podcast_index"

	return podcastindex.NewClientOptions{
		APIKey:    key,
		APISecret: secret,
		UserAgent: userAgent,
	}

}
