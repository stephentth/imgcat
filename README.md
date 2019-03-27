# imgcat

## Introduce

A simple utily to render an image to terminal. Support load image from local disk and from image url.

Features:
* Ouput image 2 pixel per glyph.
* Can fetch image from url.
* Auto analize your terminal size and scale output image base on it.

## Screenshots

<center>
    <p>Simple cat an image local</p>
    <img src="docs/images/screenshot_02.png">
    <p>Simple cat an image remote</p>
    <img src="docs/images/screenshot_01.png">
</center>

## Install

If you have go installed in your system.

```
go get -u github.com/stephentt-me/imgcat
```

Make sure `$GOPATH/bin` is in your `$PATH`.

Or [download binary here](#) and put it into your `$PATH`.

## Usage

"Cat" image from file
```
imgcat foo.jpg
```

"Curl" image from url
```
imgcat <url>
```

## License

MIT