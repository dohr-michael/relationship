package validators

import (
	"github.com/asaskevich/govalidator"
	"github.com/dohr-michael/relationship/apis/tools/models"
)

func init() {
	// Validator for OptionString
	govalidator.CustomTypeTagMap.Set("nullOrNonEmpty", func(i interface{}, o interface{}) bool {
		first, ok := i.(models.OptionalString)
		return ok && (first.Value == nil || len(*first.Value) > 0)
	})
}