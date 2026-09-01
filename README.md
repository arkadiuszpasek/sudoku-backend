# sudoku-backend

## Prerequisites

- Docker (SAM will require it, no need for any images)

## Testing and running

- Testing should be done, run with `go test .`
- Running locally done with docker and `sam local start-api`

```
sam local start-api  --warm-containers EAGER --host 0.0.0.0 --port 4001
```

## Deploying

- Make you have aws user configured, `aws sts get-caller-identity`
- Make sure you have `go` and `sam` installed
- Sam config is commited to the repo, the deployment can be done via `sam build` and `sam deploy`

## Manual testing

- Feel free to extend console/ scripts for manual testing locally
  - run with `go run -C console .` from root
  - run with `go run .` from `console/`
