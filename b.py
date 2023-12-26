import requests

headers = {
    'authority': 'tastedive.com',
    'accept': 'application/json, text/plain, */*',
    'referer': 'https://tastedive.com/shows',
    'user-agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36',
}

response = requests.get(
    'https://tastedive.com/api/search?query=Strager%20things&take=9&page=1&types=urn:entity:artist,urn:entity:movie,urn:entity:tv_show',
    headers=headers,
)

with open('a.json', 'wb') as f:
    f.write(response.text.encode('utf-8'))