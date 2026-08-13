package main

import (
	"testing"
	"time"
)

func TestMondayOffset(t *testing.T) {
	tests := []struct {
		day      time.Weekday
		want     int
		date     string
	}{
		{time.Monday, 0, "2026-01-05"},
		{time.Tuesday, 1, "2026-01-06"},
		{time.Sunday, 6, "2026-01-04"},
		{time.Thursday, 3, "2026-01-01"},
	}

	for _, tt := range tests {
		d, err := time.ParseInLocation("2006-01-02", tt.date, time.UTC)
		if err != nil {
			t.Fatal(err)
		}
		if d.Weekday() != tt.day {
			t.Fatalf("%s: got weekday %v, want %v", tt.date, d.Weekday(), tt.day)
		}
		if got := mondayOffset(d); got != tt.want {
			t.Errorf("%s: mondayOffset=%d, want %d", tt.date, got, tt.want)
		}
	}
}

func TestSplitEmails(t *testing.T) {
	got := splitEmails(" a@b.c , d@e.f,, ")
	if len(got) != 2 || got[0] != "a@b.c" || got[1] != "d@e.f" {
		t.Fatalf("unexpected: %#v", got)
	}
}
