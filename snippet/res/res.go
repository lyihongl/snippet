package res

const (
	VIEWS = "snippet/views"
)

func CheckErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

type ErrorMessage struct {
	ErrorMessage []string 
}
