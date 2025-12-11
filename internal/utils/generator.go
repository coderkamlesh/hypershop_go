package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	lastTimestamp int64
	sequence      uint32
	mu            sync.Mutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// ==================== ID GENERATORS ====================

// GenerateID generates unique, sortable, collision-free ID
// Example: ORD173564123456789012 (21 characters)
func GenerateID(prefix string) string {
	mu.Lock()
	defer mu.Unlock()

	timestamp := time.Now().UnixMilli()

	if timestamp == lastTimestamp {
		sequence++
		if sequence > 9999 {
			time.Sleep(time.Millisecond)
			timestamp = time.Now().UnixMilli()
			sequence = 0
		}
	} else {
		sequence = 0
		lastTimestamp = timestamp
	}

	return fmt.Sprintf("%s%d%04d", prefix, timestamp, sequence)
}

// User related IDs
func GenerateUserID() string   { return GenerateID("USR") }
func GenerateSellerID() string { return GenerateID("SLR") }
func GenerateRiderID() string  { return GenerateID("RDR") }

// Product related IDs
func GenerateProductID() string  { return GenerateID("PRD") }
func GenerateCategoryID() string { return GenerateID("CAT") }
func GenerateBrandID() string    { return GenerateID("BRD") }

// Order related IDs
func GenerateOrderID() string   { return GenerateID("ORD") }
func GenerateCartID() string    { return GenerateID("CRT") }
func GeneratePaymentID() string { return GenerateID("PAY") }

// Inventory related IDs
func GenerateInventoryID() string { return GenerateID("INV") }
func GenerateWarehouseID() string { return GenerateID("WHS") }
func GenerateStoreID() string     { return GenerateID("STR") }
func GenerateBatchID() string     { return GenerateID("BAT") }

// Other IDs
func GenerateAddressID() string      { return GenerateID("ADR") }
func GenerateReviewID() string       { return GenerateID("REV") }
func GenerateCouponID() string       { return GenerateID("CPN") }
func GenerateNotificationID() string { return GenerateID("NOT") }

// ==================== OTP GENERATORS ====================

// GenerateOTP generates 6-digit numeric OTP
func GenerateOTP() string {
	otp := rand.Intn(900000) + 100000
	return fmt.Sprintf("%d", otp)
}

// GenerateOTPWithLength generates OTP of custom length
func GenerateOTPWithLength(length int) string {
	digits := "0123456789"
	otp := ""
	for i := 0; i < length; i++ {
		otp += string(digits[rand.Intn(len(digits))])
	}
	return otp
}

// GenerateAlphanumericOTP generates alphanumeric OTP
func GenerateAlphanumericOTP(length int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	otp := ""
	for i := 0; i < length; i++ {
		otp += string(chars[rand.Intn(len(chars))])
	}
	return otp
}

// ==================== TOKEN GENERATORS ====================

// GenerateReferralCode generates unique referral code
// Example: REF-ABC123
func GenerateReferralCode() string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := ""
	for i := 0; i < 6; i++ {
		code += string(chars[rand.Intn(len(chars))])
	}
	return "REF-" + code
}

// GeneratePromoCode generates promo code
// Example: PROMO-XYZ789
func GeneratePromoCode() string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := ""
	for i := 0; i < 6; i++ {
		code += string(chars[rand.Intn(len(chars))])
	}
	return "PROMO-" + code
}

// GenerateTrackingID generates shipment tracking ID
// Example: TRACK-173564123456789
func GenerateTrackingID() string {
	timestamp := time.Now().UnixMilli()
	return fmt.Sprintf("TRACK-%d", timestamp)
}

// ==================== RANDOM GENERATORS ====================

// GenerateRandomString generates random alphanumeric string
func GenerateRandomString(length int) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := ""
	for i := 0; i < length; i++ {
		result += string(chars[rand.Intn(len(chars))])
	}
	return result
}

// GenerateRandomNumber generates random number between min and max
func GenerateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}
