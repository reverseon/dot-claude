#!/usr/bin/env bash
INPUT=$(cat)
COMMAND=$(echo "$INPUT" | jq -r '.tool_input.command // empty')

typeset -a BLOCKED_PATTERNS=(
  'git push'
  'git push --force'
  'git reset --hard'
  'git rebase'
  'git commit'
  'rm -rf'
  'npm publish'
  'npm install -g'
  'curl.*\|.*sh'      # curl pipe sh exfil
  '\$\(.*git push'    # subshell bypass trap
  '`.*git push'       # backtick bypass trap
  '>.*.env'           # write to .env
)

for pattern in "${BLOCKED_PATTERNS[@]}"; do
  if echo "$COMMAND" | grep -qE "$pattern"; then
    echo "🚫 Blocked: matches '$pattern'" >&2
    exit 2
  fi
done

exit 0