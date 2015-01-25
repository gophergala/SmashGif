package hello

import (
	"appengine"
	"appengine/datastore"

	"log"
	"math/rand"
	"net/url"
	"strconv"
)

var (
	games = [4]string{
		"Melee",
		"SSB4",
		"64",
		"Project M",
	}
)

func storeGif(g Gif, c appengine.Context) {
	key := datastore.NewKey(c, "Gif", g.GifId, 0, nil)
	_, err := datastore.Put(c, key, &g)
	check(err)
}

func queryNext(v url.Values, c appengine.Context) Gif {
	game := getGame(v)

	// Depending on whether the count is even or odd, I pick
	// a GIF with higher or lower upvotes
	// The offset is random. By having an oscillating quality,
	// I'm hoping the user stays on longer
	upvoteFilter := "Content.Upvotes >="
	count, err := strconv.Atoi(v.Get("count"))
	check(err)

	if count%4 == 0 {
		upvoteFilter = "Content.Upvotes <="
	}

	upvotesInt := rand.Intn(1000)

	// First call doesn't have any, default to 100 and oscillate there
	upvotes := v.Get("upvotes")
	if upvotes != "" {
		upvotesInt, err = strconv.Atoi(upvotes)
		check(err)
	}

	log.Println("Searching for game: ", game)
	// Go doesn't allow line to start with . because of no semi colons :(
	q := datastore.NewQuery("Gif").
		Filter("GameTitle =", game).
		Filter(upvoteFilter, upvotesInt).
		Limit(10)

	var result []Gif
	_, err = q.GetAll(c, &result)
	check(err)

	return result[rand.Intn(len(result))]
}

func getGame(v url.Values) string {
	var validGames []string
	for _, g := range games {
		if v.Get(g) == "0" {
			log.Println("Valid Game:", g)
			validGames = append(validGames, g)
		}
	}

	// AppEngine doesn't support Composite filters so we're gonna have to
	// randomly pick which game.
	// If all is filtered, we fallback to everything
	length := 4
	if len(validGames) != 0 {
		length = len(validGames)
	}
	log.Println("Length: ", length)

	game := "Melee" // Default to Melee because it's the best :)
	if length == 4 {
		index := rand.Intn(length)
		log.Println(index, games)
		game = games[index]
	} else {
		game = validGames[rand.Intn(length)]
	}
	log.Println(game)
	return game
}
