/*
This file is:

The MIT License (MIT)

Copyright (c) 2014 Bitrise

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package main

import (
	"strings"
)

// See also: https://docs.microsoft.com/en-us/outlook/actionable-messages/message-card-reference#actions
type Message struct {
	Context    string    `json:"@context"`
	Type       string    `json:"@type"`
	ThemeColor string    `json:"themeColor,omitempty"`
	Title      string    `json:"title,omitempty"`
	Summary    string    `json:"summary,omitempty"`
	Sections   []Section `json:"sections,omitempty"`
}

type Section struct {
	ActivityTitle string   `json:"activityTitle,omitempty"`
	ActivityText  string   `json:"activityText,omitempty"`
	Facts         []Fact   `json:"facts,omitempty"`
	Images        []Image  `json:"images,omitempty"`
	Actions       []Action `json:"potentialAction,omitempty"`
}

type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func parsesFacts(s string) (fs []Fact) {
	for _, p := range pairs(s) {
		fs = append(fs, Fact{Name: p[0], Value: p[1]})
	}
	return
}

type Image struct {
	URL   string `json:"image"`
	Title string `json:"title"`
}

func parsesImages(s string) (is []Image) {
	for _, p := range pairs(s) {
		is = append(is, Image{Title: p[0], URL: p[1]})
	}
	return
}

type Action struct {
	Type    string   `json:"@type"`
	Name    string   `json:"name"`
	Targets []Target `json:"targets,omitempty"`
}

type Target struct {
	OS  string `json:"os"`
	URI string `json:"uri"`
}

func parsesActions(s string) (as []Action) {
	for _, p := range pairs(s) {
		as = append(as, Action{
			Type: "OpenUri",
			Name: p[0],
			Targets: []Target{{
				OS:  "default",
				URI: p[1],
			}},
		})
	}
	return
}

// pairs slices every lines in s into two substrings separated by the first pipe
// character and returns a slice of those pairs.
func pairs(s string) [][2]string {
	var ps [][2]string
	for _, line := range strings.Split(s, "\n") {
		a := strings.SplitN(line, "|", 2)
		if len(a) == 2 && a[0] != "" && a[1] != "" {
			ps = append(ps, [2]string{a[0], a[1]})
		}
	}
	return ps
}
