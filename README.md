# Wait For Response GitHub Action

**Wait For Response** is a composite GitHub Action designed to poll a specified URL until a desired HTTP response is received or a timeout is reached. In addition to matching an exact response code, this supports wildcard patterns like `"2xx"`, `"4xx"`, and `"5xx"`, which allow you to check for a class of responses rather than a specific code. This action works on both Windows and Unix runners.

## Features

- **Polling Mechanism:** Continuously polls the given URL until a matching HTTP response is returned or the timeout is exceeded.
- **Wildcard Support:** Specify response code patterns (e.g., `"2xx"`) to match any response within that class.
- **Configurable Options:** Customize the URL, expected response code (exact or wildcard), timeout, polling interval, initial sleep delay, and an optional string to look for in the response body.
- **Cross-Platform Compatibility:** Works seamlessly on both Windows (using PowerShell) and Unix (using Bash) runners.

## Inputs

| Input         | Description                                                                 | Required | Default            |
|---------------|-----------------------------------------------------------------------------|----------|--------------------|
| `url`         | URL to poll                                                                 | No       | `http://localhost/`|
| `responseCode`| HTTP response code to wait for, or a wildcard (e.g., `"200"` or `"2xx"`)    | No       | `200`              |
| `timeout`     | Timeout before giving up (in milliseconds)                                  | No       | `30000`            |
| `interval`    | Interval between polling attempts (in milliseconds)                         | No       | `200`              |
| `sleep`       | Delay before the initial request (in milliseconds)                          | No       | `0`                |
| `lookfor`     | String to look for in the response body (optional)                          | No       | (empty string)     |

## Outputs

| Output  | Description                                                                                           |
|---------|-------------------------------------------------------------------------------------------------------|
| `result`| Returns `0` if a matching response (exact or wildcard) is received within the timeout, otherwise `1`.  |

## How It Works

1. **Initial Delay:**  
   If a `sleep` value is provided, the action waits that many milliseconds before starting.

2. **Polling:**  
   The action repeatedly polls the specified URL at intervals defined by the `interval` input until:
   - The HTTP response code either exactly matches the provided `responseCode` **or** matches the wildcard pattern (for example, if `responseCode` is `"2xx"`, any response code starting with `2` is accepted).
   - If a `lookfor` string is provided, the action also checks whether the response body contains the specified text.

3. **Timeout:**  
   If the conditions are not met within the specified `timeout` period, the action terminates with a failure (`result` = `1`).

## Example Usage

Below is an example workflow using the **Wait For Response** action:

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
          responseCode: 2xx        # Accept any 2xx HTTP status code
          timeout: 60000           # Timeout after 60 seconds
          interval: 500            # Retry every 500ms
          sleep: 1000              # Wait 1 second before first attempt
          lookfor: 'Service is up'
```

In this example, the workflow will:

- Poll `https://example.com/health` after an initial 1-second delay.
- Accept any HTTP response code in the `2xx` range.
- Verify that the response body contains the text "Service is up."
- Poll every 500 milliseconds, and if the expected conditions are not met within 60 seconds, the action will time out.

## Author

Shane Krueger

## License

This project is licensed under the MIT License.

---

Feel free to contribute or open issues if you have any suggestions or encounter any problems with the action. Enjoy using **Wait For Response** in your CI/CD workflows!
