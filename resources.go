package erepgo

import (
	"encoding/json"
	"fmt"
)

// CitizenProfile returns the profile for the given citizen ID.
// Resource: citizen, Action: profile.
func (c *Client) CitizenProfile(citizenID int) ([]byte, error) {
	return c.Call("citizen", "profile", map[string]string{"citizenId": fmt.Sprint(citizenID)})
}

// CountryRegions returns regions for the given country ID.
// Resource: country, Action: regions.
func (c *Client) CountryRegions(countryID int) ([]byte, error) {
	return c.Call("country", "regions", map[string]string{"countryId": fmt.Sprint(countryID)})
}

// CountriesIndex returns the countries index.
// Resource: countries, Action: index.
func (c *Client) CountriesIndex() ([]byte, error) {
	return c.Call("countries", "index", nil)
}

// RegionCitizens returns citizens in a region with optional pagination.
// Resource: region, Action: citizens.
func (c *Client) RegionCitizens(regionID int, page int) ([]byte, error) {
	params := map[string]string{"regionId": fmt.Sprint(regionID)}
	if page > 0 {
		params["page"] = fmt.Sprint(page)
	}
	return c.Call("region", "citizens", params)
}

// MapData returns map data.
// Resource: map, Action: data.
func (c *Client) MapData() ([]byte, error) {
	return c.Call("map", "data", nil)
}

// IndustriesIndex returns the industries index.
// Resource: industries, Action: index.
func (c *Client) IndustriesIndex() ([]byte, error) {
	return c.Call("industries", "index", nil)
}

// BattleIndex returns battle data for the given battle ID.
// Resource: battle, Action: index.
func (c *Client) BattleIndex(battleID int) ([]byte, error) {
	return c.Call("battle", "index", map[string]string{"battleId": fmt.Sprint(battleID)})
}

// WarBattles returns battles for the given war ID.
// Resource: war, Action: battles.
func (c *Client) WarBattles(warID int) ([]byte, error) {
	return c.Call("war", "battles", map[string]string{"warId": fmt.Sprint(warID)})
}

// CitizensRegistered returns registered citizens data.
// Resource: citizens, Action: registered.
func (c *Client) CitizensRegistered() ([]byte, error) {
	return c.Call("citizens", "registered", nil)
}

// LawsActive returns active laws.
// Resource: laws, Action: active.
func (c *Client) LawsActive() ([]byte, error) {
	return c.Call("laws", "active", nil)
}

// LawsRecent returns recent laws.
// Resource: laws, Action: recent.
func (c *Client) LawsRecent() ([]byte, error) {
	return c.Call("laws", "recent", nil)
}

// DecodeJSON is a helper to decode API response bytes into v.
func DecodeJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
