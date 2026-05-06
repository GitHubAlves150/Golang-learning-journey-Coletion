# Interface

```go
INTERFACE "Campeao"
     │
     ├── Atacar()  →  todo campeão precisa saber atacar
     └── Defender() → todo campeão precisa saber defender

GUERREIRO ─────────┐
MAGO ──────────────┤ TODOS são Campeões!
ARQUEIRO ──────────┤
PALADINO ──────────┘

FUNÇÃO Batalha(campeao Campeao)
    campeao.Atacar()    ← Funciona com QUALQUER campeão! 🎉
    campeao.Defender()  ← Não importa se é Guerreiro ou Mago!

```