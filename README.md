# 🎯 **Groupie Tracker**

**Groupie Tracker** est une application web développée en **Go**, exploitant l'API Clash of Clans pour récupérer et afficher des informations sur les clans.

## 🚀 **Objectif du Projet**

L’application permet aux utilisateurs de :
- **Consulter** une collection de clans.
- **Rechercher et filtrer** ces données.
- **Naviguer** grâce à un système de **pagination**.
- **Gérer** une liste de **favoris persistante**.

📌 **Contraintes techniques respectées :**
- Développement en **Golang, HTML et CSS uniquement** (sans frameworks externes).
- Exploitation d’une **API REST** (Clash of Clans) sans quotas d’utilisation.
- Implémentation **manuelle** des fonctionnalités clés (recherche, filtres, pagination, favoris via fichiers JSON).

## 🛠️ **Fonctionnalités Implémentées**

### 🔍 **Recherche (FT1)**
Recherche avancée de clans en combinant **au moins deux critères** (ex: nom et tag).

### 🎚️ **Filtrage (FT2)**
Filtrage des clans selon **trois critères cumulables** :
- **Niveau minimum du clan**
- **Nombre minimum de membres**
- **Nombre minimum de points de clan**

### 📄 **Pagination (FT3)**
Affichage des résultats **par lots de 10, 20 ou 30 éléments** pour une navigation fluide.

### ⭐ **Gestion des Favoris (FT4)**
- **Ajout/suppression** de clans favoris.
- **Persistance des favoris** grâce aux fichiers JSON.

## 📦 **Installation et Lancement**

### ✅ **Prérequis**
- **Go** installé sur votre machine.
- Un éditeur de code (ex: **VSCode**).

### 📥 **Installation**
1. **Cloner le dépôt GitHub :**
   ```bash
   git clone <URL_DU_DEPOT>
   cd groupie-tracker
   ```
2. **Structure attendue du projet :**
   - 📂 **Code source en Go** (`*.go`)
   - 📂 **Templates HTML** (`templates/`)
   - 📂 **Fichiers statiques** (`assets/` ou `static/`)
   - 📂 **Fichiers JSON** (`data.json`, `favorites.json`)

### 🔑 **Obtention du Token API**
- L’API Clash of Clans nécessite un **token d’authentification**.
- Ce token doit être généré sur [Clash of Clans API](https://developer.clashofclans.com/).
- Il doit être renseigné dans `clashofclans.services.go`.

### 🚀 **Lancement du Serveur**
```bash
go run main.go
```
- Le serveur est accessible sur **http://localhost:8000**.
- Ouvrez un navigateur pour interagir avec l’application.

## 🔗 **Routes Implémentées**

| **Route**                 | **Méthode** | **Description** |
|---------------------------|------------|----------------|
| `/`                       | **GET**        | Page d’accueil |
| `/clans`                  | **GET**        | Liste des clans avec pagination |
| `/clans/details`          | **GET**        | Détails d’un clan |
| `/research`               | **GET**        | Recherche et filtres |
| `/favorites`              | **GET**        | Liste des favoris |
| `/favorites/list`         | **GET**        | Favoris (JSON) |
| `/favorites/add`          | **POST**       | Ajouter un favori |
| `/favorites/remove`       | **POST**       | Supprimer un favori |
| `/login`                  | **GET**        | Page de connexion |
| `/login/traitement`       | **POST**       | Validation connexion |
| `/register`               | **GET**        | Page d’inscription |
| `/register/traitement`    | **POST**       | Validation inscription |
| `/a_propos`               | **GET**        | Page d’informations |
| `/error`                  | **GET**        | Gestion des erreurs |

## 📡 **API Clash of Clans Utilisée**

### 🔎 **Recherche de Clans**
```http
GET https://api.clashofclans.com/v1/clans?name={query}&minClanLevel={minClanLevel}&minMembers={minMembers}&minClanPoints={minClanPoints}
```

### 🏆 **Détails d’un Clan**
```http
GET https://api.clashofclans.com/v1/clans/%23{tag}
```

### ⚔️ **Log des Guerres du Clan**
```http
GET https://api.clashofclans.com/v1/clans/%23{tag}/warlog?limit=5
```

## 📖 **À Propos du Projet**
Ce projet a été réalisé dans le cadre du cours **Groupie Tracker**. Toutes les fonctionnalités demandées (recherche, filtres, pagination, favoris) ont été développées en **Go, HTML et CSS**, sans frameworks externes.


🎯 _Merci d’avoir lu ce README ! 🚀_

