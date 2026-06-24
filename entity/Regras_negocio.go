package entity

import "fmt"


func NovoServico(gps GPS)*ServicoDeRastreamento{
    return &ServicoDeRastreamento{gps:gps}
}

func (s *ServicoDeRastreamento)RegistrarPosicao()(float64, float64, error){
	lat, long:= s.gps.ObterCoordenadas()

	if lat<=0 && long <=0{
		return 0,0, fmt.Errorf("Coordenadas invalida")
	}
	return lat, long, nil
}

func ObterCoordenadas()(float64, float64){
	return -43.44, -33.22
}