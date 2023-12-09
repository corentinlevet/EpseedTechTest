# Documentation de l'application

## Architecture

L'application est composée de deux dossiers principaux:
- `src/api`: contient les fichiers relatifs à l'API REST
- `src/components`: contient les fichiers relatifs aux composants React

### Dossier `src/api`

Le dossier `src/api` contient les fichiers relatifs à l'API REST.
La class GoServer permet de communiquer avec l'API REST à travers différentes méthodes:
- `signUp`: permet de créer un compte utilisateur
- `logIn`: permet de se connecter à un compte utilisateur
- `getUsers`: permet de récupérer la liste des utilisateurs
- `createNote`: permet de créer une note
- `updateNoteForUser`: permet de mettre à jour une note pour un utilisateur
- `getNotesForUser`: permet de récupérer la liste des notes pour un utilisateur
- `deleteNoteForUser`: permet de supprimer une note pour un utilisateur

### Dossier `src/components`

Le dossier `src/components` contient les fichiers relatifs aux composants React.  
Chaque composant devrait parler de lui-même.

#### Revenir au [RAPPORT](../RAPPORT.md)
