package content

type Content struct {
	url       string
	upvotes   uint
	subreddit string
}

type Content interface {
	Prepare()
}

// Gfycat content
type Gif struct {
	Content
}

// Youtube content
type Youtube struct {
	Content
}
