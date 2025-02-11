watch -n 1 \
    "ls -l --all --recursive --full-time | grep -vE node_modules | sha256sum" \
    && ./run.sh
