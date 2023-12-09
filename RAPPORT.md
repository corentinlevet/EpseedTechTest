# Rapport détaillé du projet

## Sommaire

- [Rapport détaillé du projet](#rapport-détaillé-du-projet)
  - [Sommaire](#sommaire)
  - [Présentation du projet](#présentation-du-projet)
  - [Prérequis](#prérequis)
  - [Lancement du projet](#lancement-du-projet)
    - [Point d'attention](#point-dattention)
  - [Choix techniques](#choix-techniques)
    - [API REST](#api-rest)
    - [Application web](#application-web)
    - [Base de données](#base-de-données)
  - [Améliorations possibles](#améliorations-possibles)
    - [API REST](#api-rest-1)
    - [Application web](#application-web-1)
  - [Conclusion](#conclusion)

## Présentation du projet

Ce projet est constitué d'une [API REST](https://www.redhat.com/fr/topics/api/what-is-a-rest-api), d'une [application web](https://fr.wikipedia.org/wiki/Application_web) et d'une [base de données](https://fr.wikipedia.org/wiki/Base_de_donn%C3%A9es).

Vous retrouverez dans ce rapport les détails de la mise en place de ce projet, ainsi que les choix techniques effectués.

Pour plus de détails sur le fonctionnement de chaque partie, veuillez vous référer aux fichiers README.md de chaque dossier:
- [API REST](./server/API.md)
- [Application web](./app/APPLICATION.md)
- [Base de données](./db/DATABASE.md)

## Prérequis

Pour lancer ce projet, vous aurez besoin de:
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Lancement du projet

Pour lancer le projet, il vous suffit de lancer la commande suivante à la racine du projet:
```bash
docker compose up
```

### Point d'attention

Lors du premier lancement du projet, il est possible que l'API REST ne se lance pas correctement du au fait que le temps de démarrage de la base de données est plus long que celui de l'API REST.

Si cela se produit, veuillez attendre quelques secondes afin que la base de données soit correctement lancée, puis relancer la commande:
```bash
docker compose up server
```

## Choix techniques

### API REST

L'API REST est développée en [Go](https://golang.org/). Elle utilise un [ORM](https://fr.wikipedia.org/wiki/Mapping_objet-relationnel) ([GORM](https://gorm.io/)) pour communiquer avec la base de données.
Cette technologie a été choisie pour trois principales raisons:
- La découverte du langage Go afin de correspondre au mieux aux attentes de l'entreprise Epseed
- La rapidité de développement de l'API REST grâce à l'utilisation d'un ORM
- La facilité de déploiement de l'API REST grâce à la compilation du code en un seul fichier binaire

Retrouvez plus de détails sur l'API REST dans le fichier [API.md](./server/API.md).

### Application web

L'application web est développée en [React](https://fr.reactjs.org/).
Cette technologie a été choisie pour deux principales raisons:
- Le framework React est une technologie que je connais bien et qui m'a permis de rapidement développer l'application web
- La facilité de déploiement d'une application [Node.js](https://nodejs.org/en/)

Retrouvez plus de détails sur l'application web dans le fichier [APPLICATION.md](./app/APPLICATION.md).

### Base de données

La base de données est une base de données [MariaDB](https://mariadb.org/).
Cette technologie a été choisie pour deux principales raisons:
- Comme pour l'API REST, correspondre au mieux aux attentes de l'entreprise Epseed
- Le langage SQL, simple et pratique pour ce type de projet

Retrouvez plus de détails sur la base de données dans le fichier [DATABASE.md](./db/DATABASE.md).

## Améliorations possibles

### API REST

Comme je ne connais pas encore très bien le langage Go, il est possible que l'architecture de l'API REST ne soit pas optimale.
Il est donc éventuellement possible de revoir l'architecture de l'API REST afin de la rendre plus performante.

### Application web

L'application web est actuellement très basique. Il serait possible de l'améliorer en ajoutant des fonctionnalités comme:
- La possibilité de modifier/supprimer un utilisateur
- La possibilité de rajouter du contenu multimédia dans une note (image, vidéo, etc.)

## Conclusion

Ce projet a été très intéressant à réaliser. Il m'a permis de découvrir de nouvelles technologies, et de me perfectionner sur d'autres.
