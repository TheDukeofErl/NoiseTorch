<h1 align="center"> yant</h1>
<h3 align="center"> Yet another noise tool, a noise suppression program for Pipewire</h3>
<p align="center"><img src="https://raw.githubusercontent.com/TheDukeofErl/yant/master/assets/icon/noisetorch.png" width="100" height="100"></p> 


<div align="center">
    
  <a href="">[![Licence][licence]][licence-url]</a>
  <a href="">[![Latest][version]][version-url]</a>
    
</div>

[licence]: https://img.shields.io/badge/License-GPLv3-blue.svg
[licence-url]: https://www.gnu.org/licenses/gpl-3.0
[version]: https://img.shields.io/github/v/release/noisetorch/NoiseTorch?label=Latest&style=flat
[version-url]: https://github.com/noisetorch/NoiseTorch/releases
[stars-shield]: https://img.shields.io/github/stars/noisetorch/NoiseTorch?maxAge=2592000
[stars-url]: https://github.com/noisetorch/NoiseTorch/stargazers/

Yet another noise tool, or yant, is an easy to use open source application for Linux with PipeWire. A fork from NoiseTorch/NoiseTorch-ng, it creates a virtual microphone that suppresses noise in any application using [RNNoise](https://github.com/xiph/rnnoise). Use whichever conferencing or VOIP application you like and simply select the filtered Virtual Microphone as input to torch the sound of your mechanical keyboard, computer fans, trains and the likes.

PulseAudio support is being dropped when compared to NoiseTorch-ng to allow for distribution via flatpak.

Don't forget to leave a star ⭐ if this sounds useful to you! 

## Roadmap
* Update roadmap as needed
* Gut autoupdating code
* Remove support for PulseAudio
* Test
* Deploy via flathub

## Features
* Simple setup for microphone denoising

## Download & Install

Planning on distributing this via flathub

## Usage

Select the microphone you want to denoise, and click "Load", yant will create a virtual microphone called "Filtered Microphone" that you can select in any application. Output filtering works the same way, simply output the applications you want to filter to "Filtered Headphones".

When you're done using it, simply click "Unload" to remove it again, until you need it next time.

The slider "Voice Activation Threshold" under settings, allows you to choose how strict yant should be in only allowing your microphone to send sounds when it detects voice.. Generally you want this up as high as possible. With a decent microphone, you can turn this to the maximum of 95%. If you cut out during talking, slowly lower this strictness until you find a value that works for you.

If you set this to 0%, yant will still dampen noise, but not deactivate your microphone if it doesn't detect voice.

Please keep in mind that you will need to reload yant for these changes to apply.

Once ya nt has been loaded, feel free to close the window, the virtual microphone will continue working until you explicitly unload it. The yant process is not required anymore once it has been loaded.

## FAQs

### Latency

Yant may introduce a small amount of latency for microphone filtering. The amount of inherent latency introduced by noise supression is 10ms, this is very low and should not be a problem.

Output filtering currently introduces something on the order of ~100ms. This should still be fine for regular conferences, VOIPing and gaming. Maybe not for competitive gaming teams.

### Alternatives

- [noise-suppression-for-voice](https://github.com/werman/noise-suppression-for-voice): Denoising software which uses rnnoise. More complex to configure but offers more options. Requires more use of the terminal.

- [Easy Effects](https://github.com/wwmm/easyeffects): Package which offers a large number of different audio effects such as echo cancellation or noise removal. More complex to configure and only supports PipeWire. Denoising uses rnnoise.

## Building (dev) from source

Install the Go compiler from [golang.org](https://golang.org/). And make sure you have a working C++ compiler.

```shell
 git clone https://github.com/TheDukeofErl/yant # Clone the repository
 cd yant # cd into the cloned repository
 make dev # build it
```

## Special thanks to

* [@lawl](https://github.com/lawl), the original creator of NoiseTorch
* The [NoiseTorch-ng Community](https://github.com/noisetorch/NoiseTorch), who kept the project alive
* [xiph.org](https://xiph.org)/[Mozilla's](https://mozilla.org) excellent [RNNoise](https://jmvalin.ca/demo/rnnoise/).
* [@werman](https://github.com/werman/)'s [noise-suppression-for-voice](https://github.com/werman/noise-suppression-for-voice/) for the inspiration
* [@aarzilli](https://github.com/aarzilli/)'s [nucular](https://github.com/aarzilli/nucular) GUI toolkit for Go.
* [Sallee Design](https://www.salleedesign.com) (info@salleedesign.com)'s Microphone Icon under [CC BY 4.0](https://creativecommons.org/licenses/by/4.0/)
