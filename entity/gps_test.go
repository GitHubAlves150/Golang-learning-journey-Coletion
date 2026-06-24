// gps_test.go

package entity

import "testing"

// =====================================
// --Mock: dados de entrada fajuta
// =====================================
type Mockgps struct {
	lat float64
	lng float64
}
func (f *Mockgps)ObterCoordenadas()(float64, float64){
	return f.lat, f.lng
}

//=====================================
//--Mock: dados de entrada fajuta
//=====================================

func TestRegistraPosicao(t *testing.T) {
	mockgps := &Mockgps{
		lat: -34.4,
		lng: -45.5,
	}
	servico := NovoServico(mockgps)

	lat, lng, err:= servico.RegistrarPosicao()

	if err !=nil{
		t.Errorf("Erro inesperado %v", err)
	}

	t.Errorf("lat: %v long: %v", lat, lng)

}

//quero testar essa funcao :
//func (s *ServicoDeRastreamento)RegistrarPosicao()(float64, float64, error){
