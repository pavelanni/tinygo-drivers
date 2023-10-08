# TinyGo drivers

This repository contains TinyGo drivers either written by me,
or the drivers from the original repo slightly modified for specific devices (e.g. `tm1637`).

## TM1637

This driver was modified by changing the delay time to address an issue with a specific implementation.
Thanks to @bxparks for the solution. See more in the README.

## Rotary encoder

This rotary encoder driver is a slightly modified version of https://github.com/bgould/tinygo-rotary-encoder
I added the encoder channel to read the movement from and the switch channel.

