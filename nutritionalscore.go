package main

type ScoreType int

const (
	// Food is the score type for food
	Food ScoreType = iota
	// Beverage is the score type for beverage
	Beverage
	Water
	Cheese
)

type NutritionalScore struct {
	value     int
	Positive  int
	Negative  int
	ScoreType ScoreType
}

var scoreToLetter = []string{"A", "B", "C", "D", "E"}

type EnergyKcal float64
type SaturatedFattyAcid float64
type SodiumMg float64
type SugarGram float64
type FruitPercent float64
type FiberGram float64
type ProteinGram float64

type NutritionalData struct {
	Energy             EnergyKcal
	Sugars             SugarGram
	SaturatedFattyAcid SaturatedFattyAcid
	Sodium             SodiumMg
	Fruits             FruitPercent
	Fiber              FiberGram
	Proteins           ProteinGram
	isWater            bool
}

var energyLevel = []float64{335, 670, 1005, 1340, 1675, 2010, 2345, 2680, 3015, 3350}
var sugarLevel = []float64{4.5, 9, 13.5, 18, 22.5, 27, 31, 36, 40, 45}
var saturatedFattyAcidLevel = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var sodiumLevel = []float64{90, 180, 270, 360, 450, 540, 630, 720, 810, 900}
var fruitLevel = []float64{40, 36, 33, 30, 27, 24, 20, 16, 12, 8}
var fiberLevel = []float64{3.5, 3, 2.5, 2, 1.5, 1, 0.5, 0, -0.5, -1}
var proteinLevel = []float64{8, 7.5, 7, 6.5, 6, 5.5, 5, 4.5, 4, 3.5}

var energyLevelBeverage = []float64{0, 30, 60, 90, 120, 150, 180, 210, 240, 270}
var sugarLevelBeverage = []float64{0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, 4.5, 5}

func getPointsFromRange(value float64, level []float64) int {
	for i, v := range level {
		if value <= v {
			return i
		}
	}
	return 10
}

func (e EnergyKcal) GetPoints(st ScoreType) int {

	if st == Food {
		return getPointsFromRange(float64(e), energyLevel)
	} else if st == Beverage {
		return getPointsFromRange(float64(e), energyLevelBeverage)
	} else if st == Water {
		return 0
	} else if st == Cheese {
		return getPointsFromRange(float64(e), energyLevel)
	}
	return 0
}

func (s SaturatedFattyAcid) GetPoints(st ScoreType) int {
	if st == Beverage {
		return getPointsFromRange(float64(s), saturatedFattyAcidLevel)
	}
	return getPointsFromRange(float64(s), saturatedFattyAcidLevel)

}

func (s SodiumMg) GetPoints(st ScoreType) int {
	if st == Beverage {
		return getPointsFromRange(float64(s), sodiumLevel)
	}
	return getPointsFromRange(float64(s), sodiumLevel)
}

func (s SugarGram) GetPoints(st ScoreType) int {

	if st == Food {
		return getPointsFromRange(float64(s), sugarLevel)
	} else if st == Beverage {
		return getPointsFromRange(float64(s), sugarLevelBeverage)
	} else if st == Water {
		return 0
	} else if st == Cheese {
		return getPointsFromRange(float64(s), sugarLevel)
	}
	return 0
}

func (f FruitPercent) GetPoints(st ScoreType) int {

	if st == Food {
		return getPointsFromRange(float64(f), fruitLevel)
	} else if st == Beverage {
		return 0
	} else if st == Water {
		return 0
	} else if st == Cheese {
		return 0
	}
	return 0
}

func (f FiberGram) GetPoints(st ScoreType) int {

	if st == Food {
		return getPointsFromRange(float64(f), fiberLevel)
	} else if st == Beverage {
		return 0
	} else if st == Water {
		return 0
	} else if st == Cheese {
		return 0
	}
	return 0
}

func (p ProteinGram) GetPoints(st ScoreType) int {

	if st == Food {
		return getPointsFromRange(float64(p), proteinLevel)
	} else if st == Beverage {
		return 0
	} else if st == Water {
		return 0
	} else if st == Cheese {
		return 0
	}
	return 0
}

func (n NutritionalScore) GetNutriScore() string {
	if n.ScoreType == Food {
		return scoreToLetter[n.value]
	}
	if n.ScoreType == Beverage {
		return scoreToLetter[n.value]
	}
	if n.ScoreType == Water {
		return "A"
	}
	if n.ScoreType == Cheese {
		return scoreToLetter[n.value]
	}
	return "A"
}

func GetNutritionalScore(n NutritionalData, st ScoreType) NutritionalScore {
	value := 0
	positive := 0
	negative := 0

	if st != Water {
		fruitPoints := n.Fruits.GetPoints(st)
		fiberPoints := n.Fiber.GetPoints(st)
		negative = n.Energy.GetPoints(st) + n.SaturatedFattyAcid.GetPoints(st) + n.Sodium.GetPoints(st) + n.Sugars.GetPoints(st)
		positive = fruitPoints + fiberPoints + n.Proteins.GetPoints(st)
		if st == Cheese {
			value = negative - positive
		}

	}

	return NutritionalScore{value, positive, negative, st}

}
