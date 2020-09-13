being a small? project? about blaseball?

This repository holds code that does stuff. I'm not really sure if you can use
`go get` to fetch the binaries, or if you have to build them yourself.

Building `cmd/server` will build you a web server that operates on port 8080
(or whatever `PORT` says). This server will (someday) mirror the
[blaseball database API](https://github.com/Society-for-Internet-Blaseball-Research/blaseball-api-spec). It will also
provide a few key extra features, discussed below. Right now this server stores
its data in memory. If and when I deploy this server somewhere for more general
use, I should (will) set up a more reliable data storage layer.

Building `cmd/cli` will build you a command-line tool that will query against
a blaseball API. It defaults to the official one, but you can use flags to set
a custom url or pick a hardcoded local url. This tool will tell you many of
the things that the server could tell you, but in your terminal instead of in a
browser. Honestly it's only here as a test bed before I build an honest
front-end for the server.

# Contributions

At this time I am not accepting contributions. I don't even know how blaseball works yet, and this is just a fun pet project to distract myself from systemic
racism, global pandemics, and devastating wildfires.

# Server features

Running the server gives you several extra features on top of the official
blaseball API:

- not yet implemented

In the future, I plan on adding more several extra features, such as:

- forbidden knowledge will be hidden unless a query parameter is passed in
- time-series data such as a list of teams and roles a player has played for over time
