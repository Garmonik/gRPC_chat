package validate_lib

import (
	passwordvalidator "github.com/wagslane/go-password-validator"
	"net/mail"
	"regexp"
	"strconv"
)

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidPassword(password string) bool {
	minEntropy := 50.0
	err := passwordvalidator.Validate(password, minEntropy)
	if err != nil {
		return false
	}
	return true
}

func ConversionStringToUint(str string) (uint, error) {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	result := uint(num)
	return result, nil
}

func IsSafeInput(input string) bool {
	patterns := []string{
		// SQL-инъекции
		`(?i)(\bunion\b|\bselect\b|\binsert\b|\bupdate\b|\bdelete\b|\bdrop\b|\btruncate\b|\bexec\b|\bexecute\b|\bcreate\b|\balter\b)`,
		`(\bOR\b|\bAND\b)\s+\d+\s*=\s*\d+`,
		`['"]\s*\+\s*['"]`,
		`--|\/\*|\*\/`,

		// XSS
		`<script.*?>.*?<\/script>`,
		`<.*?on\w+\s*=\s*["'].*?["']`,
		`javascript:\s*`,

		// Shell-инъекции
		`;\s*\w+`,
		`\|\s*\w+`,
		`&\s*\w+`,
		`\$\s*\(`,
		`\`,
	}

	for _, pattern := range patterns {
		matched, err := regexp.MatchString(pattern, input)
		if err != nil {
			return false
		}
		if matched {
			return false
		}
	}

	if len(input) > 3000 {
		return false
	}

	return true
}
