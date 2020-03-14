/*
Given a name, return a string with the message:

One for X, one for me.
Where X is the given name.

However, if the name is missing, return the string:

One for you, one for me.
Here are some examples:

Name	String to return
Alice	One for Alice, one for me.
Bob	One for Bob, one for me.
One for you, one for me.
Zaphod	One for Zaphod, one for me.
*/

package twofer

import "fmt"

func ShareWith(name string) string {
	//assign a value if no param was passed
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}

func twofer() {
	fmt.Println(ShareWith("Alice"))
}
