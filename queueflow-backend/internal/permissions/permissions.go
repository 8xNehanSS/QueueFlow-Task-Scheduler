package permissions

const (
	USER_CREATE = "user:create"

	USER_READ = "user:read"

	JOB_CREATE = "job:create"

	JOB_DELETE = "job:delete"

	JOB_VIEW = "job:view"
)

var Roles = map[string][]string{

	"admin": {
		"job:create",
		"job:view",
		"job:delete",
		"user:create",
	},

	"user": {
		"job:create",
		"job:view",
	},

	"worker": {
		"job:view",
	},
}

func HasPermission(
	role string,
	permission string,
) bool {

	for _, p := range Roles[role] {

		if p == permission {
			return true
		}
	}

	return false
}
