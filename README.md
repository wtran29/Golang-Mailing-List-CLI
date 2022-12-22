# Setup

Project requires a `gcc` compiler installed and `protobuf` code generation tools.

# Install protobuf compiler

Download the appropriate binary for your operating system from [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases).

# Windows Installation and Setup

A. Install Git and Visual Studio 2019 or later.

  1. Open Visual Studio and go to the `Tools` menu.

  2. Select `"Get Tools and Features"` to open the Visual Studio Installer.

  3. In the Visual Studio Installer, go to the `"Workloads"` tab.

  4. Scroll down to the `"Linux development with C++"` workload and select the `"VC++ tools for CMake and Linux"` option.

  5. Click on the "Install" button to install the workload package.

  6. Follow the prompts to complete the installation.

B. Open a terminal window and clone the vcpkg repository using the following command:
```
git clone https://github.com/microsoft/vcpkg.git
```
C. Navigate to the vcpkg directory:

```
cd vcpkg
```
D. Run the following command to bootstrap vcpkg:
```
.\bootstrap-vcpkg.bat
```
This will build vcpkg and prepare it for use.

E. Run the following command to install protoc:
```
vcpkg install protobuf
```
This will install the protoc binary and the necessary dependencies.

F. Add the C:\vcpkg\installed\x64-windows\tools\protobuf directory to the PATH environment variable. You can do this by modifying the system environment variables or by adding the following line to the PowerShell profile file (%USERPROFILE%\Documents\WindowsPowerShell\Microsoft.PowerShell_profile.ps1):
```
$env:Path += ";C:\vcpkg\installed\x64-windows\tools\protobuf"
```
G. Reload the PowerShell profile by running the following command:

```
. $Profile
```
H. Install by placing the extracted binary somehwere in your `$PATH`.

# VS Code Settings

To make the protoc binary available in the Bash terminal of VS Code:

1. Open VS Code and go to the File menu.

2. Select "Preferences" and then "Settings".

3. In the Settings editor, search for "terminal.integrated.env".

4. This will show you the terminal environment variables that are set in VS Code.

5. To add the path to the protoc binary to the PATH environment variable, you can add the following line to the terminal environment variables:

```
"terminal.integrated.env.linux": {
    "PATH": "${env:PATH}:/path/to/protoc"
},

"terminal.integrated.env.windows": {
    "PATH": "${env:PATH};/path/to/protoc"
}
```
This will add the /path/to/protoc directory to the PATH environment variable in the Bash terminal.

6. Save the changes to the settings by clicking on the "Save" button or by pressing Ctrl + S.

7. Close and reopen the terminal window to apply the changes.
## Install Go protobuf codegen tools

`go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

## Generate Go code from .proto files

```
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/mail.proto
```