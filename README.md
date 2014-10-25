# Atelier Revel

Cet atelier nécessite:

 * golang 1.2.1
 * mercurial

## Setup

Lancer le script d'installation et recharger le .bashrc:

```bash
$ bash setup.sh
$ source ~/.bashrc
```
Le téléchargement des dépendances de revel peut être long.
Ce téléchargement échouera si mercurial n'est pas présent sur le poste. Dans ce cas, les message suivant devrait apparaître:

```bash

package github.com/revel/revel
    imports code.google.com/p/go.net/websocket: exec: "hg": executable file not found in $PATH

```

Après une première exécution, le setup devrait être complet.
