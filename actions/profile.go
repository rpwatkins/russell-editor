package actions

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

// Profile data
type Profile struct {
	Author      string
	Profession  string
	Institution string
	Email       string
	SAOffice    string
	SAHours     string
	VOffice     string
	VHours      string
	Phone       string
}

// LoadProfile loads the toml data into a Profile struct
func LoadProfile(dataLocation string) (Profile, error) {
	// path
	path := filepath.Join(dataLocation, "data", "profile.toml")

	var profile Profile

	tomlBytes, err := os.ReadFile(path)
	if err != nil {
		return profile, err
	}

	// unmarshall toml data to course struct
	if err := toml.Unmarshal([]byte(tomlBytes), &profile); err != nil {
		return profile, err
	}

	// return profile toml
	return profile, nil

}
