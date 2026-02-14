package bookmark

import "errors"

var ErrBookmarkNotFound = errors.New("bookmark not found")
var ErrInvalidBookmarkID = errors.New("invalid bookmark ID")
var ErrBookmarkAlreadyExists = errors.New("bookmark already exists")
var ErrBookmarkDeleteFailed = errors.New("failed to delete bookmark")
