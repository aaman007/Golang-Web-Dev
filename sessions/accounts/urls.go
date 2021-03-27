package accounts

import "../database"

var UrlPatterns = []database.Path {
	{"/login/", LoginController},
	{"/signup/", SignUpController},
	{"/logout/", LogoutController},
	{"/dashboard/", DashboardController},
}
