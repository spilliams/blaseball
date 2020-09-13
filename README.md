being a small? project? about blaseball?

This repository holds code that does stuff. I'm not really sure if you can use
`go get` to fetch the binaries, or if you have to build them yourself.

Building `cmd/server` will build you a web server that operates on port 8080
(maybe someday that'll be configurable). This server will (someday) mirror the
[blaseball database API](https://github.com/Society-for-Internet-Blaseball-Research/blaseball-api-spec). It will also
provide a few key extra features, discussed below. Right now this server stores
its data in memory. If and when I deploy this server somewhere for more general
use, I should (will) set up a more reliable data storage layer.

Building `cmd/cli` will build you a command-line tool that will query against
a blaseball API. I say "a" blaseball API because right now it will build
against the local server (it assumes you're running one). Someday this will be
configurable, so you can tell it to hit your local server, or the official API,
or maybe a hosted third-party API. The future is so uncertain!

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
