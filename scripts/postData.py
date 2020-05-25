import requests
url = 'http://localhost:8082/cache/set'
payload = {'ID':'1','Value':'first value'}

#r = requests.get(url)

#r = requests.get(url,params=payload)

r=requests.post(url,data=payload)

import json
r = requests.post(url, data = json.dumps(payload))
payload = {'ID':'2','Value':'2 value'}
r = requests.post(url, data = json.dumps(payload))
payload = {'ID':'3','Value':'3 value'}
r = requests.post(url, data = json.dumps(payload))
payload = {'ID':'4','Value':'4 value'}
r = requests.post(url, data = json.dumps(payload))
payload = {'ID':'5','Value':'5 value'}
r = requests.post(url, data = json.dumps(payload))
payload = {'ID':'6','Value':'6 value'}
r = requests.post(url, data = json.dumps(payload))

for i in range (100):
    payload = {'ID':str(i),'Value': str(i) +'value'}
    r = requests.post(url, data = json.dumps(payload))

r.text
r.status_code
