#!bin/sh

os=$(uname -s)
arch=$(uname -m)
installer_name='axilock-installer-secret-prevention'


echo_red(){
    echo "\033[31m$@\033[0m"
}
echo_green(){
    echo "\033[32m$@\033[0m"
}

if [ -e $installer_name ]; then
    echo "Installer file $installer_name already exists."
    echo "This may be due to a previous installation attempt."
    echo "Please remove the file $installer_name and try again."
    echo "Run the following command to remove the file:"
    echo "rm $installer_name"
    exit 1
fi

failed(){
    echo_red "Installation failed. Please try again."
    rm $installer_name
    exit 1
}
set -eE
trap failed ERR INT

if ! which trufflehog > /dev/null; then
    echo "Installing dependencies"
    if [ "$os" = "Linux" ]; then
        sudo apt-get install -y trufflehog
    elif [ "$os" = "Darwin" ]; then
        HOMEBREW_NO_ENV_HINTS=true HOMEBREW_NO_INSTALL_CLEANUP=true HOMEBREW_NO_AUTO_UPDATE=true brew install -q trufflehog
    elif [ "$os" = "Windows_NT" ]; then
        choco install trufflehog
    else
        echo "Unsupported OS: $os"
        exit 1
    fi
fi

if [ "$os" = "Linux" ]; then
    os="OS_LINUX"
elif [ "$os" = "Darwin" ]; then
    os="OS_DARWIN"
elif [ "$os" = "Windows_NT" ]; then
    os="OS_WINDOWS"
else
    echo "Unsupported OS: $os"
    exit 1
fi

if [ "$arch" = "x86_64" ]; then
    arch="ARCH_AMD64"
elif [ "$arch" = "arm64" ]; then
    arch="ARCH_ARM64"
elif [ "$arch" = "aarch64" ]; then
    arch="ARCH_ARM64"
else
    echo "Unsupported architecture: $arch"
    exit 1
fi

echo "Downloading Axilock installer"
curl --progress-bar --location 'https://api.axilock.ai/v1/client/update' \
--header 'Content-Type: application/json' \
--data '{
    "os" : "'$os'",
    "arch" : "'$arch'"
}' > $installer_name

chmod +x $installer_name
./$installer_name install

if [ $? -ne 0 ]; then
    echo_red "Installation failed"
    echo "Please contact us at contact@axilock.ai"
    rm $installer_name
    exit 1
fi

echo_green "Installation completed successfully"
rm $installer_name