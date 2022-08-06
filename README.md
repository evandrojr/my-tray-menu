# my-tray-menu

Delivers an easy set up tray menu.

![alt text](media/screen1.png "my-tray-menu")

[demo.webm](https://user-images.githubusercontent.com/939608/183258199-f92279f0-3a65-4c05-910b-db7bea126766.webm)

## Bug

I was unable to create channels dynamically, yet. It is necessary to edit main.go to add menu options.

## Dead simple configuration

Sample my-tray-meny.yaml:

```
Turn off screen: /bin/sh /home/j/lab/my-tray-menu/scripts/turn-off-screen.sh
Toggle microphone: /bin/sh /home/j/lab/my-tray-menu/scripts/toggle-microphone.sh
Kill process (xkill): xkill
Shutdown: sudo shutdown -n now
```

## Requirements

- Go 1.9

## Dependencies' set-up

Follow instructions for specific OS dependencies at:

https://github.com/getlantern/systray

## Usage

```
go get https://github.com/evandrojr/my-tray-menu
my-tray-menu
```

Manual:

```
git clone git@github.com:evandrojr/my-tray-menu.git
cd my-tray-menu
go get
go build
./my-tray-menu
```

🍻

## References:

1. https://github.com/Osuka42g/simple-clock-systray
1. https://github.com/getlantern/systray
