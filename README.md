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
| `travel/` | `travel/vX.Y.Z` | a new stop is a `^` minor; a reordered one is a `~` patch |
| `travel/onsen/` | `travel/onsen/vX.Y.Z` | **nested** in `travel/`: a new bath moves onsen alone — the longest declared prefix wins |
| `camp/` | `camp/vX.Y.Z` | a new item is a `^` minor; a swapped one is a `~` patch |

Tags follow the Go multi-module convention (`<dir>/vX.Y.Z`), so each module is
also fetchable with `go get` at its own version. Nothing is built or published
beyond the draft releases — the modules are the smallest thing that has a
version.

Like [glyph-test](https://github.com/akira-toriyama/glyph-test), real pull
requests, releases and tags here are fair game, and its history and tags are
frozen coordinates once an end-to-end workflow references them.

## Baselining a new line

A line with no tag walks its whole history and reports `v0.0.0` as its current
version. To say "nothing before here is a release of this module", add the
module with a `=` commit (it moves no line) and, on the commit that lands it,
cut `<path>/v0.0.0` by hand:

```sh
git tag -a travel/v0.0.0 <merge-commit> -m 'travel/v0.0.0 — the baseline'
git push origin travel/v0.0.0
```

The next change under `travel/` then steps from `v0.0.0`. Publishing a draft is
the human act that creates every other tag; glyph never tags.

## What the workflows prove

| workflow | fires | proves |
|---|---|---|
| `release.yml` | every push to `main` | glyph's reusable in packages mode: the four scalar outputs are `""`, `packages` lists the five lines in config order without body or url, the live releases list holds at most one draft per line and the tag each verdict names, and no bare `vX.Y.Z` residue stands |
| `artifact-probe.yml` | by hand, **expected to fail** | the reusable refuses `binary:` once the verdict says packages — a green run is the finding |
| `e2e.yml` | by hand, `glyph-ref` = the glyph ref to build | the attribution rules on a fixture built from nothing (no credential), the same rules at this repository's frozen coordinate over the API, one real write that converges a planted bare draft and a planted stray, and the pin sites resolving to one version |
| `commit-lint.yml` / `version-preview.yml` | every pull request | the fleet's own gates, reading `[[packages]]`: a shared-only `^` is refused before merge, and the merge preview posts one headline per touched line |

Frozen coordinates the e2e reads (never rewrite them):

- #2 — a `^` under `haiku/` and a `~` under `curry/` in one squash: the two lines move differently.
- #4 — open on purpose: the preview probe for haiku/curry.
- #10 — a `^` under `travel/` and a `~` under `camp/` in one squash; `travel/onsen` is not mentioned.
- #11 — a `^` under `travel/onsen/` alone: the nested line moves, its parent does not.
- merge commit of #11 (`380db9d`): every literal in the walk tier was measured there.

Measured by hand rather than in e2e, because each consumes a tag or a version:

- publish one line's draft, merge again: the published line gets a fresh draft at the next version and every other line's draft is updated under the same id (2026-09-10, `haiku/v0.1.0` then `curry/v0.0.1`).
- GitHub's **Latest** badge lands on whichever line was published last; glyph never sets `make_latest`.
