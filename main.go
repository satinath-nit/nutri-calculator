package main

import fmt "fmt"

func main() {

	ns := GetNutritionalScore(NutritionalData{
		Energy:             EnergyKcal(50),
		Sugars:             SugarGram(105),
		SaturatedFattyAcid: SaturatedFattyAcid(10),
		Sodium:             SodiumMg(10),
		Fruits:             FruitPercent(10),
		Fiber:              FiberGram(105),
		Proteins:           ProteinGram(10),
	}, Food)

	fmt.Println("HNutritional Score %d", ns.value)
	fmt.Println("Nutriscore %d", ns.GetNutriScore())

}
