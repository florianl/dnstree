package dnstree

import (
	"regexp"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name string
		dns  string
		rank int
		err  string
	}{
		{name: "google.com", dns: "google.com", rank: 1},
		{name: "google.ch", dns: "google.ch", rank: 2},
		{name: "google.de", dns: "google.de", rank: 3},
		{name: "abc.xyz", dns: "abc.xyz", rank: 4},
		{name: "github.com", dns: "github.com", rank: 5},
		{name: "www.google.com", dns: "www.google.com", rank: 6},
		{name: "www.google.com - Exists", dns: "www.google.com", rank: 7, err: "Element already exists"},
		{name: "www.☺️.tld", dns: "www.☺️.tld", rank: 8},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := Insert(tc.dns, tc.rank)
			if err != nil || len(tc.err) > 0 {
				if matched, _ := regexp.MatchString(tc.err, err.Error()); matched == false {
					t.Fatalf("Error matching regex: %v \t Got: %v", tc.err, err)
				}
			}
		})
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name        string
		destination string
		rank        int
		err         string
	}{
		{name: "google.com", destination: "google.com", rank: 1},
		{name: "abc.xyz", destination: "abc.xyz", rank: 4},
		{name: "oogle.com", destination: "oogle.com", rank: 1, err: "Not found"},
		{name: "baidu.com", destination: "baidu.com", rank: 1, err: "Not found"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rank, err := Search(tc.destination)
			if err != nil || len(tc.err) > 0 {
				if matched, _ := regexp.MatchString(tc.err, err.Error()); matched == false {
					t.Fatalf("Error matching regex: %v \t Got: %v", tc.err, err)
				}
			} else {
				if rank != tc.rank {
					t.Fatalf("%s\texpected: %d\tgot: %d", tc.name, tc.rank, rank)
				}
			}
		})
	}
}
