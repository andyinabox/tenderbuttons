# Tender Buttons

http://tenderbuttons.click

 - [About this project](ABOUT.md)
 - [Original text](corpus/pg15396.txt) from project Gutenberg and my [modified version](corpus/tb.txt) used in this project.
 - See the [Makefile](Makefile) for hints if you want to try and run this project yourself.

## Releasing

Releases are cut with `make release VERSION=vX.Y.Z`, starting at `v1.0.0`. The
[`sh/release.sh`](sh/release.sh) script:

1. Validates the version (must match `vX.Y.Z`) and that the docker daemon is running.
2. Refuses to proceed if the working tree is dirty or the tag already exists.
3. Creates an annotated git tag.
4. Builds the image as `code.andydayton.com/andy/tenderbuttons:vX.Y.Z` (also tagged `:latest`).
5. Pushes both image tags to the registry.
6. Pushes the git tag to `origin`.

This replaces the previous Docker Hub workflow — images now live on the
self-hosted registry at `code.andydayton.com`. Before running a release, make
sure you're logged in: `docker login code.andydayton.com`.

For local development, `make docker-build` and `make docker-run` still work
and tag the image as `:dev` so it can't be confused with a real release.
