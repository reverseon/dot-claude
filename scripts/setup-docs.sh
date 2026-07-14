#!/bin/bash
# Fetch and set up Kubernetes EN documentation as sparse checkout and yt-dlp

DOCS_DIR="knowledgebase/k8s-docs"
YT_DLP_DIR="knowledgebase/yt-dlp"

if [ -d "$DOCS_DIR" ]; then
    echo "K8s docs already present"
else
    echo "Fetching Kubernetes EN docs (sparse checkout)..."
    git clone --depth 1 --filter=blob:none --sparse https://github.com/kubernetes/website.git "$DOCS_DIR"

    cd "$DOCS_DIR"
    git sparse-checkout set content/en/docs
    cd ../..

    rm -rf "$DOCS_DIR/.git"

    echo "Docs ready at: $DOCS_DIR/content/en/docs"
fi

if [ -d "$YT_DLP_DIR" ]; then
    echo "yt-dlp already present"
else
    echo "Fetching yt-dlp..."
    git clone https://github.com/yt-dlp/yt-dlp.git "$YT_DLP_DIR"
    rm -rf "$YT_DLP_DIR/.git"

    echo "yt-dlp ready at: $YT_DLP_DIR"
fi