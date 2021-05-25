package main

import (
	"bufio"
	"fmt"
	"gitlab.com/bzholmyrza/hw3_bench/user"
	"io"
	"os"
	"strings"
)

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	seenBrowsers := make(map[string]bool)
	sc := bufio.NewScanner(file)
	user := user.User{}

	for i := 0; sc.Scan(); i++ {
		err := user.UnmarshalJSON([]byte(sc.Text()))
		if err != nil {
			panic(err)
		}

		isAndroid := false
		isMSIE := false

		for _, browser := range user.Browsers {
			if strings.Contains(browser, "Android") {
				isAndroid = true
				_, notSeenBefore := seenBrowsers[browser]
				if !notSeenBefore {
					seenBrowsers[browser] = true
				}
			}

			if strings.Contains(browser, "MSIE") {
				isMSIE = true
				_, notSeenBefore := seenBrowsers[browser]
				if !notSeenBefore {
					seenBrowsers[browser] = true
				}
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		email := strings.Replace(user.Email, "@", " [at] ", -1)
		fmt.Fprintf(out, "[%d] %s <%s>\n", i, user.Name, email)
	}
	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}
