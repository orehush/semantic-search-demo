from fastapi.testclient import TestClient
from main import app

client = TestClient(app)


def test_get_synonyms_valid_phrase():
    response = client.post(
        "/semantic-tools/synonyms/", json={"phrase": "happy"})
    assert response.status_code == 200
    json_response = response.json()
    assert "synonyms" in json_response
    assert len(json_response["synonyms"]) > 0


def test_get_synonyms_empty_phrase():
    response = client.post("/semantic-tools/synonyms/", json={"phrase": " "})
    assert response.status_code == 400
    assert response.json() == {"detail": "Phrase must not be empty"}


def test_get_synonyms_nonexistent_word():
    response = client.post(
        "/semantic-tools/synonyms/", json={"phrase": "qwertyuiop"})
    assert response.status_code == 200
    json_response = response.json()
    assert "synonyms" in json_response
    assert len(json_response["synonyms"]) == 0
