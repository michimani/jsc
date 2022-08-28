package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/michimani/jsc/cmd"
	"github.com/michimani/jsc/slack"
)

type channels []string

func (c *channels) String() string {
	return fmt.Sprintf("%s", *c)
}

func (c *channels) Set(value string) error {
	*c = append(*c, value)
	return nil
}

var (
	inputC   channels
	inputT   *string = flag.String("t", "", "Slack OAuth Token")
	inputL   *int    = flag.Int("l", 10, "limit of messages to be retrieved for each channel (default: 10)")
	inputO   *string = flag.String("o", "", "output type (default: timeline)")
	version  string
	revision string
)

func usage() {
	format := `
     _
    (_)___  ___
    | / __|/ __|
    | \__ \ (__
   _/ |___/\___|
  |__/ Version: %s-%s

Usage:
  jsc [flags] [values]
Flags:
  -c (string)
    ID of the channels you want to join.
    This option can be used multiple times to specify multiple channels.

  -l (integer)
    Limit of messages to be retrieved for each channel (default: 10)

  -o (string)
    Specify the output method as 'json' or 'timeline'. The default is 'timeline',
    which outputs in a format suitable for display.

  -t (string)
    Specify the OAuth Token for the Slack API. It must either be specified
    with this option or set in the environment variable 'SLACK_OAUTH_TOKEN'.
    If both are set, the one specified in this option takes precedence.

Example:
  jsc -c CHANNEL1 -c CHANNEL2 -l 5
`
	fmt.Fprintln(os.Stderr, fmt.Sprintf(format, version, revision))
}

func main() {
	flag.Usage = usage
	flag.Var(&inputC, "c", "channel IDs for join")
	flag.Parse()

	os.Exit(run())
}

func run() int {
	// Slack OAuth Token
	token := os.Getenv("SLACK_OAUTH_TOKEN")
	if inputT != nil && *inputT != "" {
		token = *inputT
	}

	if token == "" {
		fmt.Println("OAuth Token for Slack is not specified. Set the environment variable `SLACK_OAUTH_TOKEN` or specify it with the `-t` option.")
		return 1
	}

	// Channels
	if len(inputC) == 0 {
		fmt.Println("At least one channel must be specified")
		return 1
	}

	c, err := slack.NewSlackClient(&slack.NewSlackClientInput{
		Token: token,
	})
	if err != nil {
		fmt.Printf("Failed to generate Slack API client: %v\n", err)
		return 1
	}

	joined, err := cmd.GetJoinedHistory(context.TODO(), c, &cmd.HistoryInput{
		Limit:      *inputL,
		ChannelIDs: inputC,
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output := cmd.OutputTypeTimeline
	if inputO != nil && *inputO == "json" {
		output = cmd.OutputTypeJSON
	}

	switch output {
	case cmd.OutputTypeTimeline:
		for _, m := range joined.Messages {
			fmt.Println("------------------------------------------------------------")
			fmt.Printf("%s\n", m.PostedAt.Format(time.RFC3339))
			fmt.Printf("%s\n", m.Username)
			fmt.Printf("%s\n\n", m.MessageURL)
			fmt.Printf("%s\n\n", m.Text)
		}
	case cmd.OutputTypeJSON:
		printAsJSON(joined)
	}

	return 0
}

func printAsJSON(res any) {
	j, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(string(j))
}
