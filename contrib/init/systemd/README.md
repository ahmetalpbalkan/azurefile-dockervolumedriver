# Ubuntu (systemd) installation instructions

> NOTE: These instructions are valid for Ubuntu versions using `systemd` init
> system.

## TL;DR
1. Get the latest [release](https://github.com/Azure/azurefile-dockervolumedriver/releases)
2. Put the binary into `/usr/bin/azurefile-dockervolumedriver`
3. Get the .default and .service files from the .tar.gz file and deploy them
4. Reload systemd

## In-depth walkthrough
To ssh onto machines created with the `docker-machine` commands, you'll probably need to ssh via `docker-machine ssh [machinename]` to perform the following.

0. `sudo -s`
0. Download the tar.gz from "Releases" tab of the repo for copies of config files
    + Currently this would be `wget https://github.com/Azure/azurefile-dockervolumedriver/archive/0.2.1.tar.gz` 
    + Decompress the archive:  `tar -xvf 0.2.1.tar.gz`
0. Download the binary from the "Releases" tab of the repo to `/usr/bin/azurefile-dockervolumedriver`
    + Use wget: `wget -qO/usr/bin/azurefile-dockervolumedriver https://github.com/Azure/azurefile-dockervolumedriver/releases/download/0.2.1/azurefile-dockervolumedriver`
    + Make it executable `chmod +x /opt/bin/azurefile-dockervolumedriver`
0. Save the `.default` file to `/etc/default/azurefile-dockervolumedriver`
    + Copy the file: `cp azurefile-dockervolumedriver-0.2.1/contrib/init/systemd/azurefile-dockervolumedriver.default /etc/default/azurefile-dockervolumedriver`
0. Edit `/etc/default/azurefile-dockervolumedriver` with your Azure Storage Account credentials.
    + Edit in vi (or another editor): `vi /etc/default/azurefile-dockervolumedriver`
    + Use ESC, :, x, ENTER to save an exit if you use vi
0. Save the `.service` file to `/etc/systemd/system/azurefile-dockervolumedriver.service`
    + Make the requisite directories if they don't exist: `mkdir /etc/systemd && mkdir /etc/systemd/system`
    + Copy the relevant file: `cp azurefile-dockervolumedriver-0.2.1/contrib/init/systemd/azurefile-dockervolumedriver.service /etc/systemd/system/`
0. Run `systemctl daemon-reload`
0. Run `systemctl enable azurefile-dockervolumedriver`
0. Run `systemctl start azurefile-dockervolumedriver`
0. Check status via `systemctl status azurefile-dockervolumedriver`

To test, from your base machine i.e. not inside your docker-machine, try by creating a volume and running a container with it:

    docker volume create -d azurefile --name myvol -o share=myvol
    docker run -i -t -v myvol:/data busybox
    # cd /data
    # touch a.txt

You can find the logs at `journalctl -fu azurefile-dockervolumedriver`.
