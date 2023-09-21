package databaseTable_test

import (
	"github.com/alecthomas/units"
	structureSpec "github.com/taubyte/go-specs/structure"
	databaseTable "github.com/taubyte/tau-cli/table/database"
)

func ExampleQuery_key() {
	database := &structureSpec.Database{
		Id:          "QmbAA8hRosp5BaXFXikADCtpkQCgQCPdRVhnxjiSHfXdWH",
		Name:        "someProject",
		Description: "this is a database of some type",
		Tags:        []string{"apple", "orange", "banana"},
		Match:       "/test/v1",
		Regex:       false,
		Local:       false,
		Key:         "someKey",
		Min:         15,
		Max:         30,
		Size:        uint64(units.MB),
	}

	databaseTable.Query(database)

	// Output:
	// ┌─────────────┬────────────────────────────────────────────────┐
	// │ ID          │ QmbAA8hRosp5BaXFXikADCtpkQCgQCPdRVhnxjiSHfXdWH │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Name        │ someProject                                    │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Description │ this is a database of some type                │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Tags        │ apple, orange, banana                          │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Encryption  │ true                                           │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Access      │                                                │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │  -  Network │ all                                            │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Replicas    │                                                │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │  -  Min     │ 15                                             │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │  -  Max     │ 30                                             │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Storage     │                                                │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │  -  Size    │ 1MB                                            │
	// └─────────────┴────────────────────────────────────────────────┘
}

func ExampleQuery_no_key() {
	database := &structureSpec.Database{
		Id:          "QmbAA8hRosp5BaXFXikADCtpkQCgQCPdRVhnxjiSHfXdWH",
		Name:        "someProject",
		Description: "this is a database of some type",
		Tags:        []string{"apple", "orange", "banana"},
		Match:       "/test/v1",
		Regex:       false,
		Local:       false,
		Min:         15,
		Max:         30,
		Size:        uint64(units.MB),
	}

	databaseTable.Query(database)

	// Output:
	// ┌─────────────┬────────────────────────────────────────────────┐
	// │ ID          │ QmbAA8hRosp5BaXFXikADCtpkQCgQCPdRVhnxjiSHfXdWH │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Name        │ someProject                                    │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Description │ this is a database of some type                │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Tags        │ apple, orange, banana                          │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Encryption  │ false                                          │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Access      │                                                │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │  -  Network │ all                                            │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Replicas    │                                                │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │  -  Min     │ 15                                             │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │  -  Max     │ 30                                             │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │ Storage     │                                                │
	// ├─────────────┼────────────────────────────────────────────────┤
	// │  -  Size    │ 1MB                                            │
	// └─────────────┴────────────────────────────────────────────────┘
}


/* 
Imports:
The code brings in some external tools and packages to help with the testing process. For instance:
It uses a package called "github.com/alecthomas/units" to work with data sizes, like converting them from bytes to kilobytes or megabytes.
There's another package called "github.com/taubyte/go-specs/structure" which seems to be custom-made for describing data structures, possibly for databases.
"github.com/taubyte/tau-cli/table/database" is likely a package that contains functions for interacting with database tables.
Lastly, it utilizes the standard Go package called "testing" to create and run tests.

Example Functions:
There are two example functions: ExampleQuery_key and ExampleQuery_no_key. These functions serve as practical demonstrations.
ExampleQuery_key shows how to use the databaseTable.Query function when dealing with a database that has a specific key.
ExampleQuery_no_key demonstrates how to use the same function when the database doesn't have a key.

Example Data:
Inside each example function, there's a pretend database configuration represented by an object named structureSpec.Database. This object contains details like the database's ID, name, description, tags, and size.
It's as if we're setting up an example database with certain characteristics.

Function Calls:
Both example functions make use of the databaseTable.Query function. This function is responsible for querying and formatting information about a database.

Expected Output:
Each example function includes a comment block labeled Output:. This comment section describes what we expect the databaseTable.Query function to produce when we call it with the example database data.

Testing Framework Usage:
These example functions are designed to work with Go's built-in testing framework. When you run tests for this code, the testing framework will execute these functions and check if the actual output matches the expected output specified in the Output: comments.
If everything aligns as expected, the test passes. If there's a mismatch, it means something isn't working correctly, and the test fails.
In essence, this code is a way to verify that the databaseTable.Query function behaves correctly when used with different database scenarios. It's like having a practice run to make sure everything works as intended before using it in real applications.
*/