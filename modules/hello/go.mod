module vitozap/hello

go 1.26.2

replace vitozap/greetings => ../greetings

require (
	vitozap/greetings v0.0.0-00010101000000-000000000000
	vitozap/slices v0.0.0-00010101000000-000000000000
	vitozap/structures v0.0.0-00010101000000-000000000000
)

replace vitozap/slices => ../slices

replace vitozap/structures => ../structures
