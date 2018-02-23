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

## Environment Variables

### `SUCCESS_RATE`

This environment variable must be a float in the range `[0.0, 1.0)`. If not set, or invalid,
defaults to `0.9`. The app generates a random number. When that number is in the range
`[0.0, SUCCESS_RATE]`, the app returns an HTTP 200 response and message. When in
the range `(SUCCESS_RATE, 1.0)`, an HTTP 500 response and message are returned.

### `MIN_DELAY_MS` and `MAX_DELAY_MS`

These environment variables must be an unsigned 64 bit integer. If not set, or invalid,
`MIN_DELAY_MS` defaults to `0` and `MAX_DELAY_MS` defaults to `MIN_DELAY_MS`.
Before responding, the app will wait for a random period of milliseconds in the range
`[MIN_DELAY_MS, MAX_DELAY_MS)`.
