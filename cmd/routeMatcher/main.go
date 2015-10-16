package main
import (
	"regexp"
	"fmt"
	"mime"
)

func main() {
	fmt.Println(mime.TypeByExtension(".html"))

	str := "/example/url/{something}/else"

	fmt.Println(str)

	reg, _ := regexp.Compile("{.*}")

	regexdString := reg.ReplaceAllString(str, "(.*)")

	reg, _ = regexp.Compile(regexdString)

//	fmt.Println("Matches: " + string(reg.Match([]byte("/example/url/if/else"))))
//	fmt.Println("Sub: " + string(reg.Match([]byte("/example/url/if/else"))))

	fmt.Println(regexdString)

}