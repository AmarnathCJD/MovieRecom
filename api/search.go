package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
)

type Title struct {
	Name       string  `json:"name,omitempty"`
	Id         string  `json:"id,omitempty"`
	Type       string  `json:"type,omitempty"`
	ImdbRating float64 `json:"imdb_rating,omitempty"`
	Poster     string  `json:"poster,omitempty"`
}

func SearchSeriesHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")
	if q == "" {
		http.Error(w, "Bad Request, missing query", http.StatusBadRequest)
		return
	}

	results, err := searchSeries(q)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(results); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func searchSeries(query string) ([]Title, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://tastedive.com/api/search?query=%s&take=50&page=1&types=urn:entity:artist,urn:entity:movie,urn:entity:tv_show,urn:entity:videogame,urn:entity:person", url.QueryEscape(query)), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("authority", "tastedive.com")
	req.Header.Set("accept", "application/json, text/plain, *")
	req.Header.Set("referer", "https://tastedive.com/shows")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var resultsRaw map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&resultsRaw); err != nil {
		return nil, err
	}

	results := []Title{}
	for _, _result := range resultsRaw["results"].([]interface{}) {
		var result = _result.(map[string]interface{})
		resultW := Title{
			Name: result["name"].(string),
			Id:   result["entity_id"].(string),
			Type: result["types"].([]interface{})[0].(string),
		}

		if properties, ok := result["properties"]; ok {
			if external, ok := properties.(map[string]interface{})["external"]; ok {
				if imdb, ok := external.(map[string]interface{})["imdb"]; ok {
					if user_rating, ok := imdb.(map[string]interface{})["user_rating"]; ok {
						resultW.ImdbRating = user_rating.(float64)
					}
				}
			}

			if image, ok := properties.(map[string]interface{})["image"]; ok {
				if url, ok := image.(map[string]interface{})["url"]; ok {
					resultW.Poster = strings.ReplaceAll(url.(string), "420x", "840x")
				}
			}
		}

		results = append(results, resultW)
	}

	return results, nil
}

func getSimilar(Id string, _type string) ([]Title, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://tastedive.com/api/getRecsByCategory?page=%d&entityId=%s&category=%s", rand.Intn(5)+1, Id, strings.ReplaceAll(strings.ReplaceAll(_type, "tv_show", "shows"), "movie", "movies")), nil)
	req.Header.Set("authority", "tastedive.com")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	var resultRaw []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&resultRaw); err != nil {
		return nil, err
	}

	for _, resultW := range resultRaw {
		var _ = resultW.(map[string]interface{})

	}

	return nil, nil // TODO: Complete This
}

/*
def get_similar_series(id: str, type: str):
    req = get("https://tastedive.com/api/getRecsByCategory?page={}&entityId={}&category={}".format
              (

                  random.randint(1, 5),
                    id,
                    type.split(":")[2].replace("tv_show", "shows").replace("movie", "movies")
              ), headers={
            "authority": "tastedive.com",
            "accept": "application/json, text/plain, *",
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
*/
