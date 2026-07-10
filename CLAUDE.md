## Golden Rules

### Discovery vs Editing
- Treat all work as being in discovery phase by default.
- During discovery, you MAY read, inspect, analyze, summarize, and propose changes.
- During discovery, you MUST NOT make any changes to files, content, configuration, or state unless I explicitly tell you to make the changes.
- Do not interpret analysis, suggestions, demonstrations, drafts, or examples as permission to apply edits.
- Before making any change, state the proposed change and wait for explicit approval.
- Only enter editing phase when I clearly and specifically instruct you to apply changes.
- Approval is limited to the exact scope I specify; do not make additional edits, refactors, formatting changes, or cleanup unless I explicitly ask for them.

### Script Execution
- You MAY create temporary scripts via `cat` or heredoc **only** for read/parse/inspect purposes
  (e.g., extracting data, grepping output, formatting a report).
- Temporary scripts MUST be pure readers — no `write`, `sed -i`, `awk` with output redirect,
  or any operation that mutates files or state.

### File Changes
- ALL file modifications MUST use the Claude Code Edit tool (`str_replace`, `create`, `insert`).
- NEVER use shell to write code changes; Use Edit tool only.
- This applies to ALL file types: source code, configs, dotfiles, markdown.

### Comment, Don't Delete
- When code needs to be removed, comment it out instead of deleting it.
- Preserves history, allows easy reversal, and makes diffs clearer.
- Exception: Obvious boilerplate, auto-generated code, or lines you explicitly approve for deletion.
- If unsure, ask first before deleting anything.

### Verify External System Data

For ANY values/details from external systems that could be outdated, environment-specific, or security-critical:
- API endpoints, service URLs, configuration formats
- AWS service details, IAM policies, API responses
- Third-party service configurations (CircleCI, GitHub, etc.)
- Command syntax, CLI parameters, authentication methods
- Certificate/cryptographic values, tokens, keys, thumbprints

**Rule: If it came from training data and the system could have changed it, verify against the actual current source first.**

When in doubt:
- Check official documentation
- Run the actual command to verify output
- Query the live API/service
- Ask before proceeding if unverifiable

Training data has a knowledge cutoff. External systems don't. For anything that could reasonably change or be environment-specific, verify rather than rely on training knowledge.

### Prefer MCP Servers Over General Tools

When an MCP server is available for a system (e.g., mcp-docs for GitHub/Docker/Claude docs), use it instead of general tools:
- Use `mcp__mcp-docs__*` tools for official documentation (GitHub, Docker, Claude)
- Use `WebFetch` only when no specialized MCP server exists
- Check `<system-reminder>` deferred tools and available skills before defaulting to general tools
- MCP servers provide better coverage, structured access, and fresher documentation than general web fetching