package main

import "github.com/stellayazilim/neptune_cms/internal/rest"

func main() {
	r := rest.Rest()

	r.Run(":8082")
}
