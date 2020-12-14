# photosquared

## Description

Small command line utility to generate a square image with a border out of a source image of arbitrary aspect ratio suitable for posting to Instagram.

As an amateur photographer I do all my raw processing in [https://www.captureone.com/en](Capture One) and want to keep the aspect ratio that makes the most sense for the composition. Unfortunately, Instagram works exclusively with square images. Before I wrote PhotoSquared, I'd have to take the processed image, copy it, make a new all white square image that was 300 pixels wider than the widest side of the source image and insert my processed image. So much work!

Enter PhotoSquared. I built PhotoSquared in an afternoon and there are probably a ton of features I could add to it. For now, it only has the single feature I needed. PhotoSquared is written in [https://golang.org/](Go) so you can build it on any machine that supports Go. Now I just process my images as I normally do and before transferring them to my phone I just run them through PhotoSquared and I'm done.

## Downloading

All of these instructions assume you are using a Mac because it's what I have. If someone finds this useful and has a Windows machines feel free to contribute some instructions.

* Download [https://golang.org/dl/](Go) from [https://golang.org/dl/](https://golang.org/dl/)
* Follow the [https://golang.org/doc/install](installation instructions) for your machine. They may be found here [https://golang.org/doc/install](https://golang.org/doc/install)
* Get a copy of the PhotoSquared source code here [https://github.com/denmark/photosquared](https://github.com/denmark/photosquared) and put it in your $GOPATH/src/ directory
* If you have `make` installed then simply go to $GOPATH/src/photosquared/ and type `make`
* If you do not have `make` you can compile PhotoSquared using the following command: `go build -o ${GOPATH}/bin/photosquared main.go`

## Usage

PhotoSquared takes as its only argument the full path to a JPEG, PNG, or GIF file. For example:

```
$ photosquared DSCF2300.jpg
Reading from: [DSCF2300.jpg]
Input Dimensions: 6000x4000; Output Dimensions: 6600x6600
Writing to: [DSCF2300-squared.jpg]
$
```

view your output:

```
$ open DSCF2300-squared.jpg
```

PhotoSquared should write the output file to the same directory where the input image is.

## Customization

There are two parameters embedded in the code that enable you to customize the width of the border and the color of the border.

```
borderPadding := 300                        // 300 pixels
borderColor := color.RGBA{255, 255, 255, 0} // White
```

`borderPadding` is the number of pixels wide the border is.
`borderColor` is the RGB value of the border color.

If you want to adjust either just edit `main.go` and rebuild PhotoSquared. At some point in the future I might take either one or both of these parameters as command line arguments.

## TODO

* Tests
* Windows support
* Try to wrap the built binaries into a Mac or Windows App so they can be run straight from a Capture One process recipe
* Take command line arguments for both borderPadding and borderColor
