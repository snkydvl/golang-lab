module example.com/mainapp

go 1.21

require (
	example.com/car v0.0.0
	example.com/sportscar v0.0.0
	example.com/transport v0.0.0
)

replace example.com/transport => ../transport

replace example.com/car => ../car

replace example.com/sportscar => ../sportscar
