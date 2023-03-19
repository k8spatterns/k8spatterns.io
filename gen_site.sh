#!/bin/bash

set -o pipefail
set -eu

set -x

git branch -l
git remote

if [ -n "$(git status -s)" ]; then
  echo "The working directory is dirty. Please commit any pending changes."
  git status
  exit 1;
fi

echo "Deleting old publication"
rm -rf public
mkdir public
git worktree prune
rm -rf .git/worktrees/public/

echo "Checking out gh-pages branch into public"
git worktree add -B gh-pages public origin/gh-pages

echo "Removing existing files"
rm -rf public/*

echo "Generating site"
hugo -v

echo "Copying over examples"
cp -r examples/* public

# Commit & push everything
cd public
if [ -n "$(git status -s)" ] ; then
  echo "Updating gh-pages branch"
  git config credential.helper 'cache --timeout=120'
  git config user.email "gh@k8spatterns.io"
  git config user.name "k8spatterns-ci"
  git add --all && git commit -m "Publishing to gh-pages [ci skip]" && git push --force origin gh-pages
fi
