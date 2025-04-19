from praisonaiagents import Agent, MCP

search_agent = Agent(
        instructions="""You help book apartement on Airbnb.""",
        llm="ollama/qwen2.5",
        tools=MCP("npx -y @openbnb/mcp-server-airbnb --ignore-robots-txt")
)

while True:
    query = input("?\n")
    search_agent.start(query)
