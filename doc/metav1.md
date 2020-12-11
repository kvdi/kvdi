kVDI CRD Reference
------------------

### Packages:

-   [kvdi.io/v1](#kvdi.io%2fv1)

Types

-   [AuthResult](#AuthResult)
-   [AuthorizeRequest](#AuthorizeRequest)
-   [CreateRoleRequest](#CreateRoleRequest)
-   [CreateSessionRequest](#CreateSessionRequest)
-   [CreateUserRequest](#CreateUserRequest)
-   [DesktopSession](#DesktopSession)
-   [DesktopSessionsResponse](#DesktopSessionsResponse)
-   [FileStat](#FileStat)
-   [JWTClaims](#JWTClaims)
-   [LoginRequest](#LoginRequest)
-   [MFAResponse](#MFAResponse)
-   [Resource](#Resource)
-   [ResourceGetter](#ResourceGetter)
-   [RolesGetter](#RolesGetter)
-   [Rule](#Rule)
-   [SessionResponse](#SessionResponse)
-   [StatDesktopFileResponse](#StatDesktopFileResponse)
-   [TemplatesGetter](#TemplatesGetter)
-   [UpdateMFARequest](#UpdateMFARequest)
-   [UpdateRoleRequest](#UpdateRoleRequest)
-   [UpdateUserRequest](#UpdateUserRequest)
-   [UsersGetter](#UsersGetter)
-   [VDIUser](#VDIUser)
-   [VDIUserRole](#VDIUserRole)
-   [Verb](#Verb)

kvdi.io/v1
----------

Package v1 contains API Schema definitions for the meta v1 API group

Resource Types:

### AuthResult

AuthResult represents a response from an authentication attempt to a
provider. It contains user information, roles, and any other auth
requirements.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>User</code> <em><a href="#VDIUser">VDIUser</a></em></td>
<td><p>The authenticated user and their roles</p></td>
</tr>
<tr class="even">
<td><code>RedirectURL</code> <em>string</em></td>
<td><p>The provider can populate this field to signify a redirect is required, e.g. for OIDC.</p></td>
</tr>
<tr class="odd">
<td><code>RefreshNotSupported</code> <em>bool</em></td>
<td><p>In the case of OIDC, the refresh tokens cannot be used. Because when the user tries to use them, there is no way to query the provider for the user’s information without initializing a new auth flow. For now, the provider can set this to false to signal to the server that a refresh is not possible.</p></td>
</tr>
</tbody>
</table>

### AuthorizeRequest

AuthorizeRequest is a request with an OTP for receiving an authorized
token.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>otp</code> <em>string</em></td>
<td><p>The one-time password</p></td>
</tr>
<tr class="even">
<td><code>state</code> <em>string</em></td>
<td><p>The state secret for the request flow</p></td>
</tr>
</tbody>
</table>

### CreateRoleRequest

CreateRoleRequest represents a request for a new role.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>name</code> <em>string</em></td>
<td><p>The name of the new role</p></td>
</tr>
<tr class="even">
<td><code>annotations</code> <em>map[string]string</em></td>
<td><p>Annotations to apply to the role</p></td>
</tr>
<tr class="odd">
<td><code>rules</code> <em><a href="#Rule">[]Rule</a></em></td>
<td><p>Rules to apply to the new role.</p></td>
</tr>
</tbody>
</table>

### CreateSessionRequest

CreateSessionRequest requests a new desktop session with the givin
parameters.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>template</code> <em>string</em></td>
<td><p>The template to create the session from.</p></td>
</tr>
<tr class="even">
<td><code>namespace</code> <em>string</em></td>
<td><p>The namespace to launch the template in. Defaults to default.</p></td>
</tr>
</tbody>
</table>

### CreateUserRequest

CreateUserRequest represents a request to create a new user. Not all
auth providers will be able to implement this route and can instead
return an error describing why.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>username</code> <em>string</em></td>
<td><p>The user name for the new user.</p></td>
</tr>
<tr class="even">
<td><code>password</code> <em>string</em></td>
<td><p>The password for the new user.</p></td>
</tr>
<tr class="odd">
<td><code>roles</code> <em>[]string</em></td>
<td><p>Roles to assign the new user. These are the names of VDIRoles in the cluster.</p></td>
</tr>
</tbody>
</table>

### DesktopSession

DesktopSession describes the properties and status of a desktop session.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>name</code> <em>string</em></td>
<td><p>The name of the desktop session.</p></td>
</tr>
<tr class="even">
<td><code>namespace</code> <em>string</em></td>
<td><p>The namespace of the desktop session.</p></td>
</tr>
<tr class="odd">
<td><code>user</code> <em>string</em></td>
<td><p>The username of the user who owns this session.</p></td>
</tr>
<tr class="even">
<td><code>socketType</code> <em>string</em></td>
<td><p>The type of display socket the desktop is using</p></td>
</tr>
<tr class="odd">
<td><code>status</code> <em><a href="#DesktopSessionStatus">DesktopSessionStatus</a></em></td>
<td><p>Connection status for the session.</p></td>
</tr>
</tbody>
</table>

### DesktopSessionsResponse

DesktopSessionsResponse contains a list of desktop sessions and
information about their statuses.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>sessions</code> <em><a href="#DesktopSession">[]*github.com/tinyzimmer/kvdi/pkg/apis/meta/v1.DesktopSession</a></em></td>
<td><p>A list of desktop sessions.</p></td>
</tr>
</tbody>
</table>

### FileStat

(*Appears on:*
[StatDesktopFileResponse](#StatDesktopFileResponse))

FileStat contains information about a queried file. Contents will only
contain nested FileStat objects when this object represents the root of
the query.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>name</code> <em>string</em></td>
<td><p>The name of the file or directory</p></td>
</tr>
<tr class="even">
<td><code>isDirectory</code> <em>bool</em></td>
<td><p>True if the file is a directory</p></td>
</tr>
<tr class="odd">
<td><code>size</code> <em>int64</em></td>
<td><p>The size of the file when IsDirectory is false</p></td>
</tr>
<tr class="even">
<td><code>contents</code> <em><a href="#FileStat">[]*github.com/tinyzimmer/kvdi/pkg/apis/meta/v1.FileStat</a></em></td>
<td><p>When IsDirectory is true, the contents of the directory</p></td>
</tr>
</tbody>
</table>

### JWTClaims

JWTClaims represents the claims used when issuing JWT tokens.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>user</code> <em><a href="#VDIUser">VDIUser</a></em></td>
<td><p>The user with their permissions when the token was generated</p></td>
</tr>
<tr class="even">
<td><code>authorized</code> <em>bool</em></td>
<td><p>Whether the user is fully authorized</p></td>
</tr>
<tr class="odd">
<td><code>renewable</code> <em>bool</em></td>
<td><p>Whether a refresh token was issued with the claims</p></td>
</tr>
<tr class="even">
<td><code>StandardClaims</code> <em>github.com/dgrijalva/jwt-go.StandardClaims</em></td>
<td><p>The standard JWT claims</p></td>
</tr>
</tbody>
</table>

### LoginRequest

LoginRequest represents a request for a session token. Different auth
providers may not always need this request, and can instead redirect
/api/login as needed. All the auth provider needs to do in the end is
return a JWT token that contains a fulfilled VDIUser.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>username</code> <em>string</em></td>
<td><p>Username</p></td>
</tr>
<tr class="even">
<td><code>password</code> <em>string</em></td>
<td><p>Password</p></td>
</tr>
<tr class="odd">
<td><code>state</code> <em>string</em></td>
<td><p>State generated by requesting client to prevent CSRF and retrieve tokens from an oidc flow</p></td>
</tr>
<tr class="even">
<td><code>request</code> <em>net/http.Request</em></td>
<td><p>the underlying request object for usage by auth providers</p></td>
</tr>
</tbody>
</table>

### MFAResponse

MFAResponse contains the response to an UpdateMFARequest or
GetMFARequest.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>enabled</code> <em>bool</em></td>
<td><p>Whether MFA is enabled for the user</p></td>
</tr>
<tr class="even">
<td><code>provisioningURI</code> <em>string</em></td>
<td><p>If enabled is set, a provisioning URI is also returned.</p></td>
</tr>
<tr class="odd">
<td><code>verified</code> <em>bool</em></td>
<td><p>If enabled is set, whether or not the user has verified their MFA setup</p></td>
</tr>
</tbody>
</table>

Resource (`string` alias)

(*Appears on:* [Rule](#Rule))

Resource represents the target of an API action

### ResourceGetter

ResourceGetter is an interface for retrieving lists of kVDI related
resources. Its primary purpose is to pass an interface to rbac
evaluations so they can check permissions against present resources.

### RolesGetter

RolesGetter is an interface that can be used to retrieve available roles
while checking user permissions.

### Rule

(*Appears on:* [CreateRoleRequest](#CreateRoleRequest),
[UpdateRoleRequest](#UpdateRoleRequest),
[VDIUserRole](#VDIUserRole))

Rule represents a set of permissions applied to a VDIRole. It mostly
resembles an rbacv1.PolicyRule, with resources being a regex and the
addition of a namespace selector.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>verbs</code> <em><a href="#Verb">[]Verb</a></em></td>
<td><p>The actions this rule applies for. VerbAll matches all actions.</p></td>
</tr>
<tr class="even">
<td><code>resources</code> <em><a href="#Resource">[]Resource</a></em></td>
<td><p>Resources this rule applies to. ResourceAll matches all resources.</p></td>
</tr>
<tr class="odd">
<td><code>resourcePatterns</code> <em>[]string</em></td>
<td><p>Resource regexes that match this rule. This can be template patterns, role names or user names. There is no All representation because * will have that effect on its own when the regex is evaluated.</p></td>
</tr>
<tr class="even">
<td><code>namespaces</code> <em>[]string</em></td>
<td><p>Namespaces this rule applies to. Only evaluated for template launching permissions. NamespaceAll matches all namespaces.</p></td>
</tr>
</tbody>
</table>

### SessionResponse

SessionResponse represents a response with a new session token

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>token</code> <em>string</em></td>
<td><p>The X-Session-Token to use for future requests.</p></td>
</tr>
<tr class="even">
<td><code>expiresAt</code> <em>int64</em></td>
<td><p>The time the token expires.</p></td>
</tr>
<tr class="odd">
<td><code>renewable</code> <em>bool</em></td>
<td><p>Whether an HttpOnly was sent back with the request enabling token refresh.</p></td>
</tr>
<tr class="even">
<td><code>user</code> <em><a href="#VDIUser">VDIUser</a></em></td>
<td><p>Information about the authenticated user and their permissions.</p></td>
</tr>
<tr class="odd">
<td><code>authorized</code> <em>bool</em></td>
<td><p>Whether the user is fully authorized (e.g. false if MFA is required but not provided yet)</p></td>
</tr>
<tr class="even">
<td><code>state</code> <em>string</em></td>
<td><p>The state secret generated by the client</p></td>
</tr>
</tbody>
</table>

### StatDesktopFileResponse

StatDesktopFileResponse contains the info for a queried file inside a
desktop dession.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>stat</code> <em><a href="#FileStat">FileStat</a></em></td>
<td></td>
</tr>
</tbody>
</table>

### TemplatesGetter

TemplatesGetter is an interface that can be used to retrieve available
templates while checking user permissions.

### UpdateMFARequest

UpdateMFARequest sets the MFA configuration for the user. If enabling, a
provisioning URI will be returned.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>enabled</code> <em>bool</em></td>
<td><p>When set, will enable MFA for the given user. If false, will disable MFA.</p></td>
</tr>
</tbody>
</table>

### UpdateRoleRequest

UpdateRoleRequest requests updates to an existing role. The existing
attributes will be entirely replaced with those supplied in the payload.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>annotations</code> <em>map[string]string</em></td>
<td><p>The new annotations for the role</p></td>
</tr>
<tr class="even">
<td><code>rules</code> <em><a href="#Rule">[]Rule</a></em></td>
<td><p>The new rules for the role.</p></td>
</tr>
</tbody>
</table>

### UpdateUserRequest

UpdateUserRequest requests updates to an existing user. Not all auth
providers will be able to implement this route and can instead return an
error describing why.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>password</code> <em>string</em></td>
<td><p>When populated, will change the password for the user.</p></td>
</tr>
<tr class="even">
<td><code>roles</code> <em>[]string</em></td>
<td><p>When populated will change the roles for the user.</p></td>
</tr>
</tbody>
</table>

### UsersGetter

UsersGetter is an interface that can be used to retrieve available users
while checking user permissions.

### VDIUser

(*Appears on:* [AuthResult](#AuthResult),
[JWTClaims](#JWTClaims),
[SessionResponse](#SessionResponse))

VDIUser represents a user in kVDI. It is the auth providers
responsibility to take an authentication request and generate a JWT with
claims defining this object.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>name</code> <em>string</em></td>
<td><p>A unique name for the user</p></td>
</tr>
<tr class="even">
<td><code>roles</code> <em><a href="#VDIUserRole">[]*github.com/tinyzimmer/kvdi/pkg/apis/meta/v1.VDIUserRole</a></em></td>
<td><p>A list of roles applide to the user. The grants associated with each user are embedded in the JWT signed when authenticating.</p></td>
</tr>
<tr class="odd">
<td><code>mfa</code> <em><a href="#UserMFAStatus">UserMFAStatus</a></em></td>
<td><p>MFA status for the user</p></td>
</tr>
<tr class="even">
<td><code>sessions</code> <em><a href="#DesktopSession">[]*github.com/tinyzimmer/kvdi/pkg/apis/meta/v1.DesktopSession</a></em></td>
<td><p>Any active sessions for the user - new field that is only populated on a /api/whoami request.</p></td>
</tr>
</tbody>
</table>

### VDIUserRole

VDIUserRole represents a VDIRole, but only with the data that is to be
embedded in the JWT. Primarily, leaving out useless metadata that will
inflate the token.

<table>
<thead>
<tr class="header">
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td><code>name</code> <em>string</em></td>
<td><p>The name of the role, this must match the VDIRole from which this object derives.</p></td>
</tr>
<tr class="even">
<td><code>rules</code> <em><a href="#Rule">[]Rule</a></em></td>
<td><p>The rules for this role.</p></td>
</tr>
</tbody>
</table>

Verb (`string` alias)

(*Appears on:* [Rule](#Rule))

Verb represents an API action

------------------------------------------------------------------------

*Generated with `gen-crd-api-reference-docs` on git commit `eade2e6`.*
