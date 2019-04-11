/*
Package auth provides boring user authentication code for golang.

Because I'm tired of writing the same things over and over again.

Provides a complete user authentication system, including:

1. Email / password

2. Facebook authentication

3. Change of password / email

4. Forgotten passwords

Tested with SQLITE and Postgresql. To use it, create a database using the sqlx
module, and then create an auth.UserDB from that, and then call auth.New() to
create an HTTP handler for "/user/" (note the trailing slash). It provides
the following endpoints which work with GET and POST. It also allows CORS and
OPTIONS requests.

All HTTP responses might have the additional "Status" header which is a
user-readable explanation of what went wrong.

Auth

/user/auth has two cases. In case one, pass "email" and "password" and you will
receive either an HTTP error, or the UserInfo structure.

In the second case, use "method" and "token" to perform oauth authentication.
This will either sign in or create a new user. If the method is "facebook" then
the token is used to get the user's email from facebook's servers.

Create

/user/create will create a password user, using the "email" and "password".

Get

/user/get will retrieve the user's information and return it as JSON, or
return code 401 if not signed in.

Signout

/user/signout will forget the user's session cookie. It always
returns code 200

Update

/user/update takes two parameters, "email" and "password".
If email is non-blank, it changes the user's email. If password
is non-blank, it changes the password.

Oauth add

/user/oauth/add performs takes three parameters, "method",
"token" and "update_email". It performs oauth authentication
and adds the authentication to the user's account so they can
later sign in. If "update_email" is true, it also changes the
user's email address to the one provided by the oauth provider.

Oauth remove

/user/oauth/remove removes the oauth method from the user's account.
The only parameter is "method" which can be "facebook"

Forgot password

/user/forgotpassword just takes an "email" parameter and "url". If the user
exists in the system, it sends an email with the password reset
token to the user's email address. Otherwise it returns a sensible
error message in the Status header.

The url parameter must have "${TOKEN}" in it which is replaced with the
token generated in the email message.

Reset password

/user/resetpassword takes the "token" parameter and "password".
It will update the user's password and also sign them in, returning
UserInfo.
*/
package auth
