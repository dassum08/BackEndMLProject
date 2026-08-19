def test_post_data(client):
    response = client.post("/getdata")

    assert response.status_code == 200