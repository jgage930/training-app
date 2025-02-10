watch --chgexit -n 1 \
    "ls -l --all --recursive --full-time | grep -vE node_modules | sha256sum" \
    && echo "Detected the modification of a file or directory"