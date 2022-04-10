package internal_test

import (
	"fmt"
	"github.com/lapacek/simple-api-example/internal/common"
	"testing"
	"time"

	"github.com/lapacek/simple-api-example/internal"
)

func Test_GetStartOfWeek(t *testing.T) {

	input := "2022-04-13"
	date, err := time.Parse(common.DateLayout, input)
	if err != nil {
		fmt.Printf("Time parsing failed, err(%v)", err)
	}

	result := internal.GetStartOfWeek(date).Format(common.DateLayout)

	expected := "2022-04-11"
	if expected != result {
		t.Errorf("Computed date(%v), expected date(%v)", result, expected)
	}
}

func Test_GetEndOfWeek(t *testing.T) {

	input := "2022-04-13"
	date, err := time.Parse(common.DateLayout, input)
	if err != nil {
		fmt.Printf("Time parsing failed, err(%v)", err)
	}

	result := internal.GetEndOfWeek(date).Format(common.DateLayout)

	expected := "2022-04-17"
	if expected != result {
		t.Errorf("Computed date(%v), expected date(%v)", result, expected)
	}
}

func Test_IsTicketAvailable(t *testing.T) {

}
