from praisonaiagents import Agent, MCP

search_agent = Agent(
        instructions="""You are a shell command operator that can execute atmos commands.""",
        llm="ollama/qwen2.5",
        tools=MCP("npx -y shell-command-mcp",
                  env={
                  "ALLOWED_COMMANDS": "atmos"
              })
)
while True:
    query = input("?\n")
    search_agent.start(query)
