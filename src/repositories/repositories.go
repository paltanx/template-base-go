package repositories

// agregar en container y newcontainer referencia  nuevo repo y luego retornar

type Container struct {
	OtpRepository IOtpRepository
}

func NewContainer(
	otpRepository IOtpRepository,
) *Container {
	return &Container{
		OtpRepository: otpRepository,
	}
}
