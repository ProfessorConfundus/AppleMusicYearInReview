package main

type song struct {
	name             string
	artist           string
	length_ms        int
	plays            int // ALL plays, including plays reported in plays_via_siri
	time_listened_ms int
	skips            int
	plays_via_siri   int
}

type artist struct {
	name             string
	plays            int
	time_listened_ms int
}

type genre struct {
	name             string
	plays            int
	time_listened_ms int
}
