from collections import Counter
from requests import get
from urllib.parse import quote
import time
import random


def search_series(name: str):
    req = get(
        "https://tastedive.com/api/search?query={}&take=9&page=1&types=urn:entity:artist,urn:entity:movie,urn:entity:tv_show".format(
            quote(name)
        ),
        headers={
            "authority": "tastedive.com",
            "accept": "application/json, text/plain, */*",
            "referer": "https://tastedive.com/shows",
            "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko)",
        },
    )

    results = []
    for result in req.json()["results"]:
        resultW = {
                "name": result["name"],
                "id": result["entity_id"],
                "type": result["types"][0],
            }
        if result.get("properties", {}).get("external", {}).get("imdb", {}).get("user_rating", None) is not None:
            resultW["imdb_rating"] = result["properties"]["external"]["imdb"]["user_rating"]
        
        results.append(resultW)
    return results

def get_similar_series(id: str, type: str):
    req = get("https://tastedive.com/api/getRecsByCategory?page={}&entityId={}&category={}".format
              (
                  random.randint(1, 5),
                    id,
                    type.split(":")[2].replace("tv_show", "shows").replace("movie", "movies")
              ), headers={
            "authority": "tastedive.com",
            "accept": "application/json, text/plain, */*",
            "referer": "https://tastedive.com/shows",
            "user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko)",
        })
    
    result = []
    for res in req.json():
        result.append({
            "name": res["entityName"],
            "id": res["id"],
            "type": res["entityTypeId"],
            "image": res["image"],
        })

    return result


def get_most_favourable_genres(user_genre_list: list):
    gen = Counter(user_genre_list)
    most_common = gen.most_common(3)
    return [genre[0] for genre in most_common]


def get_top_series(series: list, genre_list: list):
    series = [series for series in series if series["rating"] > 6]

    series = [
        series
        for series in series
        if any(genre in series["genres"] for genre in genre_list)
    ]
    print(series)
    series = sorted(series, key=lambda k: k["rating"], reverse=True)
    if len(series) > 5:
        series = series[:5]
    return series


class User:
    def __init__(self, id):
        self.id = id
        self.genre_list = []
        self.series_list = []
        self.favorite_series = []
        self.watched_series = []
        self.watchlist = []

    def add_genre(self, genre):
        self.genre_list.append(genre)

    def add_genres(self, genre_list):
        self.genre_list.extend(genre_list)

    def add_series(self, series):
        self.series_list.append(series)

    def add_series_list(self, series_list):
        self.series_list.extend(series_list)

    def get_recommendations(self):
        genre_list = get_most_favourable_genres(self.genre_list)
        series_list = get_top_series(self.series_list, genre_list)
        return series_list

    def fetch_series_bulk(self):
        import concurrent.futures
        common_cluster = []

        with concurrent.futures.ThreadPoolExecutor(max_workers=10 if len(self.series_list) > 10 else len(self.series_list)) as executor:
            futures = []
            for series in self.series_list:
                futures.append(executor.submit(get_similar_series, series["id"], series["type"]))

            for future in concurrent.futures.as_completed(futures):
                series_list = future.result()
                for series in series_list:
                    if series not in common_cluster:
                        common_cluster.append([series, 1])
                    else:
                        common_cluster[common_cluster.index(series)][1] += 1

                    self.add_series(series)
        
        common_cluster = sorted(common_cluster, key=lambda k: k[1], reverse=True)
        print(common_cluster[:5])

    def fetch_series(self):
        pass


def ask_reccomendation(user):
    print("What genres do you like? (separated by comma)")
    genre_list = input().split(",")
    user.add_genres(genre_list)
    print("What series do you like? (separated by comma)")
    series_list = input().split(",")
    for series in series_list:
        search = search_series(series)
        if len(search) > 0:
            user.add_series(search[0])
        else:
            print("No series found with name {}".format(series))
    user.fetch_series_bulk()


u = User(1)
ask_reccomendation(u)

a = time.time()
print("Fetching series/movies and their recommendations...")
u.fetch_series_bulk()
print(time.time() - a)

with open("a.json", "w") as f:
    import json
    json.dump(u.series_list, f)

print("Here are your recommendations:")
i = 0
for series in u.series_list:
    i += 1
    print("{}. {} ({})".format(i, series["name"], series["type"]))
    print("1. Add to watchlist ->")
    print("2. Add to favorites ->")
    print("3. Add to watched ->")
    print("4. Skip ->")
    choice = input()
    if choice == "1":
        u.watchlist.append(series)
    elif choice == "2":
        u.favorite_series.append(series)
    elif choice == "3":
        u.watched_series.append(series)
    elif choice == "4":
        continue
    else:
        print("Invalid choice")
        continue