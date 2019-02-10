#!/bin/sh

if [[ -n "$(git status -s)" ]]; then
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

# Commit & push everything
cd public
if [[ -n "$(git status -s)" ]] ; then
  echo "Updating gh-pages branch"
  git add --all && git commit -m "[ci skip] Publishing to gh-pages" && git push --force
fi
