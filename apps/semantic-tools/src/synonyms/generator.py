import nltk
from nltk.corpus import wordnet

# Ensure required resources are downloaded
nltk.download('wordnet')


def generate_synonyms(phrase: str):
    synonyms = []

    for syn in wordnet.synsets(phrase):
        for lemma in syn.lemmas():
            synonyms.append({"synonym": lemma.name(), "score": lemma.count()})

    # Sort the list of synonyms by score in descending order
    synonyms = sorted(synonyms, key=lambda x: x['score'], reverse=True)

    return synonyms
