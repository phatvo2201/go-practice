package shodan

const ShodanUrl = ""

type ShodanClient struct {
	Token string
}

func NewClient(token string) *ShodanClient {
	return &ShodanClient{
		Token: token,
	}
}
