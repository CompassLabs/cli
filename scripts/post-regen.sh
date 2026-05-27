#!/usr/bin/env bash
# Re-apply hand-written fixes on top of a fresh Speakeasy regen.
#
# WHY THIS EXISTS
# ---------------
# `speakeasy run -t compass-cli` rewrites everything under cli-sdk/internal/
# and cli-sdk/cmd/ from templates + the OpenAPI spec, wiping any direct edits
# to those files. The release pipeline runs that command in CI
# (.github/workflows/generate_combined_spec_and_sdks.yaml, the `generate-cli`
# job) and ships the artifact straight to CompassLabs/cli, so any mono-local
# patch to a generated file dies before it ever reaches a user.
#
# Our fixes live as unified diffs in `cli-sdk/scripts/post-regen/*.patch`.
# This script applies them all, sorted lexicographically (NN-name.patch), so
# the order is reviewable from `ls`. Each patch must end with `git apply`
# reporting success — any failure aborts the script (and therefore the
# release) with a non-zero exit code.
#
# WHEN A PATCH STOPS APPLYING
# ---------------------------
# If `git apply` rejects some `NN-foo.patch`, that's the signal that
# Speakeasy's template changed how it emits the surrounding code. The
# workflow uploads a `cli-sdk-pristine` artifact right before this step
# runs, so the recovery path is:
#
#   gh run download <RUN-ID> -n cli-sdk-pristine -D /tmp/pristine
#   # /tmp/pristine now holds the freshly-regenerated cli-sdk that the
#   # failing patch was generated against — overwrite the problem file
#   # in your working tree with /tmp/pristine/internal/.../foo.go,
#   # re-apply your intent by hand, then capture a fresh patch:
#   git diff -- cli-sdk/path/to/foo.go > cli-sdk/scripts/post-regen/NN-foo.patch
#
# DO NOT make the script "tolerant" of partial failures — silent skipping
# is exactly the regression mode that bit commit 31c3769d8.
#
# USAGE
# -----
#   bash cli-sdk/scripts/post-regen.sh                  # run from repo root
#   ( cd cli-sdk && bash scripts/post-regen.sh )        # or from cli-sdk
#
# The script normalizes its working directory to the cli-sdk repo root so
# the patches' `a/cli-sdk/...` paths resolve correctly.

set -euo pipefail

# --- locate cli-sdk root --------------------------------------------------
# Patches were generated relative to the mono root (`a/cli-sdk/internal/...`).
# git apply -p1 strips the leading component, so we must run from the
# directory that *contains* cli-sdk/ — i.e. the mono root.

script_dir=$(cd "$(dirname "$0")" && pwd)
cli_sdk_dir=$(dirname "$script_dir")              # cli-sdk/
mono_root=$(dirname "$cli_sdk_dir")               # mono/
patch_dir="$script_dir/post-regen"

cd "$mono_root"

if [ ! -d "$patch_dir" ]; then
  echo "post-regen: no patch directory at $patch_dir — nothing to do" >&2
  exit 0
fi

# --- collect patches in order --------------------------------------------
shopt -s nullglob
patches=("$patch_dir"/*.patch)
shopt -u nullglob

if [ "${#patches[@]}" -eq 0 ]; then
  echo "post-regen: no .patch files in $patch_dir — nothing to do" >&2
  exit 0
fi

echo "post-regen: applying ${#patches[@]} patch(es) from $patch_dir"

# --- apply --------------------------------------------------------------
# We use plain `git apply` (no `--3way`) so the patch lookup only touches
# the working-tree file, not git's object store. The CI workflow checks
# out with the default `fetch-depth: 1` (shallow), which means the
# pre-image blobs referenced by our patches' `index OLD..NEW` headers
# aren't present locally — `git apply --3way` then errors out with
# "does not match index" even when the on-disk file matches the patch
# base byte-for-byte. Plain `git apply` does pure context matching and
# avoids that trap. `--whitespace=nowarn` silences trailing-newline
# noise that Speakeasy's templates sometimes introduce.
#
# On hard conflict we exit immediately — half-applying a patch set ships
# a half-broken binary, which is exactly the regression mode that bit
# commit 31c3769d8. When that happens, grab the `cli-sdk-pristine`
# artifact uploaded by the workflow step right before this one and
# regenerate the failing patch (recipe in cli-sdk/PRESERVED_EDITS.md).

for p in "${patches[@]}"; do
  name=$(basename "$p")
  if ! git apply --whitespace=nowarn "$p"; then
    echo "  ✗ $name — patch no longer applies; regenerate it (see header of post-regen.sh)" >&2
    exit 1
  fi
  echo "  ✓ $name"
done

# --- sanity build --------------------------------------------------------
# Same guardrail the existing risk-import sed has: if anything we did
# breaks the build, fail loudly here rather than shipping a binary that
# can't compile.
( cd "$cli_sdk_dir" && go build ./cmd/compass >/dev/null )

echo "post-regen: all patches applied; build passes"
