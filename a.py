import requests
import gzip, csv
from urllib.parse import quote

graph = []

# initialy ask user for random picks from 5 lists
# then get common related of all 5 lists
# then shuffle and pick 5 random from that list
# get more like for each of those 5
# shuffle and pick 5 random from that list
# show user 5 random from that list

#https://api.themoviedb.org/3/search/tv?api_key=d56e51fb77b081a9cb5192eaaa7823ad&query=Avatar&callback=jQuery20303312176821254378_1703523248502&_=1703523248503

def get_keyword(keyword: str):
    response = requests.get(
        "https://api.themoviedb.org/3/search/keyword?query="+quote(keyword)+"&page=1&api_key=d56e51fb77b081a9cb5192eaaa7823ad"
    )
    response = response.text
    with open('a.json', 'wb') as f:
        f.write(response.encode('utf-8'))
        

def unzip_gz_and_Save(filename):
    with gzip.open(filename, 'rb') as f:
        file_content = f.read()
    with open(filename[:-3], 'wb') as f:
        f.write(file_content)

unzip_gz_and_Save('title.basics.tsv.gz')