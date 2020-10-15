package logging

import "log"

func PrintEror(err error) {
	if err != nil {
		log.Println(err)
	}
}

func PrintErrorWithMessage(err error, message string) {
	if err != nil {
		log.Println(message, ":", err)
	}
}

func FatalEror(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
