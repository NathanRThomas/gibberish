/*! \file gibberish.go
    \brief Main source file for the gibberish project
    This is an attempt to do basic nlp


    This file is used only for individual testing
    When used as a module for other projects you'd want to directly import from the sub-directories of interest
*/

package main 

import (
    "fmt"

    "github.com/NathanRThomas/gibberish/intent"
)


func testy (statement string, trainings []string, minValue, maxValue float64) bool {
    gibberish := intent.Intent_c{}
    res := gibberish.Classify(statement, trainings)

    if minValue <= res || maxValue >= res {
        return true
    } else {
        fmt.Printf("ERROR: '%s' expecting %f but got %f\n", statement, minValue, res)
        return false
    }
}


  //-------------------------------------------------------------------------------------------------------------------------//
 //----- PUBLIC FUNCTIONS --------------------------------------------------------------------------------------------------//
//-------------------------------------------------------------------------------------------------------------------------//

func main () {
    
    training := make([]string, 0)
    training = append(training, "no")
    training = append(training, "nope")
    training = append(training, "stop")
    training = append(training, "no thanks")
    training = append(training, "please stop")
    training = append(training, "block")
    training = append(training, "unsubscribe")
    training = append(training, "don't message me")
    training = append(training, "stop messaging me")
    training = append(training, "shutup")
    training = append(training, "shut up")

    //negatives
    testy("no", training, 0.99, 1.1)
    testy("stop! no thanks", training, 0.5, 1.1)
    testy("No thanks, please stop", training, 0.5, 1.1)
    testy("Please Don't message me", training, 0.99, 1.1)
    testy("Stop it now", training, 0.51, 1.1)

    //positives
    testy("yes", training, 0.00, 0.4)
    testy("yes please", training, 0.00, 0.4)
    testy("thanks", training, 0.00, 0.4)
    testy("yep", training, 0.00, 0.4)
}