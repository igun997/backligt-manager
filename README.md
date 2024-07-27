# Backlight Manager

This is a simple GUI application to manage the keyboard backlight using the `brightnessctl` tool. The application allows you to set the backlight level and specify the device name.

## Requirements

- Go 1.16 or higher
- `brightnessctl`
- `sudo` permissions to run `brightnessctl` commands

### Installing `brightnessctl`

#### For Debian-based systems (e.g., Ubuntu):
```sh
sudo apt-get update
sudo apt-get install brightnessctl
```

### Building the Application

First, ensure you have Go installed. Then, clone the repository and build the application with CGO enabled:

```sh
git clone https://github.com/igun997/backlight-manager.git
cd backlight-manager
CGO_ENABLED=1 go build -o backlight_manager
```

### Screenshoot 

![Backlight Manager Screenshot](https://raw.githubusercontent.com/igun997/backligt-manager/master/screenshot.png)