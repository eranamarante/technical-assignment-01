package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
)

type request struct {
	Keywords  []string `json:"keywords"`
	EmailText string   `json:"emailText"`
}

type response struct {
	IsCensored        bool   `json:"isCensored"`
	ModifiedEmailText string `json:"modifiedEmailText"`
}

type errResponse struct {
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	// POST handler
	e.POST("/email/classify", func(c echo.Context) error {
		req := new(request)

		// binds request body to struct
		if err := c.Bind(req); err != nil {
			errResp := errResponse{Message: fmt.Sprintf("error mapping request body: %s", err.Error())}
			return c.JSON(http.StatusBadRequest, errResp)
		}

		isCensored, modifiedEmailText := classifyEmail(req.Keywords, req.EmailText)
		return c.JSON(http.StatusOK, response{IsCensored: isCensored, ModifiedEmailText: modifiedEmailText})
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func classifyEmail(keywords []string, emailTxt string) (bool, string) {
	flag := false
	modifiedEmail := emailTxt

	// loop through keywords
	for i, word := range keywords {

		// skip empty strings
		if word == "" {
			continue
		}

		// trim excess spaces
		keywords[i] = strings.TrimSpace(word)

		// using regexp for more reliable match of words or phrases
		re := regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(keywords[i]) + `\b`)
		modifiedEmail = re.ReplaceAllString(modifiedEmail, "*****")

		// set flag to true if modifiedEmail is not the same as original email
		if !strings.EqualFold(modifiedEmail, emailTxt) {
			flag = true
		}
	}

	return flag, modifiedEmail
}
