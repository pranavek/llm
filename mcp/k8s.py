from praisonaiagents import Agent, MCP

search_agent = Agent(
        instructions="""You are a Kubernetes operator.""",
        llm="ollama/qwen2.5",
        tools=MCP("npx -y mcp-server-kubernetes --ignore-robots-txt")
)
while True:
    query = input("?\n")
    search_agent.start(query)
