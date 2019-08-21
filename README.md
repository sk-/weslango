# weslango
GRPC server to detect the language of a text. This is an evolution of [weslang](https://github.com/deezer/weslang) which uses now Google's
[CLD3](https://github.com/google/cld3), Grpc and Go.

## Test
```
bazel test //...
```

## Build
```
bazel build //weslango -c opt
```

## Local Testing
For local testing you can start the server with:

```
bazel run //weslango
```

and then use [evans](https://github.com/ktr0731/evans) to make requests to the
server.

## Contribute
Feel free to send Pull Requests or create Issues. Even Issues asking for better
coding styles and practices are welcome, as it has been a while since I last
coded in Go.
