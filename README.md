# goad
A load-testing tool in Go implementing some of a similar CLI to `curl`.

## Implemented curl flags
* `--request/-X` for setting an HTTP method
* `--data/-d` for setting a request body
* `--header/-H` for settings headers (repeatable)

## Unique flags
* `--url` for specifying the URL to hit
* `--steps` a repeatable flag for setting the stepping for tests
* `--cooldown` to set the cooldown period between test runs/steps