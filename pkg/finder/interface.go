/* The finder interface defines the interface for a book finder implementation*/
package finder

import "github.com/thealamu/bookfinder/internal/pkg/bookdetails"

type Finder interface {
	Find(s string) *bookdetails.BookDetails
}
