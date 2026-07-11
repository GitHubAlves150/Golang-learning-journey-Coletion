CREATE TABLE IF NOT EXISTS usuarios(
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	nome varchar(30),
	email varchar(30),
	sexo varchar(6),
	criado_em TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO usuarios (nome, email, sexo) VALUES
('Lucas Alves', 'lucas@gmail.com', 'Masc'),
('Roberto Henrique', 'roberto@mail.com', 'Masc'),
('Ana Fagundes', 'aninha@ukmail.com', 'Fem'),
('Cristiane Oliveira', 'cris@hotmail.com', 'Fem'),
('Matheus Gonçalves', 'mateus@gmail.com', 'Masc');
