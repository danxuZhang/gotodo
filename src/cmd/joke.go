// Package cmd
// Copyright © 2023 dxZhang <dxzhang49@gmail.com>
// /*
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

// Joke Represents a requested joke.
type Joke struct {
	Id     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	fmt.Println("Let's get a random joke...")
	const URL string = "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(URL)
	joke := Joke{}

	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		log.Printf("Canot unmarshal joke: %v.", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseApi string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseApi,
		nil)

	if err != nil {
		log.Printf("Cannot request a joke: %v.", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "gotodo "+
		"(github.com/danxuZhang/gotodo)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Cannot request: %v.", err)
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Cannot request: %v.", err)
	}

	return responseBytes
}

// jokeCmd represents the joke command
var jokeCmd = &cobra.Command{
	Use:   "joke",
	Short: "Get random dad joke in your terminal",
	Long: `This command fetches a dad joke from the icanhazdadjoke API and 
			prints on the terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(jokeCmd)
}
