package store

import "../database"

var UrlPatterns = []database.Path {
	{"/store/", StoreController},
}