# Étape finale : Préparer l'image avec Alpine et installer KICS
FROM --platform=linux/amd64 alpine:latest
WORKDIR /app

# Installer les dépendances nécessaires pour Alpine
RUN apk update && apk add --no-cache \
    bash \
    curl \
    wget \
    musl-dev \      
    clang \         
    gcc \           
    tar \
    git \
    libstdc++ \
    make \     
    && rm -rf /var/cache/apk/*

# Installer Go 1.22.7
ENV GO_VERSION=1.22.7
RUN wget https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz && \
    rm go${GO_VERSION}.linux-amd64.tar.gz

# Ajouter Go au PATH
ENV PATH="/usr/local/go/bin:$PATH"

# Commande par défaut
CMD ["/bin/bash"]
