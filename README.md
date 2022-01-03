# Q2DMflags
A command-line utility to generate an integer bitmask for Quake 2 deathmatch flags or convert a bitmask back to a list of flag names.

## Possible `dmflags` and Values
```
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
```

## Compiling
To build for your current system:
`# go build q2dmflags.go` 

To build for another OS or ARCH (ex: 32bit windows):
`# GOOS=windows GOARCH=386 CGO_ENABLED=0 go build q2dmflags.go`

## Usage
`# q2dmflags <integer bitmask>`
`# q2dmflags <space separated list of flag names>`

