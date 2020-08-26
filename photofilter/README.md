# Photo Filter Project

This project leads you through a series of exercises for creating photographic filters. It involves:

- Using [the `image` package](https://golang.org/pkg/image/).
- Creating a command-line program.
- Working with `interface`s.
- Fun photo manipulation.
- Creating a web server.

## Resources

Before you jump into playing with photos, here are some resources you may want
to take a look at (not required, but interesting):

- [A video about Color](https://www.youtube.com/watch?v=FTKP0Y9MVus&feature=emb_logo)
- [color theory](https://99designs.com/blog/tips/the-7-step-guide-to-understanding-color-theory/) (not as engaging as the video, but I wanted a second link)

## Exercises

### Exercise #1: Filter #1

Create a program that is given a photograph and applies a filter
of your own making to it. It will be invoked something like this...

```
myprogram --infile /path/to/my/file.jpg --outfolder /path/for/created/image
```

...and will end up creating the file `/path/for/created/image/DATE-file.jpg` (for
example: `/path/for/created/image/2020-09-22_14-05-22-file.jpg`), which is the
original file with the filter applied to it.

__"But what filter?"__ Your first filter will _convert the source image to grayscale_.

From this:

![color image](balloons-color.jpg)

To this:

![gray image](balloons-gray.jpg)

Your program will:

- Process both JPEG and PNG images.
- Output all images as JPEGs.

Your code should:

- have at least one package (not just `main`)
- have tests
- be checked into github (I recommend making a repo for all your exercise work.)

### Exercise #2: Image from URL

Change your program to allow the user to specify an image from the Internet
instead of from a local file:

```
myprogram --infile https://i.pinimg.com/originals/36/df/d6/36dfd61e6eeda1f9fc442dcc20c99bef.jpg --outfolder /path/for/created/image
```

Your program will:

- Determine if the image is local or on the Internet and respond appropriately.
- If the image is on the Internet, download it to the local machine before processing.


### Exercise #3: Second Filter

Create a second filter that divides an image into same-sized rectangles. For
each rectangle, adjust the color of each pixel to be its original color
blended with the average color of the rectangle.

Something like this:

![mosaic filter](color-mosaic.jpg)

Add an argument to your program so the user can pick which filter. (You will likely want to give them names.) Something like:

```
myprogram --infile https://foo.com/some/internet/photo.jpg --outfolder /path/for/created/image --filter mosaic
```

### Exercise #4: Third Filter & Apply Multiple Filters

Create a _third_ filter that divides the source image into rectangles and then
rearranges those rectangles randomly, producing a scrambled image.

Also, allow user to specify one _or more_ filters to apply. Something like:

```
myprogram --infile https://foo.com/some/internet/photo.jpg --outfolder /path/for/created/image --filters "mosaic,scramble,gray"
```

By this point, I am hoping that your program has taken advantage of creating
an `interface` or two to simplify working with multiple filters.


### Exercise #5: Web Server

Make this available via a web page. You should be able to type this into a
browser address bar and see the filtered image:

```
http://127.0.0.1/filter?file="https://foo.com/some/internet/photo.jpg"&filters="scramble,blend"
```