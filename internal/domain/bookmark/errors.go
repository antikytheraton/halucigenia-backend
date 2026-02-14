package bookmark

import "errors"

var ErrBookmarkSaveFailed = errors.New("failed to save bookmark")
var ErrBookmarkNotFound = errors.New("bookmark not found")
var ErrInvalidBookmarkID = errors.New("invalid bookmark ID")
var ErrBookmarkAlreadyExists = errors.New("bookmark already exists")
var ErrBookmarkDeleteFailed = errors.New("failed to delete bookmark")
