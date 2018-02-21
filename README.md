# Sample Error App

This app will throw an error randomly.

## Build

```
git clone https://github.com/jghiloni/sample-error-app.git
cd sample-error-app
dep ensure
go build
```

## Run

### Local
```
PORT=8000 ./sample-error-app
```

### Cloud Foundry
```
dep ensure
cf push
```

## Options

### `ERROR_THRESHOLD` environment variable

This environment variable must be a float in the range `[0.0, 1.0)`. If not set, or invalid,
defaults to `0.9`. The app generates a random number. When that number is in the range
`[0.0, ERROR_THRESHOLD]`, the app returns an HTTP 200 response and message. When in
the range `(ERROR_THRESHOLD, 1.0)`, an HTTP 500 response and message are returned.
