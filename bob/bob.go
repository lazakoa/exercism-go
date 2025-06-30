// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)
	alpha := "abcdefghijklmnopqrstuvwxyz"
	capitalAlpha := strings.ToUpper(alpha)
	if trimmedRemark == "" {
		return "Fine. Be that way!"
	} else if strings.HasSuffix(trimmedRemark, "?") && strings.ToUpper(trimmedRemark) == trimmedRemark {
		if strings.ContainsAny(trimmedRemark, alpha+capitalAlpha) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Sure."
		}
	} else if strings.HasSuffix(trimmedRemark, "?") {
		return "Sure."
	} else if strings.ToUpper(trimmedRemark) == trimmedRemark {
		if strings.ContainsAny(trimmedRemark, alpha+capitalAlpha) {
			return "Whoa, chill out!"
		} else {
			return "Whatever."
		}
	} else {
		return "Whatever."
	}
}
