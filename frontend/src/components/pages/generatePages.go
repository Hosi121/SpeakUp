package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	list := []string{
		"alert",
		"settings",
		"login",
		"home",
		"sessionList",
		"waiting",
		"micCheck",
		"session",
		"sessionInterval",
		"sessionFinished",
		"signUp",
		"record",
		"memo",
		"message",
	}

	for _, s := range list {
		//title := capitalizeFirstLetter(s)
		//content := cmp(s)
		//saveStringToFile(content, fmt.Sprintf("%s.tsx", title))
		fmt.Print(router(s))
	}
}

func cmp(s string) string {
	title := capitalizeFirstLetter(s)
	res := `import { Stack } from "@mui/system";` + "\n\n"
	res += fmt.Sprintf("export const %s = () => {\n", title)
	res += `    return (` + "\n"
	res += `        <Stack>` + "\n"
	res += fmt.Sprintf("            %s\n", s)
	res += `        </Stack>` + "\n"
	res += `    )` + "\n"
	res += `}` + "\n"
	return res
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func saveStringToFile(content, filename string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func router(s string) string {
	runes := []rune(s)
	runes2 := []rune{}
	for _, r := range runes {
		runes2 = append(runes2, unicode.ToLower(r))
	}
	url := string(runes2)

	com := capitalizeFirstLetter(s)
	res := "{\n"
	res += "path: " + fmt.Sprintf("\"%s\",\n", url)
	res += fmt.Sprintf("element: <%s />\n", com)
	res += "},\n"
	return res
}
