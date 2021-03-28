package accounts

import "../database"

var UrlPatterns = []database.Path {
	{"/login/", RedirectSignedInUser(LoginController)},
	{"/signup/", RedirectSignedInUser(SignUpController)},
	{"/logout/", IsAuthorized(LogoutController)},
	{"/dashboard/", IsAdminUser(DashboardController)},
}
