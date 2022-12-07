package main

import (
	"fmt"
	"learningGo/errorHandlingInGo"
)

func main() {
	/*** ifElseIf ***/
	//gettingStartedWithGo.FindEvenOdd(20)
	//largestNo := gettingStartedWithGo.ToCheckGraterOfThreeNo(4, 12, 18)
	//fmt.Println("the largest no is", largestNo)
	//gettingStartedWithGo.ToCheckIfNoIsInGivenRange(1000)

	///*** loopsInGo ***/
	//
	//gettingStartedWithGo.ToPrintNoFrom1to10()
	//gettingStartedWithGo.ToPrintNoFrom1to5()
	//gettingStartedWithGo.ToPrintOnlyOddNos()
	//gettingStartedWithGo.ToPrintStarPattern()
	//gettingStartedWithGo.ToShowOuterBreak()
	//
	///*** additional info
	//-> it supports multiple initialization
	//-> multiple update
	//-> for{} can be used for infinite loop
	//***/
	//
	///** switches in go **/
	//gettingStartedWithGo.FindingTheFinger()
	//gettingStartedWithGo.ConditionLessSwitches(80)
	//gettingStartedWithGo.FallthroughInSwitch(101)
	//

	///*** arrays in go ***/
	//collectionInGo.MultipleWayOfInitializing()
	//collectionInGo.CreateAndPrintMultiDimensionalArray(5, 5)
	//
	////collectionInGo.SlicesInGo()
	//collectionInGo.AppendFunctionAndCapacityOfSlices()
	//collectionInGo.SamethingUsingSlice(20, []int{1, 5, 6, 8, 20})
	//collectionInGo.VariadicFunctionAndRelationWithSlice(20, 1, 5, 6, 8, 20)
	//
	////to call an empty slice
	//collectionInGo.SamethingUsingSlice(20, []int{})
	//collectionInGo.VariadicFunctionAndRelationWithSlice(20)
	//
	////slice can be passed to the variadic parameter if we some how tell that i am passing an slice you do't convert it
	//nums := []int{1, 2, 3, 4, 89, 20, 40, 60}
	//collectionInGo.VariadicFunctionAndRelationWithSlice(20, nums...)
	//
	//welcome := []string{"hello", "world"}
	//collectionInGo.Change(welcome...)
	//fmt.Println(welcome)

	/*** maps in go ***/
	//collectionInGo.BasicsOfMap()
	//collectionInGo.MapsWithStructure()

	/*** concurrency in go ***/

	errorHandlingInGo.CalculateAreaOfRectangle()
	fmt.Println("Main function terminated")

}
