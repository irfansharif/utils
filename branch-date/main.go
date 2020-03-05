// Prints out the current date in the format "YYMMDD".
//
// Usage:
//  $ git co -b `branch-date`.<branch-name>
//  $ git branch
//    190925.kv-contention
//    191110.raftlog-pkg
//    200129.trunc-panic
//    200130.snapshot-log-entries
//    200208.workload-region-failure
//    200212.dan/migration
//    200212.migration
//    200228.make-faster
//    200304.short-running-hooks
//    200305.gitmodules
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("%02d%02d%02d", t.Year()-2000, t.Month(), t.Day())
}
