# BMO

Home assistant.

## Raspberry Pi setup

```
sudo apt install libc6-dev libglu1-mesa-dev libgl1-mesa-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev libasound2-dev
```

Compile SDL 2.10 and link shared lib:
```
ln -s /usr/local/lib/libSDL2-2.0.so.0 /usr/lib/libSDL2-2.0.so.0
```

## Docs

YeeLight API: https://www.yeelight.com/download/Yeelight_Inter-Operation_Spec.pdf
