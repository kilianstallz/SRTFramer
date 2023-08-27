
# SrtFramer

## Introduction

SrtFramer is a command-line tool designed to extract (key)-frames from video files captured by DJI drones. This tool utilizes the generated SRT (SubRip Subtitle) files to enrich the extracted frames with GPS metadata.

## How to use

### Prerequisites

- ffmpeg (https://ffmpeg.org/)

```shell
brew install ffmpeg
```

- exiftool (https://exiftool.org/)

```shell
brew install exiftool
```

### Installation

- Install from source
```shell
go install github.com/kilianstallz/srtframer
```

- Install from latest GitHub [Release](https://github.com/kilianstallz/SRTFramer/releases)

### Usage

```shell
srtframer -i <input> -o <output> -s <srt>
```


## License

[MIT](https://choosealicense.com/licenses/mit/)
