/*
	Package largo implements command-line-like string/[]string parsing.
	Unlike standard `flag` package, largo allows to provide flag args
	in **unordered** way, such as:

		dosomething -upper -count 3 say hello
		dosomething -upper say -count=3 hello
		dosomething say -upper hello -count 3

	All should be parsed and resulted to the same args, such as:

		Command:	dosomething
		Rest args:	[say, hello]
		Flags:
			upper:	true
			count:	3

	`largo` parses given string/[]string as follows:

		dosomething say -count 3 hello -upper
		-----------
		  command
		   	            -------        ------
		                 flag           flag
		            ----       --------
		            rest         rest

	// TODO: Add more information here in near future ;)
*/
package largo
