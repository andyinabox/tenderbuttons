#!/usr/bin/env bash
set -euo pipefail

REGISTRY="code.andydayton.com"
IMAGE="${REGISTRY}/andy/tenderbuttons"

usage() {
	echo "usage: $0 vX.Y.Z" >&2
	exit 1
}

[[ $# -eq 1 ]] || usage

VERSION="$1"

if [[ ! "$VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
	echo "error: version must match vX.Y.Z (got: $VERSION)" >&2
	exit 1
fi

if ! docker info >/dev/null 2>&1; then
	echo "error: docker daemon is not running" >&2
	exit 1
fi

if [[ -n "$(git status --porcelain)" ]]; then
	echo "error: working tree is dirty; commit or stash changes before releasing" >&2
	exit 1
fi

if git rev-parse -q --verify "refs/tags/$VERSION" >/dev/null; then
	echo "error: tag $VERSION already exists locally" >&2
	exit 1
fi

if git ls-remote --exit-code --tags origin "refs/tags/$VERSION" >/dev/null 2>&1; then
	echo "error: tag $VERSION already exists on origin" >&2
	exit 1
fi

echo "==> Tagging $VERSION"
git tag -a "$VERSION" -m "Release $VERSION"

echo "==> Building $IMAGE:$VERSION (also tagged :latest)"
docker build -t "$IMAGE:$VERSION" -t "$IMAGE:latest" .

echo "==> Pushing $IMAGE:$VERSION"
docker push "$IMAGE:$VERSION"

echo "==> Pushing $IMAGE:latest"
docker push "$IMAGE:latest"

echo "==> Pushing tag $VERSION to origin"
git push origin "$VERSION"

echo "==> Released $VERSION"
