# Book Data

The `data` directory contains data required for `booklist.Service` and
 `bookdetails.Service`.
 
 Each `.txt` file here represents a book.
 
 And a book must be included in `_booklist.csv` in order to be available from
  the API.
  
 These files are **not used directly**, instead they are compiled into `data.go` by `compile-data.sh`. **Whenever a change is made to a file here, run
  `compile-data.sh` from this directory in order to update `data.go`.**
  