# qnap-tool

A useful command line program to manage QNAP NAS.

## Example

```shell
    # Power on (Wake-on-LAN)
    qnap-tool wake 11:22:33:44:55:66 # MAC of NAS

    # Power off
    qnap-tool shutdown 1.2.3.4 # IP of NAS
```

## Full usage

```shell
    qnap-tool --help

    Usage of ./qnap-tool:

    Utility of QNAP NAS.

    qnap-tool [flags] action arg

    Actions:
    wake:  Wake up a NAS by Wake-on-LAN.
        qnap-tool wake MAC-ADDRESS
    shutdown:  Shuts down a NAS.
        qnap-tool [-user -password -port -https] shutdown HOST-OF-NAS.
    
    Flags:
    -https
            Whether the admin page is https enabled
    -password string
            The password of user to login
    -port uint
            The port of the admin web page URL (default 5000)
    -user string
            The user name to login
```
