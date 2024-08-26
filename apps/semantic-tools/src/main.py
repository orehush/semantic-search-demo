from fastapi import FastAPI, HTTPException
from pydantic import BaseModel, Field
from synonyms.generator import generate_synonyms

app = FastAPI()


class GenerateSynonymsRequest(BaseModel):
    phrase: str = Field(..., min_length=1,
                        description="The word or phrase to find synonyms for")


@app.post("/semantic-tools/synonyms/")
async def get_synonyms(request: GenerateSynonymsRequest):
    phrase = request.phrase.strip()
    if not phrase:
        raise HTTPException(status_code=400, detail="Phrase must not be empty")

    synonyms = generate_synonyms(phrase)
    return {"synonyms": synonyms}
