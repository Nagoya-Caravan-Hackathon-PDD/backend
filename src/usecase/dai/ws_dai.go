package dai

type WSDai interface {
	Start(times int, token string) error
	Turn(times int, token string) error
	Result(times int, token string) error
	End(times int, token string) error
}
