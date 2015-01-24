package hello

type Content struct {
	comments  string
	upvotes   uint64
	subreddit string
}

// type Content interface {
// 	Prepare()
// }

// Gfycat content
type Gif struct {
	Content
	gameTitle string
	gifTitle  string
	gifId     string
}

// Youtube content
type Youtube struct {
	Content
}
