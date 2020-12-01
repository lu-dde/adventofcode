package alch

import (
	"fmt"
	"math"
)

//Synthesis struct
type Synthesis struct {
	goalProduct string
	Reactions   map[string]*Reaction
}

//New synthesis
func New(goal string) *Synthesis {
	return &Synthesis{
		goalProduct: goal,
		Reactions:   map[string]*Reaction{},
	}
}

//AddLine parse and build Reactions
func (synth *Synthesis) AddLine(line string) {
	reaction := parseReactionStr(line)
	synth.Reactions[reaction.Product.Element] = &reaction
}

//FuelCost calculates cost of goal in ORE
func (synth *Synthesis) FuelCost() int {
	wishList := synth.MaterialCost("", 1, synth.Reactions[synth.goalProduct])

	fmt.Println()
	fmt.Println("Wish list")
	sum := make(map[string]int, 10)
	for _, material := range wishList {
		sum[material.Element] += material.Amount
	}

	var cost float64 = 0
	for m, goalAmountInt := range sum {
		goalAmount := float64(goalAmountInt)
		costPerAmount := float64(synth.Reactions[m].Reagents[0].Amount)
		amountPerReaction := float64(synth.Reactions[m].Product.Amount)
		fmt.Println(m, goalAmount, "at", costPerAmount, "ORE", "per", amountPerReaction, m)
		needed := math.Ceil(goalAmount / amountPerReaction)
		cost += needed * costPerAmount
		fmt.Println(needed*costPerAmount, "ORE", "=>", needed*amountPerReaction, m)

	}
	fmt.Println()
	return int(cost)
}

//MaterialCost calculates cost of reaction in ORE
func (synth *Synthesis) MaterialCost(level string, amount int, r *Reaction) []Material {

	wishList := []Material{}

	levelPrint := level
	fmt.Println(levelPrint, amount, "of", r.Product.Amount, r.Product.Element)
	for _, agent := range r.Reagents {
		if agent.Element == "ORE" {
			fmt.Println(level, "add to wish list:", amount, "of", r.Product.Element)
			//wishList = append(wishList, Material{r.Product.Element, amount})
		} else {
			element := synth.Reactions[agent.Element]
			amountMaterialNeeded := int(math.Ceil(float64(amount) / float64(agent.Amount)))
			amountProductNeeded := element.Product.Amount * amountMaterialNeeded
			wishes := synth.MaterialCost(level+"  ", amountProductNeeded, element)
			wishList = append(wishList, wishes...)
		}
	}

	return wishList
}
