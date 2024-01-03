from curl_cffi import requests
from requests import get as g


chatgpt_backend_url = "https://wall.alphacoders.com/"

def login_chatgpt(email, password) -> str:
    csrf = requests.get("https://dodi-repacks.site",impersonate="chrome110").status_code

    print(csrf)


def x():
    r  = requests.get(chatgpt_backend_url, impersonate="chrome110")

    print(r.status_code)

#from concurrent.futures import ThreadPoolExecutor
#
#with ThreadPoolExecutor(max_workers=100) as executor:
 #   for i in range(1000):
  #      executor.submit(login_chatgpt, "chrome110", "chrome110")
print(g(chatgpt_backend_url).status_code)
print(requests.get(chatgpt_backend_url, impersonate="chrome110").status_code)