package fisiks

import "math"

const (
	SpeedOfLight = 299_792_458

	// Normal gravity value, in m.s⁻2
	Gravity = 9.806_65
	// Gravitational constant
	GravitationalConstant = 6.674_30e-11

	// Elementary charge, in Coulomb (C)
	ElementaryCharge = 1.602_176_620_8e-19
	// Void permeability
	Mu0 = 4e-7 * math.Pi
	// Void permittivity
	Epsilon0 = 1 / (Mu0 * SpeedOfLight * SpeedOfLight)
	// Coulomb constant
	K = 1 / (4 * math.Pi * Epsilon0)

	// time interval
	dt = 0.027

	// Shock absorption coefficient
	bounce = -0.5
)
