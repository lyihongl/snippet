package res

const (
	VIEWS = "snippet/views"
	LOGIN_ALERT="<script>alert('You are not logged in');</script>"
)

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type ErrorMessage struct {
	ErrorMessage []string 
}
