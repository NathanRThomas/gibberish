/*! \file intent.go
    \brief Used to compare a statement against an array of training statements to determine the likelyhood of them being a match
*/

package intent

import (
//    "fmt"
//    "strings"
    "math"
    

    "github.com/NathanRThomas/gibberish/toolz"
)

  //-------------------------------------------------------------------------------------------------------------------------//
 //----- PUBLIC STRUCTS ----------------------------------------------------------------------------------------------------//
//-------------------------------------------------------------------------------------------------------------------------//


type Intent_c struct {
    tool        toolz.Toolz_c
}

  //-------------------------------------------------------------------------------------------------------------------------//
 //----- PRIVATE FUNCTIONS -------------------------------------------------------------------------------------------------//
//-------------------------------------------------------------------------------------------------------------------------//

func (i *Intent_c) rngChk (in float64) float64 {
    if in > 1.0 { return 1.0 }
    if in < 0.0 { return 0.0 }
    return in
}

/*! \brief Scores how well this training statement matches the input statement

    The closer the number of words in the training to the number in the statement the better
    also the earlier in the statement we match the better

*/
func (i *Intent_c) score (statementLen, trainingLen, statementIdx int) float64 {
    overlap := float64(trainingLen) / float64(statementLen)     //the percent the training set was to the statement, the higher the better
    start := float64(statementIdx - trainingLen) / float64(statementLen) //how early in the users's statement the training statement showed up, earlier the better
    //fmt.Printf("overlap: %f - start: %f\n", overlap, start)

    //do some hard checks first
    if overlap > 0.95 { return 1.0 }    //this is totally it
    if overlap < 0.1 { return 0.0 }     //this totally isn't it
    if start > 0.8 { return 0.0 }       //this probably isn't it either

    //ok, if we're here, let's use a fancy equation
    res := 1.0 / (1.0 + math.Pow(math.E, -5.0 * (start - 0.6))) + 0.1
    //fmt.Println(res)
    if overlap >= res { //ok, we met our threshold for this counting, used as a gate
        weight := (overlap * 0.5) + ((overlap - res) * 1.5)    //simple weighting, absolute value and the ratio
        //fmt.Println(weight)
        return i.rngChk(weight)
    } else {
        return 0.0  //this is a hard cutoff
    }
}

  //-------------------------------------------------------------------------------------------------------------------------//
 //----- PUBLIC FUNCTIONS --------------------------------------------------------------------------------------------------//
//-------------------------------------------------------------------------------------------------------------------------//

func (i *Intent_c) Classify (statement string, training []string) (rank float64) {

    stemed := i.tool.Stem(statement)    //parse out what the input is
    if len(stemed) == 0 { return 0 }    //nothing to do here

    //now we want to go through the training statements and record the one with the highest correlation

    for _, t := range(training) {
        tStem := i.tool.Stem(t) //parse the training statements the same way
        if len(tStem) > 0 && len(tStem) <= len(stemed) {    //we can't match a training statment if it's longer than what the user typed
            statementIdx := 0
            trainingIdx := 0

            for statementIdx < len(stemed) && trainingIdx < len(tStem) {    //move linearly through what the user typed
                if stemed[statementIdx] == tStem[trainingIdx] { //this is a good thing
                    trainingIdx++
                } else if trainingIdx > 0 && stemed[statementIdx] == tStem[trainingIdx-1] { //this happens if the user repeats themselves, ie please please stop
                    //nothing to do in this case but move along to the next statement idx
                } else {
                    trainingIdx = 0 //we want to reset this
                }
                statementIdx++  //always move through the statement index
            }

            //see why we finished
            if trainingIdx >= len(tStem) {  //we were successful
                //let's see how successful we were
                //fmt.Printf("\n%s\n", strings.Join(tStem, " "))
                if tmp := i.score(len(stemed), len(tStem), statementIdx); tmp > rank {  //this is the best so far
                    if tmp >= 0.99 { return 0.99 }   //this is a really good one!!
                    rank = tmp  //save this as our best so far
                }
            } else {
                continue    //no match for this training statment
            }
        }
    }

    return rank  //we're done
}