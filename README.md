# http-loupe

A command line http inspection server. `http-loupe` allows you to inspect, save or debug http requests you are sending.
You can now  view the headers, encoding and body of your requests as they would be received at the expected endpoint.

Requests are currently only persistant within sessions. You will therefore lose all your requests once the server is killed
or you close the program. This is because `http-loupe` currently stores requests in memory.

## Ideas
- Customize the response for specific or all paths (need to workout how)
- save requests and do the same requests on a different endpoint? (use http-client to send request. `http-loupe` will act as a proxy)
- persist across sessions
- Add https://asciinema.org/ demo

## Dependencies
- gopkg.in/readline.v1