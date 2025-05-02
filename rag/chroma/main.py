# https://www.youtube.com/watch?v=E4l91XKQSgw

from langchain_ollama.llms import OllamaLLM
from langchain_core.prompts import ChatPromptTemplate
from vector import retriver


model = OllamaLLM(model="qwen2.5")

template = """
You are an expert in answering questions about the tv shows.

Here are some relevant reviews: {reviews}

Here is the question to answer: {question}  
"""

prompt = ChatPromptTemplate.from_template(template)
chain = prompt | model

while True:
    question = input("Enter your question: ")
    if question == "q":
        break
    
    reviews = retriver.invoke(question)
    result = chain.invoke({"reviews": reviews,"question": question,})
    print("-" * 30)
    print(result)
    print("+" * 30)
    
    
    """
    provide stats about game of throne
     list top rated tv shows on action genre
    what is the best tv show to watch that came after year 2010
    """