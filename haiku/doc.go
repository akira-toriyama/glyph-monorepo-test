// Package haiku is one independently versioned module of the harness: a small
// saijiki of Japanese seasonal poems. Its version line is the tag prefix
// haiku/ — nothing here depends on curry.
//
// Stable from v1.0.0. Inside a major line an exported name is permanent and a
// poem is always three lines. A poem's text is not part of that promise: a
// closer reading of the original ships as a patch, so a caller must not compare
// against a copy it has kept.
package haiku
