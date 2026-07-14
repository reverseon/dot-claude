#!/bin/bash
# MCP Server Setup Script
# Registers MCP servers with Claude Code (user scope)

# Add AWS Knowledge MCP server (HTTP)
claude mcp add --transport http aws-knowledge https://knowledge-mcp.global.api.aws -s user

# MCP Docs

claude mcp remove mcp-docs -s user || true && claude mcp add-json mcp-docs \
  '{
    "type": "stdio",
    "command": "uvx",
    "args": [
      "--from", "mcpdoc", "mcpdoc",
      "--urls",
      "Claude:https://claude.com/docs/llms-full.txt",
      "Docker:https://docs.docker.com/llms-full.txt",
      "GitHub:https://docs.github.com/llms.txt",
      "Ollama:https://docs.ollama.com/llms-full.txt"
    ]
  }' \
  -s user

# Tools: Pathfinder
claude mcp add pathfinder-docs --transport http https://mcp.pathfinder.copilotkit.dev/mcp -s user

# Pathfinder Actual
claude mcp add local-knowledgebase --transport http http://localhost:3001/mcp -s user

echo "MCP servers configured successfully!"
claude mcp list
