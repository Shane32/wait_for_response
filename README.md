# Wait For Response docker action

This action makes a GET requests to a given URL until the required response code is retrieved or the timeout is met.

## Inputs

### `url`

The URL to poll. Default `"http://localhost/"`

### `responseCode`

Response code to wait for. Default `"200"`

### `timeout`

Timeout before giving up in milliseconds. Default `"30000"`

### `interval`

Interval between polling in ms. Default `"200"`

### `sleep`

Amount to sleep in ms before initial request. Default `"0"`

### `lookfor`

String to look for in the response. Default `""`

## Example usage

```yml
uses: Shane32/wait_for_response@v2
with:
  url: 'http://localhost:8081/'
  responseCode: 200
  timeout: 2000
  interval: 500
  sleep: 2000
  lookfor: 'OK'
```
