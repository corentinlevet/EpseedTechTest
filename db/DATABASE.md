# Documentation de la base de données

## Schéma de la base de données

La base de données est composée de 2 tables:
- `users`
- `notes`

### Table `users`

La table `users` est composée de 7 colonnes:
- `id`: identifiant unique de l'utilisateur
- `email`: adresse email de l'utilisateur
- `username`: nom d'utilisateur
- `password`: mot de passe de l'utilisateur
- `salt`: sel utilisé pour le hashage du mot de passe
- `created_at`: date de création de l'utilisateur
- `updated_at`: date de dernière modification de l'utilisateur

### Table `notes`

La table `notes` est composée de 6 colonnes:
- `id`: identifiant unique de la note
- `user_id`: identifiant de l'utilisateur à qui appartient la note
- `title`: titre de la note
- `content`: contenu de la note
- `created_at`: date de création de la note
- `updated_at`: date de dernière modification de la note

## Scripts d'import/export

Le script d'import permet d'importer les données du fichier `db/dump.sql` dans la base de données.
Pour lancer le script d'import, veuillez vous placer à la racine du projet et lancer la commande suivante:
```bash
./db/scripts/import.sh
```

Le script d'export permet d'exporter les données de la base de données dans le fichier `db/dump.sql`.
Pour lancer le script d'export, veuillez vous placer à la racine du projet et lancer la commande suivante:
```bash
./db/scripts/dump.sh
```

Il se peut que vous n'ayez pas les droits nécessaires pour lancer ces scripts.
Dans ce cas, veuillez lancer les commandes suivantes:
```bash
chmod +x ./db/scripts/import.sh
chmod +x ./db/scripts/dump.sh
```

#### Revenir au [RAPPORT](../RAPPORT.md)
