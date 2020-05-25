# caching_api
- A fullstack Web-App implementing caching Routine
- Implemented LRU (Least Recently Used) algorithm to handle the object
replacement from the cache.
- Developed HTTP based server implementation of connection protocol to
enable the cache routine to listen to a particular port to enable client to interact
with it for accessing the objects directly from the cache.
- It utilizes RESTful API protocols
- Developed frontend using React to access backend services and display live
state of the cache .
- Used Jmeter for load testing

## Live Version -> [here](http://cache-routine.herokuapp.com)

# Usage

First start the server

```
go run main.go
``` 
Use python script to POST ID-value pairs 
```
python postData.py
```

## cache API
```
GET			/cache/get/{id}
PUT 		/cache/set
```

## state API
```
state 		/cache/state
```

# To-Do
- [x] Add API support for current state.
- [x] Implement Front-end using react.
- [x] Deploy
- [x] User input size of cache
- [ ] Add support for more caching algorithms.
