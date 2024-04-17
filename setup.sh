#!/bin/bash

# Setup git hooks
cp git_hooks/* .git/hooks/
chmod +x .git/hooks/*
echo "Git hooks installed."
