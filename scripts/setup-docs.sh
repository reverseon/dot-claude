#!/bin/bash
# Fetch and set up Kubernetes EN documentation as sparse checkout

DOCS_DIR="knowledgebase/k8s-docs"

if [ -d "$DOCS_DIR" ]; then
    echo "K8s docs already present"
    exit 0
fi

echo "Fetching Kubernetes EN docs (sparse checkout)..."
git clone --depth 1 --filter=blob:none --sparse https://github.com/kubernetes/website.git "$DOCS_DIR"

cd "$DOCS_DIR"
git sparse-checkout set content/en/docs
cd ../..

echo "Docs ready at: $DOCS_DIR/content/en/docs"