package internal_errors

import "errors"

var ErrNoCategory = errors.New("category with provided id is not found")