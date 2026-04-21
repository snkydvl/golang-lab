module github.com/snkydvl/vehicle-lab/mainapp

go 1.21

require (
    github.com/snkydvl/vehicle-lab/transport v0.0.0
    github.com/snkydvl/vehicle-lab/car v0.0.0
    github.com/snkydvl/vehicle-lab/sportscar v0.0.0
)

replace github.com/snkydvl/vehicle-lab/transport => ../transport
replace github.com/snkydvl/vehicle-lab/car => ../car
replace github.com/snkydvl/vehicle-lab/sportscar => ../sportscar
