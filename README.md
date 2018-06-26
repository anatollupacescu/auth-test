# auth-test

to run just execute 

```bash
ENDPOINT_LIST=http://golang.org,http://google.com,http://unu go run cmd/main/main.go
```

sample output
```bash
2018/06/26 16:35:00 request to url 'http://golang.org' returned code: 200
2018/06/26 16:35:00 request to url 'http://google.com' returned code: 200
2018/06/26 16:35:00 Get http://unu: dial tcp: lookup unu: no such host
```
