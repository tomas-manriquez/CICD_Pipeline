package utils


/*
Find path string from command line arguments
Input: flag input, cargs
Output: string
*/
func ParseArgs(path string, args []string) string {
	if path != "" {
		return path
	} else {
		//path not passed via flag
		//check if exists on args
		if len(args) != 1 {
			return path //return ""
		} else {
			return args[0]
		}
	}
}
