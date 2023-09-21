package validate

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/taubyte/tau-cli/common"
	"github.com/taubyte/tau-cli/constants"
)

func SliceContains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}

	return false
}

func VariableNameValidator(val string) error {
	if val != "" {
		err := matchAllString(val, NameRegex)
		if err != nil {
			return err
		}
		if len(val) == 0 || len(val) > 250 {
			return errors.New(Between0And250)
		}
	}

	return nil
}

func VariableDescriptionValidator(val string) error {
	if val != "" {
		if len(val) > 250 {
			return errors.New(GreaterThan250)
		}
	}

	return nil
}

func VariableTagsValidator(val []string) error {
	// TODO validate

	// Example:
	// for _, v := range val{
	// 	  err := matchAllString(val, TagRegex)
	// }
	return nil
}

func VariablePathValidator(path string) error {
	if path != "" {
		if !strings.HasPrefix(path, "/") {
			return fmt.Errorf(PathMustStartWithSlash, path)
		}
	}

	// TODO REGEX
	return nil
}

func VariableTime(val string) error {
	if val != "" {
		_, err := time.ParseDuration(val)
		if err != nil {
			return fmt.Errorf(InvalidTimeInput, val, err)
		}
	}

	return nil
}

func VariableBool(val string) error {
	if val != "" {
		if _, err := strconv.ParseBool(val); err != nil {
			return fmt.Errorf(InvalidBoolInput, val, err)
		}
	}

	return nil
}

func VariableRequiredValidator(val string) error {
	if val != "" {
		if len(val) == 0 || len(val) > 250 {
			return errors.New(Between0And250)
		}
	}

	return nil
}

func VariableProviderValidator(val string) error {
	if val != "" {
		// TODO: add || gitlab || bitbucket when implemented
		if strings.ToLower(val) != "github" {
			return fmt.Errorf(ProviderNotSupported, val)
		}
	}

	return nil
}

func VariableIntValidator(val string) error {
	if val != "" {
		if _, err := strconv.Atoi(val); err != nil {
			return fmt.Errorf(InvalidIntegerValue, val, err)
		}
	}

	return nil
}

func SizeUnitValidator(val string) error {
	if val != "" {
		if !SliceContains(common.SizeUnitTypes, strings.ToUpper(val)) {
			return fmt.Errorf(InvalidSizeUnit, val, common.SizeUnitTypes)
		}
	}

	return nil
}

func FQDNValidator(val string) error {
	if val != "" {
		if !govalidator.IsDNSName(val) {
			return fmt.Errorf(InvalidFqdn, val)
		}
	}

	return nil
}

func RequiredNoCharLimit(val string) error {
	if val != "" {
		if len(val) == 0 {
			return errors.New(NoEmpty)
		}
	}

	return nil
}

func ApiMethodValidator(val string) error {
	if val != "" {
		if !SliceContains(common.HTTPMethodTypes, strings.ToLower(val)) {
			return fmt.Errorf(InvalidMethodType, val, common.HTTPMethodTypes)
		}
	}

	return nil
}

func MethodTypeValidator(val string) error {
	if val != "" {
		if !SliceContains(common.FunctionTypes, strings.ToLower(val)) {
			return fmt.Errorf(InvalidApiMethodType, val, common.FunctionTypes)
		}
	}

	return nil
}

func CodeTypeValidator(val string) error {
	var types = constants.CodeTypes
	if val != "" {
		if !SliceContains(types, strings.ToLower(val)) {
			return fmt.Errorf(InvalidCodeType, val, types)
		}
	}

	return nil
}

func BucketTypeValidator(val string) error {
	if val != "" {
		if !SliceContains(common.BucketTypes, val) {
			return fmt.Errorf(InvalidBucketType, val, common.BucketTypes)
		}
	}

	return nil
}

func VariableSizeValidator(val string) error {
	if val != "" {
		if !IsAny(val, IsInt, IsBytes) {
			return fmt.Errorf(InvalidSize, val)
		}
	}

	return nil
}


/*  Imports:
The code brings in various packages to help with different tasks. It's like assembling a team of experts with unique skills to assist in our work.
We have some trusted members from the standard Go library like "fmt," "errors," and "strings."
Additionally, we've invited external experts like "github.com/asaskevich/govalidator."
We've also got some special members from our own "github.com/taubyte/tau-cli/common" and "github.com/taubyte/tau-cli/constants" repositories who provide essential tools and constants for our validation functions.

Validation Functions:
We've prepared a set of validators, each with a specific job to ensure that the data we work with meets certain rules and standards.
For example, the "VariableNameValidator" makes sure that variable names are both valid according to a regular pattern and fall within an acceptable length.
The "VariableDescriptionValidator" checks the length of variable descriptions to ensure they are not too long.
We also have a placeholder, "VariableTagsValidator," waiting for specific validation rules to be added in the future.
These validators act as guardians to our data, making sure it's in good shape.

Error Handling:
If any of these validators discover that the data doesn't meet the rules, they don't panic; they simply raise their hand and say, "I found a problem!" They do this by returning an error message that helps us understand what went wrong.
This way, if something doesn't match our expectations, we know exactly why.

Constants and Helper Functions:
To help with these validations, we have constants and helper functions. Think of them as reference materials or tools that the validators use to check things.
For instance, if we need to know what an acceptable size unit is, we can consult our "common.SizeUnitTypes" reference.
Or, when we need to check if a Fully Qualified Domain Name (FQDN) is valid, we have a function from an external expert package, "govalidator.IsDNSName," to assist us.
In a nutshell, this code represents a team effort to ensure that our data is checked, validated, and maintained according to specific rules and standards, with clear error messages to guide us when something goes wrong.
*/