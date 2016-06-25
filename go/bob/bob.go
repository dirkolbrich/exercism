package bob

// import "strings"
import "regexp"

const testVersion = 2

func Hey(s string) string{
    var answer [4]string 
    answer[0] = "Whatever."
    answer[1] = "Whoa, chill out!"
    answer[2] = "Sure."
    answer[3] = "Fine. Be that way!"

    switch {
    // find if all words are upper case
    case regexp.MustCompile(`\A[^a-z]+\z`).MatchString(s) == true:
        return answer[1]                
    // find question mark at end of string
    case regexp.MustCompile(`\?\z`).MatchString(s) == true:
        return answer[2]        
    // string is empty or contains only whitespace
    case regexp.MustCompile(`\s+^\S`).MatchString(s) == true:
        return answer[3]        
    default:
        return answer[0]        
    }

}