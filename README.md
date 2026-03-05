# erepgo

Go client for the [eRepublik API](https://api.erepublik.com/doc/) (eAPI). Authenticated REST client for citizens, countries, regions, battles, wars, laws, and more.

## Installation

```bash
go get github.com/darkmantle/erepgo
```

## Configuration

You need API keys from eRepublik (public key and secret key). Keep the secret key private and never send it over the network; the client uses it only to sign requests.

## Usage

### Create a client

```go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/darkmantle/erepgo"
)

func main() {
    client := erepgo.NewClient(
        os.Getenv("EREPUBLIC_PUBLIC_KEY"),
        os.Getenv("EREPUBLIC_SECRET_KEY"),
    )
}
```

### Generic calls

Call any resource/action with optional query parameters:

```go
// Raw response body
body, err := client.Call("citizen", "profile", map[string]string{"citizenId": "123"})
if err != nil {
    log.Fatal(err)
}

// Decode JSON into a struct
var profile map[string]interface{}
err = client.CallJSON("citizen", "profile", map[string]string{"citizenId": "123"}, &profile)
if err != nil {
    log.Fatal(err)
}
fmt.Println(profile)
```

### Typed resource methods

Use the provided helpers for each API resource:

```go
// Citizen
body, err := client.CitizenProfile(123)

// Country & countries
body, err := client.CountryRegions(42)
body, err := client.CountriesIndex()

// Region (page 0 for first page)
body, err := client.RegionCitizens(regionID, 1)

// Map & industries
body, err := client.MapData()
body, err := client.IndustriesIndex()

// Battle & war
body, err := client.BattleIndex(battleID)
body, err := client.WarBattles(warID)

// Citizens & laws
body, err := client.CitizensRegistered()
body, err := client.LawsActive()
body, err := client.LawsRecent()
```

### Decoding responses

Responses are `[]byte`; decode with standard `encoding/json` or the helper:

```go
body, err := client.CitizenProfile(123)
if err != nil {
    log.Fatal(err)
}

var data map[string]interface{}
if err := erepgo.DecodeJSON(body, &data); err != nil {
    log.Fatal(err)
}
```

### XML responses

Request XML instead of JSON (default):

```go
client.SetFormat("xml")
body, err := client.CitizenProfile(123)
// body is XML
```

## API resources

| Resource   | Actions        | Methods / Call example |
| ---------- | -------------- | ----------------------- |
| citizen    | profile        | `CitizenProfile(citizenID)` |
| country    | regions        | `CountryRegions(countryID)` |
| countries  | index          | `CountriesIndex()` |
| region     | citizens       | `RegionCitizens(regionID, page)` |
| map        | data           | `MapData()` |
| industries | index          | `IndustriesIndex()` |
| battle     | index          | `BattleIndex(battleID)` |
| war        | battles        | `WarBattles(warID)` |
| citizens   | registered     | `CitizensRegistered()` |
| laws       | active, recent | `LawsActive()`, `LawsRecent()` |

For custom resources or parameters, use `client.Call(resource, action, params)`.

## License

See repository license.
