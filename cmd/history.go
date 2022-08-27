package cmd

type HistoryInput struct {
	// history limit of each channel
	Limit int
	// channel IDs
	ChannelIDs []string
}
