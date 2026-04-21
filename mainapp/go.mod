module github.com/snkydvl/golang-lab/mainapp

go 1.21

require (
    github.com/snkydvl/golang-lab/transport v0.0.0
    github.com/snkydvl/golang-lab/car v0.0.0
    github.com/snkydvl/golang-lab/sportscar v0.0.0
)

replace github.com/snkydvl/golang-lab/transport => ../transport
replace github.com/snkydvl/golang-lab/car => ../car
replace github.com/snkydvl/golang-lab/sportscar => ../sportscar
