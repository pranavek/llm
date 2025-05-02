import asyncio
from pydantic_ai import Agent
from pydantic_ai.models.openai import OpenAIModel
from pydantic_ai.providers.openai import OpenAIProvider
from pydantic_ai.mcp import MCPServerStdio

assistant_server = MCPServerStdio(  
    'go',
    args=[
        'run',
        'main.go',
    ]
)

local_rag_server = MCPServerStdio(
    'uvx',
    args=[
        "--from",
        "git+https://github.com/nkapila6/mcp-local-rag",
        "mcp-local-rag"
    ]
)

ollama_model = OpenAIModel(
    model_name='qwen2.5',  
    provider=OpenAIProvider(base_url='http://localhost:11434/v1'), 
)
agent = Agent(
    model=ollama_model, 
    mcp_servers=[assistant_server, local_rag_server],
    instuctions="""
    You are a helpful assistant. You can use the DuckDuckGo search engine to find information.
    """,
)


async def main():
    async with agent.run_mcp_servers():
        while True:
            prompt = input("Enter a prompt (or press Enter to exit): ")
            if not prompt:
                break
            result = await agent.run(prompt)
            print(result.output+"\n\n")

if __name__ == "__main__":
    asyncio.run(main())
