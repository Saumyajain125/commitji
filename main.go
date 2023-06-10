package main
import "os/exec"
import "bufio"
import "os"
import "github.com/fatih/color"
import "strings"
func main() {
	out, error := exec.Command("git","status").Output();
	if error != nil {
		color.Red("Error: Not a git repository")
		os.Exit(1)
	}
	if !strings.Contains(string(out),"Changes to be committed") {
		color.Red("No files added to staging!")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(os.Stdin)
	color.Blue("Welcome to Commit Message Generator")
	color.Yellow("Select the type of change you are committing: ")
    scanner.Scan()
    name := scanner.Text()
    color.Yellow("What is the scope of this change? (class or file name):  ")
    scanner.Scan()
    scope := scanner.Text()
	color.Yellow("Write a short and imperative summary of the code changes: ")
    scanner.Scan()
    message := scanner.Text()
	if len(message) == 0 {
		color.Red("Error: Commit Message cannot be empty")
		os.Exit(1)
	}
	final := name + "[" + scope + "]" + ": " + message
	_, err := exec.Command("git","commit","-m",final).Output()
	if err != nil {
		color.Red("Error: Commit Failed, Consider Staging the files first")
	}
	if err == nil {
		color.Green("\nCommand Successfully Executed")
	}
}