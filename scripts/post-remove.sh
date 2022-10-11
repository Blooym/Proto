#!/bin/bash
for home in /home/*; do
    if [ -d "$home/.config/proto" ]; then
        rm -rf "$home/.config/proto"
    fi
done