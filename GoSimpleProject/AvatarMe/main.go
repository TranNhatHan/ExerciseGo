package main

import (
	"AvatarMe/identicon"
)

func main() {
	ava := identicon.Generate("han2k3x2@gmail.com")
	ava.WriteImage("avatar")
}
