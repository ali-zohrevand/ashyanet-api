package Words

var KeyForEncription = "dkfl4(kdlfndlfdl"
var PermissionModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _ , _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub)  && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
`
var PermissionPolicy = `
p, admin, /*, (GET)|(POST)
p, user, /*, (GET)|(POST)
p, user, /, (GET)|(POST)
p, user, /device, POST
p, user, /aud, POST
p, user, /akd, POST
p, data_group_admin, /bob_data/*, POST
g, alice, data_group_admin
`
