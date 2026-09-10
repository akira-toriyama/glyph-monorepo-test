# glyph-monorepo-test

Permanent live-fire harness for [glyph](https://github.com/akira-toriyama/glyph)'s
monorepo support: several independently versioned Go modules in one repository,
each with its own tag line and its own rolling draft release.

The property the harness exists to prove: a pull request that touches `haiku/`
moves **haiku's** version and nothing else, and a squash merge never loses which
of the pull's commits touched which module.

| module | tag line | what a change looks like |
|---|---|---|
| `haiku/` | `haiku/vX.Y.Z` | a new season is a `^` minor; a typo fix is a `~` patch |
| `curry/` | `curry/vX.Y.Z` | a new ingredient is a `^` minor; a swapped one is a `~` patch |

Tags follow the Go multi-module convention (`<dir>/vX.Y.Z`), so each module is
also fetchable with `go get` at its own version. Nothing is built or published
beyond the draft releases — the modules are the smallest thing that has a
version.

Like [glyph-test](https://github.com/akira-toriyama/glyph-test), real pull
requests, releases and tags here are fair game, and its history and tags are
frozen coordinates once an end-to-end workflow references them.
