from langchain_ollama import OllamaEmbeddings
from langchain_chroma import Chroma
from langchain_core.documents import Document
import os
import pandas as pd

df = pd.read_csv("./data_set/imdb_top_5000_tv_shows.csv")
embeddings = OllamaEmbeddings(model="nomic-embed-text")

db_location = "./chroma_db"
add_doucuments = not os.path.exists(db_location)

if add_doucuments:
    documents = []
    ids = []
    
    for i, row in df.iterrows():
        document = Document(
            page_content=row["primaryTitle"] +" "+ str(row["averageRating"]) ,
            metadata={
                "genres": row["genres"],
                "startYear": row["startYear"],
            },
            id=str(i)
            )
        ids.append(str(i))
        documents.append(document)

vector_store = Chroma(
    collection_name="my_collection",
    persist_directory=db_location,
    embedding_function=embeddings,
    )

if add_doucuments:
    vector_store.add_documents(documents, ids=ids)

retriver = vector_store.as_retriever(
    search_kwargs={"k": 5}
    )

