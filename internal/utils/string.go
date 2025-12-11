package utils

import "strings"

// MaskMobile masks mobile number
// Example: 9876543210 -> 98****3210
func MaskMobile(mobile string) string {
	if len(mobile) < 4 {
		return mobile
	}
	return mobile[:2] + "****" + mobile[len(mobile)-4:]
}

// MaskEmail masks email address
// Example: john@example.com -> jo**@example.com
func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 || len(parts[0]) < 2 {
		return email
	}

	username := parts[0]
	domain := parts[1]

	if len(username) <= 2 {
		return username[:1] + "*@" + domain
	}

	return username[:2] + "**@" + domain
}
