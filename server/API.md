# Documentation de l'API REST

## Architecture

L'API REST est composée de 3 dossiers principaux:
- `bin`: contient les fichiers binaires de l'API REST
- `cmd`: contient les fichiers relatifs au lancement de l'API REST
- `internal`: contient les fichiers relatifs à l'API REST

### Dossier `bin`

Le dossier `bin` contient l'exécutable de l'API REST.

### Dossier `cmd`

Le dossier `cmd` contient les fichiers relatifs au lancement de l'API REST.
Le fichier `cmd/server/main.go` est le point d'entrée de cette application.

### Dossier `internal`

Le dossier `internal` contient différents dossiers:
- `config`: contient les fichiers relatifs à la configuration de l'API REST, notamment le fait de lire les variables d'environnement ainsi que de définir la zone horaire
- `db`: contient les fichiers relatifs à la base de données
- `handler`: contient les fichiers relatifs aux handlers de l'API REST
- `server`: contient les fichiers relatifs au serveur de l'API REST, notamment le fait de lancer le serveur et de définir les routes

#### Revenir au [RAPPORT](../RAPPORT.md)
