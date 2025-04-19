from praisonaiagents import Agent, MCP

search_agent = Agent(
        instructions="""You help me to search """,
        llm="ollama/qwen2.5",
        tools=MCP("npx -y duckduckgo-mcp-server --ignore-robots-txt")
)
while True:
    search_query = input("What would you like to search for?\n")
    search_agent.start(search_query)