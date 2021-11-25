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

type compareFunc func(x, y *Track) bool

type tier struct {
	name      string
	lessFunc  compareFunc
	equalFunc compareFunc
}

var (
	titleTier = tier{
		name: "title",
		lessFunc: func(x, y *Track) bool {
			return x.Title < y.Title
		},
		equalFunc: func(x, y *Track) bool {
			return x.Title == y.Title
		},
	}
	artistTier = tier{
		name: "artist",
		lessFunc: func(x, y *Track) bool {
			return x.Artist < y.Artist
		},
		equalFunc: func(x, y *Track) bool {
			return x.Artist == y.Artist
		},
	}
	albumTier = tier{
		name: "album",
		lessFunc: func(x, y *Track) bool {
			return x.Album < y.Album
		},
		equalFunc: func(x, y *Track) bool {
			return x.Album == y.Album
		},
	}
	yearTier = tier{
		name: "year",
		lessFunc: func(x, y *Track) bool {
			return x.Year < y.Year
		},
		equalFunc: func(x, y *Track) bool {
			return x.Year == y.Year
		},
	}
	lengthTier = tier{
		name: "length",
		lessFunc: func(x, y *Track) bool {
			return x.Length < y.Length
		},
		equalFunc: func(x, y *Track) bool {
			return x.Length == y.Length
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
	for _, t := range s.tiers {
		if !t.equalFunc(s.tracks[i], s.tracks[j]) {
			return t.lessFunc(s.tracks[i], s.tracks[j])
		}
	}

	return false
}

func (s MultiTierSort) Swap(i, j int) {
	s.tracks[i], s.tracks[j] = s.tracks[j], s.tracks[i]
}
