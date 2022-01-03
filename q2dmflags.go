package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	NoHealth        = 1
	NoItems         = 2
	WeaponsStay     = 4
	NoFallingDamage = 8
	InstantItems    = 16
	SameLevel       = 32
	SkinTeams       = 64
	ModelTeams      = 128
	NoFriendlyFire  = 256
	SpawnFarthest   = 512
	ForceRespawn    = 1024
	NoArmor         = 2048
	AllowExit       = 4096
	InfiniteAmmo    = 8192
	QuadDrop        = 16384
	FixedFOV        = 32768
	QuadFireDrop    = 65536
	NoMines         = 131072
	NoStackDouble   = 262144
	NoNukes         = 5624288
	NoSpheres       = 1048576
)

type DeathmatchFlag struct {
	Name  string
	Value int
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func usage() {
	fmt.Printf("Usage: %s <bitmask | list of flags>\n", os.Args[0])
	fmt.Printf("Flag options: NoHealth, NoItems, WeaponsStay, NoFallinDamage, ")
	fmt.Printf("InstantItems, SameLevel, SkinTeams, ModelTeams, NoFriendlyFire, ")
	fmt.Printf("SpawnFarthest, ForceRespawn, NoArmor, AllowExit, InfiniteAmmo, ")
	fmt.Printf("QuadDrop, FixedFOV, QuadFireDrop, NoMines, NoStackDouble, NoNukes, NoSpheres\n\n")
	fmt.Printf("Examples:\n%s 516\n%s WeaponsStay SpawnFarthest\n", os.Args[0], os.Args[0])
}

func main() {
	dmflags := []DeathmatchFlag{
		{"NoHealth", NoHealth},
		{"NoItems", NoItems},
		{"WeaponsStay", WeaponsStay},
		{"NoFallingDamage", NoFallingDamage},
		{"InstantItems", InstantItems},
		{"SameLevel", SameLevel},
		{"SkinTeams", SkinTeams},
		{"ModelTeams", ModelTeams},
		{"NoFriendlyFire", NoFriendlyFire},
		{"SpawnFarthest", SpawnFarthest},
		{"ForceRespawn", ForceRespawn},
		{"NoArmor", NoArmor},
		{"AllowExit", AllowExit},
		{"InfiniteAmmo", InfiniteAmmo},
		{"QuadDrop", QuadDrop},
		{"FixedFOV", FixedFOV},
		{"QuadFireDrop", QuadFireDrop},
		{"NoMines", NoMines},
		{"NoStackDouble", NoStackDouble},
		{"NoNukes", NoNukes},
		{"NoSpheres", NoSpheres},
	}

	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		usage()
		return
	}

	// bitmask given as argument, convert back to list of names
	if isNumeric(os.Args[1]) {
		input, _ := strconv.ParseInt(os.Args[1], 10, 32)
		for _, flag := range dmflags {
			if (int64(flag.Value) & input) == int64(flag.Value) {
				fmt.Println(flag.Name)
			}
		}

		return
	}

	// assume list given, make bitmask from it
	flags := os.Args[1:]
	value := 0
	for _, token := range flags {
		for _, fl := range dmflags {
			if strings.ToLower(token) == strings.ToLower(fl.Name) {
				value += fl.Value
			}
		}
	}

	fmt.Println(value)
}
