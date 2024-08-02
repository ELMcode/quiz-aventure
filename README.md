# Quiz Aventure

Quiz Aventure est un mini-jeu web interactif où le joueur traverse une série de niveaux en répondant correctement à des questions de quiz. Chaque bonne réponse permet de passer au niveau suivant, tandis qu'une mauvaise réponse entraîne une pénalité de points.

## Fonctionnalités

- Interface utilisateur simple
- Trois niveaux de difficulté croissants avec cinq questions chacun
- Système de score pour suivre les performances du joueur
- Animation selon le résultat de la réponse
- Sauvegarde du pseudo et du score dans le localStorage
- Option de réinitialisation du quiz et de changement de joueur

## Technologies utilisées

- Front-end : Vue.js 3.2.13
- Back-end : Go 1.22.5

## Installation

### Prérequis

- npm
- Node.js
- Go

### Étapes

1. Clonez le dépôt :

```bash
git clone https://github.com/ELMcode/quiz-aventure.git
cd quiz-aventure
```

2. Installez les dépendances pour le front-end :

```bash
cd front
npm install
```

3. Installez les dépendances pour le back-end :

```bash
cd ../back
go mod tidy
```

4. Créez un fichier `config.yaml` à la racine du répertoire `back` en vous basant sur `config.example.yaml` :

```yaml
server:
  address: ":8081" # Port que vous souhaitez utiliser pour le back-end

cors:
  allowed_origins:
    - "http://localhost:8080" # Origine autorisée pour le front-end
```

## Utilisation

### Démarrer le serveur back-end

Assurez-vous d'être dans le répertoire `back` et lancez le serveur Go :

```bash
go run cmd/quiz-aventure/main.go
```

Le serveur back-end sera démarré sur l'adresse spécifiée dans `config.yaml`

### Démarrer le serveur front-end

Assurez-vous d'être dans le répertoire `front` et lancez l'application Vue.js :

```bash
npm run serve
```

L'application front-end sera démarrée sur `http://localhost:8080`.

### Jouer au Quiz

1. Ouvrez votre navigateur et allez à `http://localhost:8080`.
2. Entrez votre pseudo et commencez le quiz.
3. Répondez aux questions et progressez à travers les niveaux.
