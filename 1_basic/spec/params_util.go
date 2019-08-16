package main

import (
	"bytes"
	"fmt"

	params "github.com/aimzeter/wuts/1_basic"
)

func main() {
	GetStudentIDSpec()
}

func GetStudentIDSpec() {
	body := bytes.NewBufferString(`
		{
			"student_id": 1
		}
	`)

	id, err := params.GetStudentID(body)

	if err != nil {
		fmt.Printf("❌ FAIL ❌: GetStudentID should not return error, got error '%s'\n", err.Error())
		return
	}

	if id != 1 {
		fmt.Printf("❌ FAIL ❌: GetStudentID did not return correct id.\n"+
			"\twant:\t%d\n"+
			"\tgot:\t%d\n", 1, id)
		return
	}

	fmt.Printf("✅ PASS ✅: GetStudentID return expected output\n")
}
