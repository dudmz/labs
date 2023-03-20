# k6-test
This lab is about testing the k6 component to do load-testing in components that
implements REST API's.

## How to run it
1. Copy the `api/` directory to your Go project's home directory (usually is `/home/your-username/go/src/github.com/username).
2. Change the go.mod declaration to point it to your username in GitHub.
3. In `script.js`, replace the URL with your local instance's. If you're using `docker`
to run k6, you need to specify the hostname of your local instance, so you either need
to issue the `hostname` command in whatever CLI that's provided in your OS.
4. Run the back-end with `go build && ./k6-test` in the project's root directory.
5. In another shell, you're going to run the k6. To run it, if using `docker`, issue the following command:

```bash
$ docker run --rm -i grafana/k6 run - <script.js
```

- This command pulls a k6 image (if not present), and executes a container with the said
image.
- `-i` indicates that the container execution is interactive, so you're thrown
into the container's shell.
- `--rm` flag indicates that the container is removed after finishing the command's
execution.

For some reason, it fails to connect the API 3 or 4 times before establishing connectivity,
I suspect that the docker's container lags in resolving the hostname.
