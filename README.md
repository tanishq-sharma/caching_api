# caching_api
This is a  HTTP server implementation of various caching algorithms. It utilizes RESTful API protocols. ID and value pairs can be of any datatype .

# Usage

First start the server

```
go run start.go
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
1.Add API support for current state.
2.Implement GUI.
3.Add support for more caching algorithms.