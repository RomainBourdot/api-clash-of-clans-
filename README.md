Groupie Tracker
Groupie Tracker est une application web réalisée en Golang qui exploite l'API officielle de Clash of Clans pour récupérer et afficher des informations sur les clans. L'objectif principal est de proposer aux utilisateurs une interface complète leur permettant de consulter, rechercher, filtrer et gérer une liste de favoris de clans.

Table des Matières
Présentation du Projet
Fonctionnalités
Installation et Lancement
Détail des Routes
Endpoints de l’API Clash of Clans
Remarques Finales
Présentation du Projet
Thème :
Groupie Tracker permet aux passionnés de Clash of Clans de consulter une collection de clans directement sur un site web.

Objectif :

Exploiter une API REST (Clash of Clans) renvoyant des données en JSON.
Implémenter manuellement des fonctionnalités essentielles telles que la recherche, les filtres, la pagination et la gestion des favoris.
Offrir une interface conviviale développée uniquement avec Golang, HTML et CSS.
Contexte :
Ce projet a été réalisé individuellement dans le cadre du cours Groupie Tracker. Il s’inscrit dans une démarche de récupération et de présentation de données d’une API publique, tout en respectant des contraintes techniques strictes (usage exclusif des librairies standards de Go, sans framework externe).

Fonctionnalités
Recherche (FT1)
L’utilisateur peut effectuer une recherche sur la collection de clans en se basant sur au moins deux propriétés (par exemple, nom et tag).

Filtres (FT2)
Possibilité de filtrer la collection selon trois critères cumulables :

Niveau minimum du clan
Nombre minimum de membres
Points minimum du clan
Pagination (FT3)
Les résultats sont affichés par lots (10, 20 ou 30 éléments par page) afin de faciliter la navigation.

Favoris (FT4)
Les utilisateurs peuvent ajouter ou retirer des clans à une liste de favoris. Cette liste est persistante grâce à l’utilisation de fichiers JSON.

Installation et Lancement
Prérequis
Go doit être installé sur votre machine.
Un éditeur de code (par exemple, VSCode) est recommandé.
Étapes d'Installation
Cloner le dépôt :
bash
Copier
git clone <URL_DU_DEPOT>
cd groupie-tracker
Vérifier l'arborescence :
Assurez-vous que le projet contient :
Les fichiers source en Go (*.go)
Les templates HTML dans le dossier templates
Les fichiers statiques (CSS, JS, images) dans le dossier assets ou static
Les fichiers de données JSON (data.json, favorites.json)
Obtention et Gestion du Token
L'API Clash of Clans requiert un token d'authentification.

Ce token est récupéré depuis le site de Clash of Clans.
Il se renouvelle automatiquement pour chaque nouvelle IP ou à l'expiration.
Vérifiez que le token dans clashofclans.services.go est valide avant de lancer l'application.
Lancement du Projet
Pour démarrer le serveur, exécutez la commande suivante à partir de la racine du projet :

bash
Copier
go run main.go
L'application sera accessible à l'adresse http://localhost:8000.

Détail des Routes Implémentées
L'application propose plusieurs routes qui gèrent l'affichage des différentes pages et fonctionnalités :

/

Méthode : GET
Description : Page d’accueil présentant l’univers du site, incluant un formulaire de recherche.
Implémentation : accueil.routes.go → controllers.AccueilController
/clans

Méthode : GET
Description : Affichage de la collection des clans avec pagination.
Implémentation : clans.routes.go → controllers.ListClans
/clans/details

Méthode : GET
Description : Affichage détaillé d’un clan, incluant statistiques, membres et log des guerres récentes.
Implémentation : clans.routes.go → controllers.DetailsClan
/research

Méthode : GET
Description : Affichage des résultats d'une recherche avec application de filtres.
Implémentation : research.routes.go → controllers.ResearchData
/favorites

Méthode : GET
Description : Affichage de la liste des favoris de l’utilisateur.
Implémentation : favoris.routes.go → controllers.FavoriteController
/favorites/list

Méthode : GET
Description : Retourne les favoris au format JSON pour une gestion côté client.
Implémentation : favoris.routes.go → controllers.ListFavoritesController
/favorites/add

Méthode : POST
Description : Ajout d’un clan à la liste des favoris.
Implémentation : favoris.routes.go → controllers.AddFavoriteController
/favorites/remove

Méthode : POST
Description : Suppression d’un clan de la liste des favoris.
Implémentation : favoris.routes.go → controllers.RemoveFavoriteController
/login

Méthode : GET
Description : Affichage de la page de connexion.
Implémentation : login.routes.go → login.LoginController
/login/traitement

Méthode : POST
Description : Traitement des identifiants de connexion d’un utilisateur.
Implémentation : login.routes.go → login.LoginTraitement
/register

Méthode : GET
Description : Affichage de la page d’inscription.
Implémentation : register.routes.go → register.RegisterController
/register/traitement

Méthode : POST
Description : Traitement de l’inscription d’un nouvel utilisateur.
Implémentation : register.routes.go → register.RegisterTraitement
/a_propos

Méthode : GET
Description : Page « À propos » contenant une FAQ sur le déroulement du projet (décomposition, gestion du temps, stratégie de documentation, etc.).
Implémentation : a_propos.routes.go → controllers.AProposController
/error

Méthode : GET
Description : Page d’erreur affichée en cas de problème avec la récupération ou le rendu des données.
Implémentation : error.routes.go → controllers.ErrorController
Endpoints de l’API Clash of Clans
Le projet exploite l'API de Clash of Clans pour récupérer des informations sur les clans via les endpoints suivants :

Recherche de Clans :

bash
Copier
GET https://api.clashofclans.com/v1/clans?name={query}&minClanLevel={minClanLevel}&minMembers={minMembers}&minClanPoints={minClanPoints}
Cet endpoint permet de récupérer une collection de clans en fonction des critères spécifiés.

Détails d’un Clan :

bash
Copier
GET https://api.clashofclans.com/v1/clans/%23{tag}
Récupère les informations détaillées d’un clan identifié par son tag (le caractère # est encodé dans l’URL).

Log des Guerres du Clan :

bash
Copier
GET https://api.clashofclans.com/v1/clans/%23{tag}/warlog?limit=5
Permet d’obtenir les dernières guerres du clan avec une limite de 5 résultats.

Remarques Finales
Gestion du Projet :
Le projet a été décomposé en plusieurs phases clés : choix de l’API, préparation des wireframes et maquettes, développement, tests et livraison. Chaque fonctionnalité a été implémentée en respectant les contraintes techniques définies (usage exclusif de Golang, HTML et CSS).

Documentation Interne :
Une FAQ détaillée sur la gestion du projet, disponible dans la page /a_propos, explique la décomposition, la répartition des tâches, la gestion du temps et la stratégie de documentation.
