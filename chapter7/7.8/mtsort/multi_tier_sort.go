package mtsort

import (
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type tier struct {
	name     string
	lessFunc func(x, y *Track) bool
}

var (
	titleTier = tier{
		name: "title",
		lessFunc: func(x, y *Track) bool {
			return x.Title < y.Title
		},
	}
	artistTier = tier{
		name: "artist",
		lessFunc: func(x, y *Track) bool {
			return x.Artist < y.Artist
		},
	}
	albumTier = tier{
		name: "album",
		lessFunc: func(x, y *Track) bool {
			return x.Album < y.Album
		},
	}
	yearTier = tier{
		name: "year",
		lessFunc: func(x, y *Track) bool {
			return x.Year < y.Year
		},
	}
	lengthTier = tier{
		name: "length",
		lessFunc: func(x, y *Track) bool {
			return x.Length < y.Length
		},
	}
)

type MultiTierSort struct {
	tracks []*Track
	tiers  []tier
}

func (s MultiTierSort) Len() int {
	return len(s.tracks)
}

func (s MultiTierSort) Less(i, j int) bool {
	// prioritize tiers by order
	// if tracks[i] and tracks[j] are not equal, determine if we should swap using the less function of the current tier
	// otherwise, proceed to the next tier
	for _, t := range s.tiers {
		equal := !t.lessFunc(s.tracks[i], s.tracks[j]) && !t.lessFunc(s.tracks[j], s.tracks[i])
		if !equal {
			return t.lessFunc(s.tracks[i], s.tracks[j])
		}
	}

	return false
}

func (s MultiTierSort) Swap(i, j int) {
	s.tracks[i], s.tracks[j] = s.tracks[j], s.tracks[i]
}
