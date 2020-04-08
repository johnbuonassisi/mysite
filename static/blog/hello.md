``` go
fmt.Println("Hello World!")
```

Hello World! My name is John Buonassisi and welcome to my website and blog. I have worked professionally for 5 years as a Software Engineer and have recently become a new father. I have some free time during my parental leave (which happens to coincide with the COVID-19 pandemic), so what better time to create a website to collect my thoughts and express my opinions on topics I am interested in!

I wanted to start with my journey making a website from scratch without the use of website builders or web frameworks. In the past, I tried to use [Wordpress](https://wordpress.com/) to create my personal site and was shocked by the terrible user experience. Constructing the site using their management system and hodge-pudge of page builders was really confusing and inflexible. I concluded that I would need to buy themes and plugins to make the site appealing and functional. I decided that instead of learning the ins and outs of Wordpress, my time could be better spent learning the underlying technologies of the web. After all, I am a Software Engineer, how hard could it be?

### Start from first principles; HTML, CSS, and Javascript

Whenever I learn something new I like to start from first principles. This usually involves taking an online course, watching tutorials on youtube, or reading documentation; anything that will give me a good base in the subject in a day or two. I have spent my career building backend systems, so HTML, CSS, and Javascript were fairly unfamiliar to me. To ramp up I turned to  the [Mozilla Developer Network](https://developer.mozilla.org/en-US/). This is an amazing resource provided by the makers of the Mozilla web browser.

> MDN's mission is simple: provide developers with the information they need to easily build projects on the open Web

I took some time to go through MDN's documentation on [HTML](https://developer.mozilla.org/en-US/docs/Learn/HTML), [CSS](https://developer.mozilla.org/en-US/docs/Learn/CSS), and [Javascript](https://developer.mozilla.org/en-US/docs/Learn/JavaScript). Their documentation is very clear and provides tons of working examples and problems that can be worked through to gain understanding. With a firmer grasp of these technologies I set out to create my personal website.

### Start prototyping

Armed with more knowledge about web technologies I decided to start playing around with HTML and CSS. I wanted my site to be a simple resume style, single page site. It should detail who I am, what I do, and the past experiences I decided to split the page into the following sections; introduction, skills, timeline, projects, and contact. At the top of the page there is a fixed navigation bar for jumping to the different sections, this navigation bar does not move when the page is scrolled. The bottom of the page has a "sticky" footer that always sticks to the end of the page. The result is the website you are currently on.

It took me a few iterations to figure out how to build each section so I want to detail each section in my nexts posts. I will describe how I built features in HTML and CSS, then I will get into how I chose to serve the site, and finally how I deployed the site using a cloud provider. I eventually added the blog page which I will also write about in future posts.

``` go
os.Exit(1)
```