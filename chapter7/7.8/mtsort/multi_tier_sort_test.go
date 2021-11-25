package mtsort

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Master Of Puppets", "Metallica", "Master Of Puppets", 1986, length("8m37s")},
}

func TestMultiTierSort(t *testing.T) {
	// 1 tier
	s := MultiTierSort{
		tracks: tracks,
		tiers:  []tier{titleTier},
	}

	sort.Sort(s)
	equalTracks(t,
		[]*Track{
			{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			{"Go", "Moby", "Moby", 1992, length("3m37s")},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			{"Master Of Puppets", "Metallica", "Master Of Puppets", 1986, length("8m37s")},
			{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		},
		s.tracks,
	)

	// 2 tiers
	s = MultiTierSort{
		tracks: tracks,
		tiers:  []tier{titleTier, yearTier},
	}

	sort.Sort(s)
	equalTracks(t,
		[]*Track{
			{"Go", "Moby", "Moby", 1992, length("3m37s")},
			{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			{"Master Of Puppets", "Metallica", "Master Of Puppets", 1986, length("8m37s")},
			{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		},
		s.tracks,
	)

	// reverse 2 tiers
	sort.Sort(sort.Reverse(s))
	equalTracks(t,
		[]*Track{
			{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
			{"Master Of Puppets", "Metallica", "Master Of Puppets", 1986, length("8m37s")},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			{"Go", "Moby", "Moby", 1992, length("3m37s")},
		},
		s.tracks,
	)

	// 3 tiers
	tracks = append(tracks, &Track{
		Title:  "Go Ahead",
		Artist: "Fake Artist",
		Album:  "As I Am",
		Year:   2007,
		Length: length("1m1s"),
	})

	s = MultiTierSort{
		tracks: tracks,
		tiers:  []tier{titleTier, yearTier, lengthTier},
	}

	sort.Sort(s)
	equalTracks(t,
		[]*Track{
			{"Go", "Moby", "Moby", 1992, length("3m37s")},
			{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			{"Go Ahead", "Fake Artist", "As I Am", 2007, length("1m1s")},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			{"Master Of Puppets", "Metallica", "Master Of Puppets", 1986, length("8m37s")},
			{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		},
		s.tracks,
	)
}

func BenchmarkMultiTierSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := MultiTierSort{
			tracks: tracks,
			tiers: []tier{
				titleTier,
				artistTier,
				albumTier,
				yearTier,
				lengthTier,
			},
		}

		sort.Sort(s)
	}
}

func BenchmarkStableMultiTierSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := MultiTierSort{
			tracks: tracks,
			tiers: []tier{
				titleTier,
				artistTier,
				albumTier,
				yearTier,
				lengthTier,
			},
		}

		sort.Stable(s)
	}
}

func equalTracks(t *testing.T, expected []*Track, actual []*Track) {
	exp := make([]Track, 0)
	for _, e := range expected {
		exp = append(exp, *e)
	}

	act := make([]Track, 0)
	for _, a := range actual {
		act = append(act, *a)
	}

	assert.Equal(t, exp, act)
}
