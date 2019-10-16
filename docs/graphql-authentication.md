# GraphQL Authentication

The GraphQL api uses a basic authentication scheme for the current implementation.
All api queries must pass an "Authentication" HTTP header in the request.
The value of the header takes the form `Server ###` where `###` is the server
code.

For example, the Authentication header for the sample manager server would
have the value `Server 125`.

## Adding the header

In the GraphQL playground, click "HTTP HEADERS" in the bottom left corner.
The value is the HTTP headers in JSON format. In this case:

```json
{
  "Authorization": "Server 478"
}
```

Again, replace the `478` with the server code you want to use.

## Javascript Fetch API

Here's an example of adding the Authentication header using Javascript's
fetch API; this example is in the form of an async function, the main idea
is passing in the options object with a `headers` key:

```javascript
async function postData(url = '/graphql', data = {}) {
  const response = await fetch(url, {
    method: 'POST', // GraphQL only response to POST
    headers: {
      'Content-Type': 'application/json', // Content type must be json
      'Authentication': 'Server 478', // Again, replace 478 with the code you want to use
    },
    body: JSON.stringify(data) // body data type must match "Content-Type" header
  });
  return await response.json(); // parses JSON response into native JavaScript objects
}
```
