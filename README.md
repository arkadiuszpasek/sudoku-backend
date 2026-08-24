# sudoku-backend

## Testing and running

- Feel free to extend console/ scripts for manual testing locally, run with `go run .`
- Testing should be done, run with `go test .`

## Deploying

- Make you have aws user configured, `aws sts get-caller-identity`
- Make sure you have `go` and `sam` installed
- Sam config is commited to the repo, the deployment can be done via `sam build` and `sam deploy`
