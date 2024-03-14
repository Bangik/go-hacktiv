### How to run with docker

```
  docker run -it --rm --name hacktiv-ass-3 -p 8080:8080 -v "$(pwd)":/usr/src/myapp -w /usr/src/myapp golang:alpine go run .
```

Open in http://localhost:8080