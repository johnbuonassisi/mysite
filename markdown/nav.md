In this post I will describe how I built a fixed Navigation Bar for my website using HTML 
and CSS. This was one of the first things I did when making my site. There were a few tricks 
I learned along the way that I wanted to share. I will assume that readers are familiar with 
basic HTML and CSS.

### Requirements for the Navigation Bar
The Navigation Bar should:

- Be fixed to the top of the page so it does not move when the page is scrolled
- Contain text Navigation Items that when clicked go to different sections of the same page or go to another page
- Span the full width of the page
- Be a fixed height

### Content for the Navigation Bar

To create the HTML for the Navigation Bar use the [navigation
section element, `<nav>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav). This element is intended to contain navigation links, either within the current page or 
other pages. Within the nav element, use a series of [anchor elements, `<a>`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a) with each [href attribute](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a#href) specifying a hyperlink to a location on the same page or links to other pages. 

In the example below, the first 5 anchors specify a hyperlink to a location on the current page using the pound symbol `#`. These hyperlinks link to an html element on the same page with an `id` attribute equal to what comes after the pound symbol. The final anchor specifies a hyperlink to a separate `blog` page. 

The HTML is not styled at this point so it will be styled using browser defaults. Anchor 
elements are inline elements so this is why they end up on the same line.

<script async src="//jsfiddle.net/05hmaoyg/2/embed/html,result/"></script>

### Styling the Navigation Bar

Now the fun part, let's add some styling so the Navigation Bar looks a bit better. The first thing to do is to fix the Navigation Bar to the top of the page and make it stretch the width of the page. This is done by setting the `position` property to `fixed`, setting `top` and `left` to `0`, forcing `nav` to be in the `top` `left` corner, and then set `width` to `100%` of the page.

<script async src="//jsfiddle.net/05hmaoyg/5/embed/html,css,result/"></script>

Next, we need to style the Navigation Items so they are equally spaced across the Navigation Bar. For this we will make the nav element a [flexbox](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Flexible_Box_Layout) container by setting the `style` to `flex`. The children elements of the flex container are aligned on the containers' major axis, which by default is horizontal. 

To space the child elements we must specify how to deal with free
space in the container. We will make all child elements of the same class `navitem` and specify the `flex` property as `auto`. This will make the navitems take all the free space
of the container and distrubute the available space evenly amongst eachother. 

Some final adjustments involve, setting `text-alignment` to `center` to make the text appear in the center of the flex item, setting `text-decoration` to none to remove the default 
underline of a hyperlink, and creating a new css property that will change the colour of
a navitem when the mouse hovers over.

<script async src="//jsfiddle.net/05hmaoyg/6/embed/html,css,result/"></script>

### Problem with Navigating to Sections Within a Page

Navigation Items containing a link starting with `#` should accomplish this... however there is a problem. If you try one of these links the top of the section will be covered by the Navigation Bar, because the Navigation Bar is fixed. What we really want is the link to a section to display the section just below the Navigation Bar, not directly under it! This is hard to describe in words, see this [example](http://jsfiddle.net/bn5whLe0/4/) (this example could not be embedded properly due to the
size of the blog page).

You will see that the `About` text for the `About` section is not shown, it is actually under 
the Navigation Bar. If you navigate to the `Skills` section you won't see the `Skills` text
until you scroll up slightly. The same issue will occur for the `Timeline` section. This could be solved by fiddling with the top padding and margin of each section, but this is 
not a general solution and would need to be applied to every section.

### Fixing Navigation to Sections Within a Page

I solved this problem by first adding `nav-spacer` under the Navigation Bar with the same height. This makes the first section appear just under the Navigation Bar when
the page first loads. I then grouped each section into a section element containing a `target` element and a `content` element. By setting the `position` of the `target` element to `absolute` and the `section` element to `relative`, I was able to move it above its containing element by the height of the Navigation Bar. I set the ids of the target elements to be the same one used by the Navigation Item links. Now, when a Navigation Item is selected it will link to the target element which is above the section content by the same height as the Navigation Bar. This will cause the content of the section to be displayed just below the Navigation Bar. Checkout the example [here](http://jsfiddle.net/bn5whLe0/5/).

### Navigation Bar Complete!

The Navigation Bar you see on this page is based on what I have described in this post. The
full source code for my website, which includes the Navigation Bar can be found [here](https://github.com/johnbuonassisi/mysite). If you are interested in just the Navigation Bar, see
this [example](http://jsfiddle.net/bn5whLe0/5/).

``` go
os.Exit(1)
```
