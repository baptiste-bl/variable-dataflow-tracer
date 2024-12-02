# Variable Dataflow Tracer

**Variable Dataflow Tracer** est un outil open-source développé par **CyberDefence**, conçu pour analyser les flux de données des variables spécifiques à travers plusieurs langages de programmation. Il permet de retracer l'origine et l'utilisation d'une variable donnée au sein d'une base de code pour générer des graphiques de flux de données détaillés. Cet outil aide les développeurs et les ingénieurs en sécurité à mieux comprendre comment les données circulent dans leurs applications, en se concentrant sur des variables individuelles. Le projet est distribué sous licence MIT et est ouvert aux contributions de la communauté.

---

## Badge

![Licence](https://img.shields.io/badge/license-apache--2.0-blue)
![Version Go](https://img.shields.io/badge/go-%3E%3D1.16-blue)
![Statut Build](https://img.shields.io/badge/build-passing-brightgreen)

---

## Table des matières

1. [Introduction](#introduction)
2. [Fonctionnalités clés](#fonctionnalités-clés)
3. [Langages supportés](#langages-supportés)
4. [Prérequis](#prérequis)
5. [Installation](#installation)
6. [Utilisation](#utilisation)
   - [En tant qu'outil en ligne de commande](#en-tant-quoutil-en-ligne-de-commande)
   - [En tant que bibliothèque](#en-tant-que-bibliothèque)
   - [Arguments](#arguments)
   - [Exemples](#exemples)
7. [Tests](#tests)
8. [Structure du code](#structure-du-code)
9. [Limitations](#limitations)
10. [Recommandations d'utilisation](#recommandations-dutilisation)
11. [Contribuer](#contribuer)
12. [Licence](#licence)
13. [Contact](#contact)

---

## Introduction

**Variable Dataflow Tracer** est un outil polyvalent conçu pour analyser les flux de données des variables spécifiques dans les programmes écrits dans divers langages de programmation. Grâce à [Tree-sitter](https://tree-sitter.github.io/tree-sitter/) pour l'analyse syntaxique, il offre des informations détaillées sur le parcours d'une variable donnée à travers une base de code, retraçant son origine et suivant son utilisation dans l'application. Cet outil est particulièrement utile pour :

- Déboguer des applications complexes.
- Effectuer des évaluations de vulnérabilité.
- Réaliser des revues de code approfondies.

---

## Fonctionnalités clés

- **Analyse spécifique aux variables** : Retrace l'origine et l'utilisation d'une variable donnée.
- **Support multilingue** : Compatible avec plusieurs langages de programmation.
- **Analyse récursive des flux de données** : Parcourt les fonctions et appels de fonctions pour tracer les sources de données.
- **Détection automatique des variables** : Identifie automatiquement les variables à une ligne donnée.
- **Suivi des variables globales** : Suit les variables globales et leurs valeurs dans tout le code.
- **Double usage** : Peut être utilisé comme outil en ligne de commande ou intégré comme bibliothèque dans des projets Go.
- **Journaux détaillés** : Fournit des options de journalisation pour diagnostiquer les problèmes ou comprendre le processus d'analyse.

---

## Langages supportés

Les langages actuellement pris en charge sont :

- Go
- Python
- Java
- JavaScript
- C
- C++
- C#
- PHP
- Ruby
- Rust

> **Remarque** : Le projet est en cours de développement. De nouveaux langages et fonctionnalités seront ajoutés prochainement.

---

## Prérequis

- [Go](https://golang.org/doc/install) version **1.16** ou ultérieure.
- Docker et Visual Studio Code (avec l’extension **Dev Containers** activée).

---

## Installation

### Étapes d'installation

1. **Cloner le dépôt** :

   ```bash
   git clone https://github.com/votre-nom-utilisateur/variable-dataflow-tracer.git
   cd variable-dataflow-tracer
   ```

2. **Ouvrir le projet dans un conteneur de développement** :

   - Ouvrez Visual Studio Code.
   - Dans la palette de commandes (`Ctrl + Shift + P`), sélectionnez **Dev Containers: Rebuild and Reopen in Container**.
   - Attendez que l'environnement de conteneur soit configuré.

3. **Installer les dépendances** :

   ```bash
   go mod tidy
   ```

   Cela garantit que toutes les dépendances nécessaires sont installées.

---

## Utilisation

### En tant qu'outil en ligne de commande

Exemple de commande pour analyser une variable spécifique :

```bash
go run main.go -f <chemin_du_fichier> -l <numéro_de_ligne> -lang <langage> -var <nom_de_la_variable> [--verbose] [--debug]
```

**Arguments principaux** :

- `-f` : Chemin vers le fichier à analyser.
- `-l` : Ligne de départ pour l'analyse.
- `-lang` : Langage de programmation (ex. `python`, `go`).
- `-var` : Nom de la variable à analyser.
- `--verbose` : Active les journaux détaillés.
- `--debug` : Active les journaux de débogage.

### En tant que bibliothèque

Exemple d'utilisation dans un projet Go :

```go
package main

import (
    "fmt"
    "github.com/votre-nom-utilisateur/variable-dataflow-tracer/core"
    "github.com/votre-nom-utilisateur/variable-dataflow-tracer/models"
)

func main() {
    config := models.Config{
        FilePath:  "chemin/vers/le/fichier",
        StartLine: 20,
        Language:  "go",
        Verbose:   true,
        Debug:     false,
        Variable:  "maVariable",
    }

    result, err := core.RunDataflowAnalysis(config)
    if err != nil {
        fmt.Println("Erreur :", err)
        return
    }
    fmt.Println("Résultat :", result)
}
```

---

## Tests

Pour exécuter les tests :

1. Accédez au répertoire des tests :

   ```bash
   cd tests
   ```

2. Lancez le script de test :

   ```bash
   go run test_all_languages.go
   ```

---

## Limitations

- **Langages partiellement supportés** : Certains langages ne sont pas encore pleinement compatibles.
- **Complexité des codes** : Les bases de code très complexes peuvent poser des défis.
- **Tests limités** : Le projet nécessite davantage de tests pour garantir la fiabilité.

---

## Contribuer

1. **Forkez le dépôt**.
2. **Créez une branche** :
   ```bash
   git checkout -b feature/nom-de-votre-feature
   ```
3. **Soumettez une pull request**.

---

## Licence

Ce projet est sous licence **MIT**.

---

## Contact

Pour toute question :

- **Entreprise** : CyberDefence
- **Email** : [contact@cyberdefence.com](mailto:contact@cyberdefence.com)  
- **Site web** : [www.cyberdefence.com](https://www.cyberdefence.com)
