package invasion

import (
	"AlienInvasion/aliens"
	"AlienInvasion/world"
	"fmt"
)

type Invasion struct {
	aliensInvading *aliens.Aliens
	worldAttacked *world.World
}

func New(worldFilename string, alienPopulation int) Invasion {
	worldAttacked := world.LoadWorld(worldFilename)
	aliensInvading := aliens.New(alienPopulation, worldAttacked.NumberOfCities())
	return Invasion{&aliensInvading, &worldAttacked}
}

func (inv Invasion) GetWorld() *world.World {
	return inv.worldAttacked
}

func (inv Invasion) GetAliens() *aliens.Aliens {
	return inv.aliensInvading
}


// Each alien can move only to one neighboring city per day.
func (anInv *Invasion) RunInvasionSync(days int) {

	fmt.Println( anInv.worldAttacked)
	fmt.Println( anInv.aliensInvading)

	fmt.Println("Initial Number of Cities = ",anInv.worldAttacked.NumberOfCities())


	for i := 0; i < days; i++ {
		fmt.Println("Day = ", i)
		//fmt.Println( anInv.worldAttacked)
		//fmt.Println( anInv.aliensInvading)
		// Sync day, move all aliens one city and then do fighting.

		// Move all aliens, one city step.
		anyMovement := false
		for a := 0; a < anInv.aliensInvading.NumberOfAliens(); a++ {
			//fmt.Println( "    Alien = ",a )
			if !anInv.aliensInvading.IsDead(a) {
				fmt.Println( "    Alien = ",a )
				aLoc := anInv.aliensInvading.Location(a)
				newCity := anInv.worldAttacked.RandomNeighboringCity(aLoc)
				fmt.Printf( "    Moving alien %d from location %d to location %d\n",a,aLoc,newCity)
				if a == 2 && aLoc == 5 && newCity ==0 {
					fmt.Println("DEBUG")
				}

				if aLoc != newCity {
					anyMovement = true
				}
				anInv.aliensInvading.MoveAlienSync(a,newCity)
			} else {
				fmt.Println( "    Alien is Dead = ",a )
			}
		}
		// Do sync fighting.
		//destroyedCities :=
		destroyedCities := anInv.aliensInvading.FightingSync()
		// Iterate destroyed cities, erase cities from graph, and mark killed aliens as dead.
		for loc, aliensDead := range destroyedCities {
			fmt.Printf("location[%d] aliensDead[%d]\n", loc, aliensDead)
			fmt.Println("Destroying city =",anInv.worldAttacked.CityName(loc))
			anInv.worldAttacked.DestroyCity(loc)
			for a := range aliensDead {
				anInv.aliensInvading.SetDead(a)
			}
		}
		fmt.Println("Number of Cities = ",anInv.worldAttacked.NumberOfCities())
		if anInv.worldAttacked.NumberOfCities() == 0 {
			fmt.Println("WARNING: All cities were destroyed!!! Stopping simulation...")
			break
		}
		if anInv.worldAttacked.NumberOfCities() == 1 {
			fmt.Println("WARNING: Only one city remaining! Stopping simulation...")
			break
		}
		if !anyMovement {
			fmt.Println("WARNING: No movements detected, aliens are trapped or all dead! Stopping simulation...")
			break
		}

	}
}



