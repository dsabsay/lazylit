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
