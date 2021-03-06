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
the things that the aforementioned server could tell you, but it does so in
your terminal instead of in a browser. Honestly it's only here as a test bed
before I build an honest front-end for the server.

# Contributions

At this time I am not accepting contributions. I don't even know how blaseball
works yet, and this is just a fun pet project to distract myself from systemic
racism, global pandemics, and devastating wildfires.

If you feel very strongly about it, you could start a new issue, or set up a
fork and Pull Request, but honestly I don't know how much time I can devote to
open-source style maintenance on top of just-farting-around style maintenance.

# Server features

Running the server gives you several extra features on top of the official
blaseball API:

- not yet implemented

In the future, I plan on adding more several extra features, such as:

- forbidden knowledge will be hidden unless a query parameter is passed in
- time-series and historical data on players

# Development todos:

- the data store is a sore spot. before hosting this I really want to get it in
a database.
- when do I release a version of this thing? when I feel comfortable publishing
it? how many of the following do I need:
    - [ ] database data store
    - [x] server is smart about freshness-checking
    - [x] server is smart about Forbidden Knowledge
    - [ ] server is API-complete, Objects edition
        - [ ] leagues
        - [ ] subleagues
        - [x] divisions
        - [x] teams
        - [x] players
        - [ ] seasons
        - [ ] games
        - [ ] playoffs
    - [ ] API-completeness, Summaries edition
        - [ ] events
        - [ ] standings
        - [ ] tiebreakers
    - [ ] API-completeness, Elections edition
        - [ ] blessings
        - [ ] decrees
        - [ ] election details
        - [ ] election results
    - [x] API-completeness, Statsheets edition
        - [x] season
        - [x] game
        - [x] team
        - [x] player
    - [ ] CLI is API-complete, for some subset of the server API
    - [ ] server is configured and running on a host somewhere
    - [ ] tasks to fetch current season, game
    - [ ] refactor server endpoints to be more restful
    - [ ] publish openAPI spec (for mine or both?)
    - [ ] health metrics for my services
    - [ ] canary for when the official api changes
