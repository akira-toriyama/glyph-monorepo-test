# glyph-monorepo-test

Permanent live-fire harness for [glyph](https://github.com/akira-toriyama/glyph)'s
monorepo support: several independently versioned Go modules in one repository,
each with its own tag line and its own rolling draft release.

The property the harness exists to prove: a pull request that touches `haiku/`
moves **haiku's** version and nothing else, and a squash merge never loses which
of the pull's commits touched which module.

| module | tag line | what its history shows |
|---|---|---|
| `haiku/` | `haiku/vX.Y.Z` | `%` is the only door to 1.0.0, and `!` takes the major once a line is 1.x — three majors so far |
| `curry/` | `curry/vX.Y.Z` | a line that grew on `^` and `~` alone before it called 1.0 |
| `travel/` | `travel/vX.Y.Z` | `!` on a 0.x line steps the **minor**: breaking things early cannot claim a stable major |
| `travel/onsen/` | `travel/onsen/vX.Y.Z` | **nested** in `travel/`: a change here moves onsen alone — the longest declared prefix wins. Both lines promoted to 1.0.0 in the same merge, each on its own commits |
| `camp/` | `camp/vX.Y.Z` | a line still in 0.x four rounds in, breaking changes included |

Tags follow the Go multi-module convention (`<dir>/vX.Y.Z`), so each module is
also fetchable with `go get` at its own version. Nothing is built and nothing is
uploaded: publishing a line's draft is what cuts that line's tag, and the module
at that tag is the whole artifact.

A release page here carries three parts, and only the middle one is glyph's:
prose written by hand **above** the sentinel comment, the generated sections
below it (one per sigil class, in `[[note.sections]]` order), and the caller's
`install-notes` footer under a `---`. glyph rewrites everything below the
sentinel on every push and never touches what is above it.

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
- the **tag landscape** as of that commit. glyph resolves a line's base with `git tag --list` and never consults reachability, so a tag cut afterwards becomes the base of a walk that starts before it and every literal collapses to `none`. The walk tier therefore clones and deletes the tags no descendant could have shown it, and asserts the surviving list as a literal (measured 2026-09-11: with `haiku/v0.2.0` and `travel/v0.1.0` planted on a descendant, arm (a) answered `haiku:v0.2.0:none|travel:v0.1.0:none`).

Measured by hand rather than in e2e, because each consumes a tag or a version:

- publish one line's draft, merge again: the published line gets a fresh draft at the next version and every other line's draft is updated under the same id (2026-09-10, `haiku/v0.1.0` then `curry/v0.0.1`).
- GitHub's **Latest** badge lands on whichever line was published last; glyph never sets `make_latest`.
- four rounds of twenty commits, three of them published (2026-09-11): 17 published releases across the five lines, from `camp/v0.1.0` to `haiku/v3.0.0`, and one standing draft per line on top. Each round was a single squash merge, and every line stepped on its own commits alone.
- a hand region written **above** the sentinel survives a rewrite: written into all five standing drafts, then read back unchanged after `e2e.yml`'s drafts tier had rewritten every one of them (2026-09-11, run 34501469251).
- the published floor is **per line**: at the frozen coordinate `release --dry-run --since-tag=camp/v0.0.0` refuses at exit 4 naming camp's own latest published release, never another line's. That is now the walk tier's arm (b2).
- `e2e.yml`'s drafts tier runs `release` with no `--footer-file`, so a dispatch leaves every draft without the `install-notes` block until the next push to `main` writes it back (2026-09-11). Publish after a push, never straight after a dispatch.

A note that touches no version line.
