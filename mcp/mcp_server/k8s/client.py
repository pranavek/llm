import asyncio
from pydantic_ai import Agent
from pydantic_ai.models.openai import OpenAIModel
from pydantic_ai.providers.openai import OpenAIProvider
from pydantic_ai.mcp import MCPServerStdio

server = MCPServerStdio(  
    'go',
    args=[
        'run',
        'main.go',
    ]
)

ollama_model = OpenAIModel(
    model_name='qwen2.5',  
    provider=OpenAIProvider(base_url='http://localhost:11434/v1'), 
)
agent = Agent(model=ollama_model, mcp_servers=[server])


async def main():
    async with agent.run_mcp_servers():
        while True:
            prompt = input("Enter a prompt (or press Enter to exit): ")
            if not prompt:
                break
            result = await agent.run(prompt)
            print(result.output+"\n")

if __name__ == "__main__":
    asyncio.run(main())
