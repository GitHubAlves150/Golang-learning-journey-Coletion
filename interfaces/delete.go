package main

import "net/http"

// delete
func Delete(w http.ResponseWriter, r *http.Request, id int) {
	//id =2 que veio na URL: /rotas/2

	//PASSO 1: Procura onde está o ID=2 no slice
	index := -1 // -1 Significa não encontrado ainda
	for i, task := range banco {
		if task.ID == id { //se id foi encontrado
			index = i //i é a posição(ex:1)
			break
		}
	}

	//PASSO 2: Se não encontrou, index continua -1 e dá um erro 404
	if index == -1 {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}

	//PASSO 3: Remove o elemento da lista
	//banco[:index]   -> pega os elementos ANTES do indice, posições de 0 até index-1
	//banco[index+1:] -> pega os elementos DEPOIS do indice posições index+1 até o fim
	//apeend junta os dois pedeços, sem o elemento do meio!
	banco = append(banco[:index], banco[index+1:]...)

	//PASSO 4: Responde se deu certo
	w.WriteHeader(http.StatusNoContent) // 204 = Deu certo, não tenho nada para te enviar
	

}
