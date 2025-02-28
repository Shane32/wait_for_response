# Wait For Response GitHub Action

**Wait For Response** is a composite GitHub Action designed to poll a specified URL until a desired HTTP response code is received (optionally verifying that a specific string appears in the response body) or a timeout is reached. It supports both Windows (PowerShell) and Unix (Bash) runners.

## Features

- **Polling Mechanism:** Continuously polls the given URL until a successful response is returned or the timeout is exceeded.
- **Configurable Options:** Easily customize the URL, expected response code, timeout, polling interval, initial sleep delay, and an optional string to look for in the response body.
- **Cross-Platform:** Compatible with both Windows and Unix runners.

## Inputs

| Input         | Description                                        | Required | Default            |
|---------------|----------------------------------------------------|----------|--------------------|
| `url`         | URL to poll                                        | No       | `http://localhost/`|
| `responseCode`| HTTP response code to wait for                     | No       | `200`              |
| `timeout`     | Timeout before giving up (in milliseconds)         | No       | `30000`            |
| `interval`    | Interval between polling attempts (in milliseconds)| No       | `200`              |
| `sleep`       | Delay before the initial request (in milliseconds) | No       | `0`                |
| `lookfor`     | String to look for in the response body            | No       | (empty string)     |

## Outputs

| Output  | Description                                                                |
|---------|----------------------------------------------------------------------------|
| `result`| Returns `0` if the expected response is received within the timeout, or `1` otherwise. |

## How It Works

1. **Initial Delay:** If specified, the action waits for the provided `sleep` duration before starting.
2. **Polling:** The action continuously polls the URL at intervals defined by `interval` until:
   - The HTTP response code matches the provided `responseCode`.
   - If a `lookfor` string is provided, the response body must contain the specified string.
3. **Timeout:** If the condition is not met within the specified `timeout`, the action times out and returns a failure (`result` = `1`).

## Example Usage

Below is an example of how to use the **Wait For Response** action in your GitHub workflow:

```yaml
name: Wait for Service to be Ready

on: [push]

jobs:
  wait-for-service:
    runs-on: ubuntu-latest
    steps:
      - name: Wait for Response from Service
        uses: Shane32/wait_for_response@v2
        with:
          url: 'https://example.com/health'
          responseCode: 200
          timeout: 60000
          interval: 500
          sleep: 1000
          lookfor: 'Service is up'
```

In this example, the workflow will:

- Poll `https://example.com/health` after an initial 1-second delay.
- Wait for a response with a status code of `200`.
- Check if the response body contains the text "Service is up".
- Poll every 500 milliseconds, and give up after 60 seconds if the conditions are not met.

## Author

Shane Krueger

## License

This project is licensed under the MIT License. 

---

Feel free to contribute or open issues if you have any suggestions or encounter any problems with the action. Enjoy using **Wait For Response** in your CI/CD workflows!
