/* containing library to calculate how long a person has lived*/
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	
	return t.Add(time.Duration(1000000000)*time.Second)
}
