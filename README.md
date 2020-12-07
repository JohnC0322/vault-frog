# vault-frog

## About this plugin
This plugin is going to leverage the hashivault encrypted password for Jfrog CLI use:
* Storing the connection details of multiple Artifactory instances with encryted password.
* Using the stored connection details for executing any Artifactory JFrog CLI command.
* Deleting the stored connection details from the Hashi Vault.

## Installation with JFrog CLI
Installing the latest version:

`$ jfrog plugin install vault-frog`

Installing a specific version:

`$ jfrog plugin install vault-frog@version`

Uninstalling a plugin

`$ jfrog plugin uninstall vault-frog`

## Usage
### Commands
* Store the Artifactory connection details is the Hashi Vault.
    - Name: store
    - Alias: s
    - Arguments:
        No arguments
    - Flags:
        - --server-id: Artifactory server ID used for accessing the configuration. **[Mandatory]**
        - --url: Artifactory URL. **[Mandatory]**
        - --user: Artifactory username. **[Mandatory]**
        - --password: Artifactory password. **[Mandatory]**
        
    - Example:
    ```
  $ jfrog vault store --server-id Artifactory-1 --url https://my-rt.com/artifactory --user my-user --password my-password
  ```
* Use the stored Artifactory connection details for executing any Artifactory JFrog CLI command.
    - Name: use
    - Alias: u
    - Arguments:
        Stored server ID
    - Flags:
        No flags
        
    - Examples:
    ```
  $ jfrog vault use Artifactory-1 search "generic-local/path/"
  $ jfrog vault use Artifactory-1 download "generic-local/path/" --flat --recursive
  ```
* Delete Artifactory connection details from the OS keyring.
    - Name: delete
    - Alias: del
    - Arguments:
        Stored server ID
    - Flags:
        No flags
        
    - Example:
    ```
  $ jfrog vault del Artifactory-1
  ```

## Release Notes
The release notes are available [here](RELEASE.md).