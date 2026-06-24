package entity


type GPS interface{
	ObterCoordenadas()(float64, float64)
}