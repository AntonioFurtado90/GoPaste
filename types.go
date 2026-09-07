package main

import "time"

// Paste represents a single saved text snippet.
type Paste struct {
	ID        string    // random identifier used in the paste's URL
	Content   string    // the text submitted by the user
	CreatedAt time.Time // when the paste was created
}
