/*! \file toolz.go
    \brief Re-used code/classes throughout the gibberish project
*/

package toolz 

import (
    "strings"
    "regexp"

    "github.com/reiver/go-porterstemmer"
)


  //-------------------------------------------------------------------------------------------------------------------------//
 //----- PUBLIC STRUCTS ----------------------------------------------------------------------------------------------------//
//-------------------------------------------------------------------------------------------------------------------------//

type Toolz_c struct {}

  //-------------------------------------------------------------------------------------------------------------------------//
 //----- PRIVATE FUNCTIONS -------------------------------------------------------------------------------------------------//
//-------------------------------------------------------------------------------------------------------------------------//

func (t *Toolz_c) safeStr (str string) string {
    withSpace := regexp.MustCompile (`[^a-zA-Z0-9 ]`)
    noSpace := regexp.MustCompile (`[-']`)
    return strings.ToLower(withSpace.ReplaceAllString(noSpace.ReplaceAllString(str, ""), " "))
}

  //-------------------------------------------------------------------------------------------------------------------------//
 //----- PUBLIC FUNCTIONS --------------------------------------------------------------------------------------------------//
//-------------------------------------------------------------------------------------------------------------------------//

/*! \brief This will generate an array of strings with the base stem of all the words, cleaned and lower-cased
*/
func (t *Toolz_c) Stem (str string) []string {
    ret := make([]string, 0)

    for _, s := range(strings.Fields(t.safeStr(str))) {
        ret = append(ret, porterstemmer.StemString(s)) //get the stem of this word
    }

    return ret
}

