# Sun, Aug 09 2020 05:35 PM
Trying to track down bug that causes all code to disappear. All the notes/docs 
on the left-hand side is there, just no code on the right side. I first noticed 
this when I tried to use lazylit at work.

I tried all the Release binaries against the artifact in the lazylit-example 
repo, but they all produced the same problem.

But I got it to work at some point because the published version of 
lazylit-example is correct. I don't know if I was using a release version to do 
that.

I tried compiling from source at several different commits, and they all 
produced the same problem...

I'm pretty stuck.

# Thu, Jul 02 2020 12:29 PM
* Summary:
    * Generated index page for each artifact. Right now links don't work because
        the actual doc HTML files aren't being put in the expected place.
* Next step is to remove existing render locations, and feed the Artifacts
    structure to the existing parser/generation functions.

# Fri, Jul 03 2020 12:49 PM
* Summary:
    * Index page links all work.
    * Created about page.
    * Completed doc page generation: removed header lines, added revision menu.
* Next steps:
    * Test with GH pages.
    * Write "init" command.
    * Test error messages (i.e. does it tell you if something is wrong?)
    * Make easy to install and use.
