Credit card validator

This is a simple html site that turns to the go server, giving it the number that the user has input in the input element in the html file. Then the go server checks if the number is a valid credit card number, puts the result into a struct and then writes a response back to the html giving it the calculated output as a JSON. The html then receives the JSON and reads the result value. Then according to the result displays on the screen one of the following responses:
1. This is a valid credit card number.
2. This is not a valid credit card number.