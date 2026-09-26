# glyph-monorepo-test

Permanent live-fire harness for [glyph](https://github.com/akira-toriyama/glyph)'s
monorepo support: several independently versioned Go modules in one repository,
each with its own tag line and its own rolling draft release.

The property the harness exists to prove: a pull request that touches `haiku/`
moves **haiku's** version and nothing else, and a squash merge never loses which
of the pull's commits touched which module.

| module | tag line | what its history shows |
|---|---|---|
| `haiku/` | `haiku/vX.Y.Z` | `%` is the only door to 1.0.0, and `!` takes the major once a line is 1.x — two published majors (`haiku/v1.0.0`, `haiku/v2.0.0`), a third standing as the `haiku/v3.0.0` draft |
| `curry/` | `curry/vX.Y.Z` | a line that grew on `^` and `~` alone before it called 1.0 |
| `travel/` | `travel/vX.Y.Z` | `!` on a 0.x line steps the **minor**: breaking things early cannot claim a stable major |
| `travel/onsen/` | `travel/onsen/vX.Y.Z` | **nested** in `travel/`: a change here moves onsen alone — the longest declared prefix wins. Both lines promoted to 1.0.0 in the same merge, each on its own commits; a `!` since stands as the `travel/onsen/v2.0.0` draft, on the `/v2` path |
| `camp/` | `camp/vX.Y.Z` | a line still in 0.x four rounds in, breaking changes included |

Tags follow the Go multi-module convention (`<dir>/vX.Y.Z`), and a module is
fetchable with `go get` at a tag only when its `go.mod` path carries that tag's
major. glyph derives the tag and never reads a `go.mod`, so moving the path is
the author's commit, landed before the draft is published: `haiku/go.mod` says
`/v3` for the standing `haiku/v3.0.0` draft and `travel/onsen/go.mod` says `/v2`
for `travel/onsen/v2.0.0`, while haiku's published v2.x tags were cut on the
suffix-less path and `go get` refuses them (measured 2026-09-26: "module path
must match major version"). `go get` takes the tag's version, not the tag —
`…/curry@v1.1.0`, never `…/curry@curry/v1.1.0`, which the proxy refuses as a
disallowed version string.
Nothing is built and nothing is uploaded: publishing a line's draft is what
cuts that line's tag, and the module at that tag is the whole artifact.

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
- four rounds of twenty commits, three of them published (2026-09-11): those three rounds published 15 releases, 17 across the five lines with `haiku/v0.1.0` and `curry/v0.0.1` from the day before — first published `haiku/v0.1.0`, last `camp/v0.3.0`, highest version `haiku/v2.1.0` — and one standing draft per line on top. Each round was a single squash merge, and every line stepped on its own commits alone.
- a hand region written **above** the sentinel survives a rewrite: written into all five standing drafts, then read back unchanged after `e2e.yml`'s drafts tier had rewritten every one of them (2026-09-11, run 34501469251).
- the published floor is **per line**: at the frozen coordinate `release --dry-run --since-tag=camp/v0.0.0` refuses at exit 4 naming camp's own latest published release, never another line's. That is now the walk tier's arm (b2).
- `e2e.yml`'s drafts tier runs `release` with no `--footer-file`, so a dispatch leaves every draft without the `install-notes` block until the next push to `main` writes it back (2026-09-11). Publish after a push, never straight after a dispatch.
- a rename **across two lines** moves both, over the API's own file listing: #22 moved `portion.go` from `curry/` to `camp/` in one commit, and the merge preview put that commit in both tables (2026-09-11). `git diff-tree --no-renames` and `previous_filename` agree.
- path boundaries are whole path segments, not string prefixes (2026-09-11, `bump --range` on a local clone): a root file named `camp.md` and a root directory `travelling/` are shared-only and refused, `travel/onsenx/` belongs to `travel`, and `travel/onsen/` belongs to `onsen`.
- the four attribution answers, measured one at a time on a local clone (2026-09-11): an empty commit carrying `^` is refused, a shared-only `^` is refused, a shared-only `^` whose scope names `haiku` moves haiku alone, and a `(haiku)` scope on a commit touching only `curry/` is refused as a contradiction.
- a pull that touches **no** declared line renders its own headline rather than an empty table: "Merging this PR moves nothing — its 1 commit(s) touch no declared package" (2026-09-11, #23, closed after the reading).
- the reusable still refuses an artifact input here: `artifact-probe.yml` went red at "Refuse an artifact input on a packages repository" against `release.yml@v3.3.0` (2026-09-11, run 34561643491).
- a **sigil-less bot subject breaks this harness**, and the fix is upstream (2026-09-13). `fleet-sync`
  wrote `:wrench:(fleet) sync 8 standard files in one push` (948d1a6) and `release` refused the whole
  range at exit 3 (run 34745527599) — this repository has no v1-acceptance window, so it is one of the
  two in the fleet that can report the defect at all. Producers fixed in `akira-toriyama/.github#226`;
  the already-merged commit was amended to carry `=` (948d1a6 -> 614d3ad, tree and parent unchanged,
  author preserved), after which `release --dry-run` exits 0 and all five lines resolve.
- `commit-lint`'s push arm **refuses rather than guesses across a force push**: after that amend it
  answered `event.before is not an ancestor of event.after (a force push, or a base this clone cannot
  see)` at exit 1 (run 34749951009). That is the designed answer, not a defect — the next ordinary push
  to `main` judges normally, which is what the commit adding these two lines demonstrates.
