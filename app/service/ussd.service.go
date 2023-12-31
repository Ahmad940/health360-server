package service

import "regexp"

func StartUSSD(sessionId, serviceCode, phoneNumber, text string) string {
	// if text == "" {
	// 	// This is the first request. Note how we start the response with CON
	// 	return `CON What would you like to check
	//       1. My account
	//       2. My phone number`
	// }
	// return "Hello World"
	if text == "" {
		// This is the first request. Note how we start the response with CON
		return `CON Welcome to health360 USSD service, please select an option below:
		1. Book appointment
		2. Cancel appointment`
	} else if text == "1" {
		// Business logic for first level response
		return `CON Select category
		1. Cardiology
		2. Hair Issue
		3. Generic CheckUp
		4. Optician
		5. Dermatology`
		// } else if isValid, _ := regexp.MatchString(`^1\*(\d)$`, text); isValid {
	} else if isValid, _ := regexp.MatchString(`^1\*[1-6]$`, text); isValid {
		return `CON Top consultants:
		1. Aaliyah Junaid
		2. Simon Adepetoye
		3. Abdulmalik Ibrahim
		4. Ahmad Muhammad
		5. Precios Victor
		6. More`
	} else if isValid, _ := regexp.MatchString(`^1\*[1-6]\*[1-2]$`, text); isValid { // 1*[1-5]*[1-2]
		return `CON Top consultants:
		1. See Details
		2. Book (₦1000)`
	} else if isValid, _ := regexp.MatchString(`^1\*[1-6]\*[1-2]\*[1-2]$`, text); isValid { // 1*[1-5]*[1-2]
		return `CON Consultant detail: 
		Dr Junaid is an Oxford graduate in cardiology, he has been practicing for 15 years and has dozen record to his belt.
		1. ₦Book`
	} else if isValid, _ := regexp.MatchString(`^1\*[1-6]\*[1-2]\*[1-2]\*[1-2]$`, text); isValid { // 1*[1-5]*[1-2]
		return `END Not available at the moment, Coming soon`
	}
	return "END Hello World"
}
