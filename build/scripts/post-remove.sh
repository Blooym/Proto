#!/bin/bash
for home in /home/*; do
    if [ -d "$home/.config/proto" ]; then
        rm -rf "$home/.config/proto"
    fi

    if [ -d "$home/.cache/proto" ]; then
        rm -rf "$home/.cache/proto"
    fi
done