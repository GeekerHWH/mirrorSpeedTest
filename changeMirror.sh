#!/bin/bash

# Check if the script is run as root
if [ "$(id -u)" != "0" ]; then
   echo "This script must be run as root" 1>&2
   exit 1
fi

# User prompt for mirror selection
echo "Choose a mirror source:"
echo "1. Tsinghua University mirrors"
echo "2. Aliyun mirrors"
read -p "Enter your choice (1 or 2): " choice

# Define variables
sources_list="/etc/apt/sources.list"
backup_file="$sources_list.bak"

# Backup /etc/apt/sources.list to the same directory as sources.list.bak
sudo cp "$sources_list" "$backup_file"

# Function to set the mirror source based on user choice
set_mirror_source() {
    local mirror_url=$1

    echo "# By default, source mirrors are commented out to improve apt update speed. Uncomment if needed.
    deb $mirror_url bookworm main contrib non-free non-free-firmware
    # deb-src $mirror_url bookworm main contrib non-free non-free-firmware

    deb $mirror_url bookworm-updates main contrib non-free non-free-firmware
    # deb-src $mirror_url bookworm-updates main contrib non-free non-free-firmware

    deb $mirror_url bookworm-backports main contrib non-free non-free-firmware
    # deb-src $mirror_url bookworm-backports main contrib non-free non-free-firmware

    deb https://security.debian.org/debian-security bookworm-security main contrib non-free non-free-firmware
    # deb-src https://security.debian.org/debian-security bookworm-security main contrib non-free non-free-firmware
    
    # deb https://deb.debian.org/debian/ bookworm non-free-firmware contrib main non-free" > "$sources_list"

}

# Set the mirror source based on user choice
case $choice in
    1)
        set_mirror_source "https://mirrors.tuna.tsinghua.edu.cn/debian/"
        ;;
    2)
        set_mirror_source "https://mirrors.aliyun.com/debian/"
        ;;
    *)
        echo "Invalid choice. Exiting."
        exit 1
        ;;
esac

# Update apt
if sudo apt update && sudo apt upgrade -y; then
    echo "System update and upgrade successful."
else
    echo "Error: System update or upgrade failed."
    exit 1
fi

# Add additional steps and error handling if needed
