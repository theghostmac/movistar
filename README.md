# Movistar
Movistar is a movie rating app using micro service architecture. It rates a movie, and provides the movie's metadata.
Metadata available:
- ID
- title
- year
- description
- director
- cast

---
## Working principle
- Store movie rating on a db. Append and delete functionality.
- Return aggregated movie ratings and store aggregations.

Movie metadata are stored separately from the rating data.

### Rating API
- Store the rating record with: ID of user giving the rating, type of record, ID of record, rating value
- Get the aggregated rating for a record by its ID and type.

### Service Implementation
* [ ] Movie metadata service (can evolve) -> metadata
* [ ] Rating service -> rating
* [ ] Movie service (can evolve) -> movie

## Add to changelog.md
- support for rating deletion

---
## Installation
Run the command:
```shell
go build github.com/theghostmac/movistar@latest
```

Or download from the Releases section of this repostory.